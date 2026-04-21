package electra

import (
	"github.com/OffchainLabs/prysm/v7/beacon-chain/score"
	"github.com/OffchainLabs/prysm/v7/beacon-chain/state"
	"github.com/OffchainLabs/prysm/v7/time/slots"
	"github.com/sirupsen/logrus"
)

func GetPendingValidatorRegistrations(st state.BeaconState) [][]byte {
	scoreService, err := score.GetService()
	if err != nil {
		return nil
	}

	validators := st.Validators()
	currentEpoch := slots.ToEpoch(st.Slot())

	out := make([][]byte, 0)
	logrus.Infof("Total validators in state: %d, current epoch: %d", len(validators), currentEpoch)
	for i := 0; i < len(validators); i++ {
		v := validators[i]
		logrus.Infof("Validator %d: pubkey=%x, activation_epoch=%d, exit_epoch=%d", i, v.PublicKey, v.ActivationEpoch, v.ExitEpoch)

		// только недавно активированные
		if v.ActivationEpoch != currentEpoch {
			continue
		}

		// не вышел
		if v.ExitEpoch <= currentEpoch {
			continue
		}

		isRegistered, err := scoreService.IsValidatorRegistered([48]byte(v.PublicKey))
		if err != nil {
			continue
		}
		if isRegistered {
			continue
		}

		// Копируем pubkey
		pubkey := make([]byte, len(v.PublicKey))
		copy(pubkey, v.PublicKey)

		out = append(out, pubkey)
	}

	return out
}
