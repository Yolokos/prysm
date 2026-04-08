package electra

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/OffchainLabs/prysm/v7/beacon-chain/core/helpers"
	"github.com/OffchainLabs/prysm/v7/beacon-chain/core/time"
	"github.com/OffchainLabs/prysm/v7/beacon-chain/core/validators"
	"github.com/OffchainLabs/prysm/v7/beacon-chain/score"
	"github.com/OffchainLabs/prysm/v7/beacon-chain/state"
	"github.com/OffchainLabs/prysm/v7/config/params"
	"github.com/OffchainLabs/prysm/v7/consensus-types/primitives"
	log "github.com/sirupsen/logrus"
)

// ProcessRegistryUpdates processes all validators eligible for the activation queue, all validators
// which should be ejected, and all validators which are eligible for activation from the queue.
//
// Spec pseudocode definition:
//
//	def process_registry_updates(state: BeaconState) -> None:
//	    # Process activation eligibility and ejections
//	    for index, validator in enumerate(state.validators):
//	        if is_eligible_for_activation_queue(validator):  # [Modified in Electra:EIP7251]
//	            validator.activation_eligibility_epoch = get_current_epoch(state) + 1
//
//	        if (
//	            is_active_validator(validator, get_current_epoch(state))
//	            and validator.effective_balance <= EJECTION_BALANCE
//	        ):
//	            initiate_validator_exit(state, ValidatorIndex(index))  # [Modified in Electra:EIP7251]
//
//	    # Activate all eligible validators
//	    # [Modified in Electra:EIP7251]
//	    activation_epoch = compute_activation_exit_epoch(get_current_epoch(state))
//	    for validator in state.validators:
//	        if is_eligible_for_activation(state, validator):
//	            validator.activation_epoch = activation_epoch
func ProcessRegistryUpdates(ctx context.Context, st state.BeaconState) error {
	currentEpoch := time.CurrentEpoch(st)
	ejectionBal := params.BeaconConfig().EjectionBalance
	activationEpoch := helpers.ActivationExitEpoch(currentEpoch)
	minScore := params.BeaconConfig().MinValidatorScore

	var err error

	scoreService, err := score.GetService()
	if err != nil {
		return fmt.Errorf("could not get score service: %w", err)
	}

	epochScore, err := scoreService.GetEpochRangeScore()
	if err != nil {
		return fmt.Errorf("could not get epoch range score: %w", err)
	}
	log.Infof("Epoch range score: %d, Current epoch: %d", epochScore, currentEpoch)
	// To avoid copying the state validator set via st.Validators(), we will perform a read only pass
	// over the validator set while collecting validator indices where the validator copy is actually
	// necessary, then we will process these operations.
	eligibleForActivationQ := make([]primitives.ValidatorIndex, 0)
	eligibleForEjection := make([]primitives.ValidatorIndex, 0)
	eligibleForActivation := make([]primitives.ValidatorIndex, 0)

	if err := st.ReadFromEveryValidator(func(idx int, val state.ReadOnlyValidator) error {
		// Score check to ensure we don't perform unnecessary state updates for validators that are not eligible for activation or ejection.
		isRegistered, err := scoreService.IsValidatorRegistered(val.PublicKey())
		if err != nil {
			return fmt.Errorf("could not check if validator %s is registered: %w", val.PublicKey(), err)
		}

		if !isRegistered {
			log.Infof("Validator %s is not registered, skipping score check", val.PublicKey())
			return nil
		}

		scoreValue, err := scoreService.GetScore(val.PublicKey())
		if err != nil {
			return fmt.Errorf("could not get score for validator %s: %w", val.PublicKey(), err)
		}
		// Collect validators eligible to enter the activation queue.
		if helpers.IsEligibleForActivationQueue(val, currentEpoch) {
			eligibleForActivationQ = append(eligibleForActivationQ, primitives.ValidatorIndex(idx))
		}

		// Collect validators to eject.
		if helpers.IsActiveValidatorUsingTrie(val, currentEpoch) &&
			(scoreValue < minScore || val.EffectiveBalance() <= ejectionBal) {
			log.Infof("Validator %s is eligible for ejection with score %d and effective balance %d", fmt.Sprintf("0x%x", val.PublicKey()), scoreValue, val.EffectiveBalance())
			eligibleForEjection = append(eligibleForEjection, primitives.ValidatorIndex(idx))
		}

		// Collect validators eligible for activation and not yet dequeued for activation.
		if helpers.IsEligibleForActivationUsingROVal(st, val) {
			eligibleForActivation = append(eligibleForActivation, primitives.ValidatorIndex(idx))
		}

		return nil
	}); err != nil {
		return fmt.Errorf("failed to read validators: %w", err)
	}

	log.Infof("Found %d validators eligible for activation queue, %d validators eligible for ejection, and %d validators eligible for activation", len(eligibleForActivationQ), len(eligibleForEjection), len(eligibleForActivation))

	// Handle validators eligible to join the activation queue.
	for _, idx := range eligibleForActivationQ {
		log.Infof("Validator %d is eligible for activation queue", idx)
		v, err := st.ValidatorAtIndex(idx)
		if err != nil {
			return err
		}
		v.ActivationEligibilityEpoch = currentEpoch + 1
		if err := st.UpdateValidatorAtIndex(idx, v); err != nil {
			return fmt.Errorf("failed to updated eligible validator %s: %w", fmt.Sprintf("0x%x", v.PublicKey), err)
		}
	}

	// Handle validator ejections.
	for _, idx := range eligibleForEjection {
		log.Infof("Initiating exit for validator %d", idx)
		var err error
		// exit info is not used in electra
		st, err = validators.InitiateValidatorExit(ctx, st, idx, &validators.ExitInfo{})
		if err != nil && !errors.Is(err, validators.ErrValidatorAlreadyExited) {
			return fmt.Errorf("failed to initiate validator exit at index %d: %w", idx, err)
		}
	}

	sort.SliceStable(eligibleForActivation, func(i, j int) bool {
		vi, _ := st.ValidatorAtIndex(eligibleForActivation[i])
		vj, _ := st.ValidatorAtIndex(eligibleForActivation[j])

		var pki, pkj [48]byte
		copy(pki[:], vi.PublicKey)
		copy(pkj[:], vj.PublicKey)

		si, err := scoreService.GetScore(pki)
		if err != nil {
			log.WithError(err).Errorf("Could not get score for validator %s, skipping", fmt.Sprintf("0x%x", vi.PublicKey))
			return false
		}
		sj, err := scoreService.GetScore(pkj)
		if err != nil {
			log.WithError(err).Errorf("Could not get score for validator %s, skipping", fmt.Sprintf("0x%x", vj.PublicKey))
			return false
		}

		pi := si >= epochScore
		pj := sj >= epochScore

		if pi != pj {
			return pi
		}

		wi := vi.EffectiveBalance * (1000 + si)
		wj := vj.EffectiveBalance * (1000 + sj)

		log.Infof("Validator %d has weight %d", eligibleForActivation[i], wi)
		log.Infof("Validator %d has weight %d", eligibleForActivation[j], wj)

		return wi > wj
	})

	limit := uint64(len(eligibleForActivation))
	churnLimit, err := scoreService.TargetValidatorsCount()
	if err != nil {
		return fmt.Errorf("could not get target validators count: %w", err)
	}
	log.Infof("Churn limit: %d", churnLimit)
	if churnLimit < limit {
		limit = churnLimit
	}

	// Activate all eligible validators.
	for _, idx := range eligibleForActivation[:limit] {
		log.Infof("Activating validator %d", idx)
		v, err := st.ValidatorAtIndex(idx)
		if err != nil {
			return err
		}
		v.ActivationEpoch = activationEpoch
		if err := st.UpdateValidatorAtIndex(idx, v); err != nil {
			return fmt.Errorf("failed to activate validator at index %d: %w", idx, err)
		}
	}

	return nil
}
