package score

import "github.com/ethereum/go-ethereum/crypto"

type MockService struct {
	Scores map[[32]byte]uint64

	EpochScore  uint64
	TargetCount uint64

	RegisterCalls int

	BlockNumber uint64
	StartBlock  uint64
	EndBlock    uint64
}

func (m *MockService) GetScore(pubkey [48]byte) (uint64, error) {
	if m.Scores == nil {
		return 0, nil
	}
	key := crypto.Keccak256Hash(pubkey[:])
	return m.Scores[key], nil
}

func (m *MockService) GetEpochRangeScore() (uint64, error) {
	return m.EpochScore, nil
}

func (m *MockService) TargetValidatorsCount() (uint64, error) {
	return m.TargetCount, nil
}

func (m *MockService) SetBlockNumber(block uint64) {
	m.BlockNumber = block
}

func (m *MockService) IsValidatorRegistered(pubkey [48]byte) (bool, error) {
	if m.Scores == nil {
		return false, nil
	}
	key := crypto.Keccak256Hash(pubkey[:])

	if _, exists := m.Scores[key]; !exists {
		return false, nil
	}
	return true, nil
}

func (m *MockService) SetBlockRange(start, end uint64) {
	m.StartBlock = start
	m.EndBlock = end
}
