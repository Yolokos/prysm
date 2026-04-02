package score

import "errors"

type Service interface {
	GetScore(pubkey [48]byte) (uint64, error)
	GetEpochRangeScore() (uint64, error)
	TargetValidatorsCount() (uint64, error)

	SetBlockNumber(block uint64)
	SetBlockRange(start, end uint64)
	IsValidatorRegistered(pubkey [48]byte) (bool, error)
}

var svc Service

func InitService(s Service) {
	svc = s
}

func GetService() (Service, error) {
	if svc == nil {
		return nil, errors.New("score service not initialized")
	}
	return svc, nil
}
