package score

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

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

func GetABI() (abi.ABI, error) {
	var err error
	parsedABI, err := abi.JSON(strings.NewReader(ScoreMetaData.ABI))
	if err != nil {
		return abi.ABI{}, err
	}
	return parsedABI, nil
}

func GetService() (Service, error) {
	if svc == nil {
		return nil, errors.New("score service not initialized")
	}
	return svc, nil
}
