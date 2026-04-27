package electra

import (
	"sync"

	"github.com/OffchainLabs/prysm/v7/beacon-chain/score"
	"github.com/OffchainLabs/prysm/v7/beacon-chain/state"
	enginev1 "github.com/OffchainLabs/prysm/v7/proto/engine/v1"
	"github.com/OffchainLabs/prysm/v7/time/slots"
	"github.com/sirupsen/logrus"
)

type RegistrationCache struct {
	registered map[[48]byte]bool
	mu         sync.RWMutex
}

var cache = &RegistrationCache{
	registered: make(map[[48]byte]bool),
}

func GetPendingValidatorRegistrations(st state.BeaconState) []*enginev1.ValidatorRegistration {
	scoreService, err := score.GetService()
	if err != nil {
		return nil
	}

	validators := st.Validators()
	currentEpoch := slots.ToEpoch(st.Slot())

	out := make([]*enginev1.ValidatorRegistration, 0)
	logrus.Infof("Total validators in state: %d, current epoch: %d", len(validators), currentEpoch)
	for i := 0; i < len(validators); i++ {
		v := validators[i]
		logrus.Infof("Validator %d: pubkey=%x, activation_epoch=%d, exit_epoch=%d", i, v.PublicKey, v.ActivationEpoch, v.ExitEpoch)

		pk := [48]byte(v.PublicKey)

		cache.mu.RLock()
		isRegistered := cache.registered[pk]
		cache.mu.RUnlock()

		if !isRegistered {
			isRegisteredRPC, err := scoreService.IsValidatorRegistered(pk)
			logrus.Infof("Validator %d registration status from RPC: %v, error: %v", i, isRegisteredRPC, err)
			if err != nil {
				continue
			}

			if isRegisteredRPC {
				cache.mu.Lock()
				cache.registered[pk] = true
				cache.mu.Unlock()
			}

			isRegistered = isRegisteredRPC
		}

		if isRegistered {
			continue
		}

		pubkey := make([]byte, len(v.PublicKey))
		copy(pubkey, v.PublicKey)

		out = append(out, &enginev1.ValidatorRegistration{
			Pubkey: pubkey,
		})
	}

	return out
}
