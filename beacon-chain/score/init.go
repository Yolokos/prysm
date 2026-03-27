package score

import "errors"

type Service interface {
	GetScore(pubkey [48]byte) uint64
	GetEpochRangeScore() uint64
	TargetValidatorsCount() uint64

	SetBlockNumber(block uint64)
	SetBlockRange(start, end uint64)
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
