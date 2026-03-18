package score

type MockService struct {
	RegisterCalls int
	TargetCount   int
}

func (m *MockService) GetScore(pubkey [48]byte) uint64 {
	return 700
}

func (m *MockService) RegisterValidator(pubkey [48]byte) error {
	m.RegisterCalls++
	return nil
}

func (m *MockService) TargetValidatorsCount() uint64 {
	return uint64(m.TargetCount)
}

func (m *MockService) GetEpochRangeScore() uint64 {
	return 200
}

func (m *MockService) SetEpochStart(startBlock uint64) {
}

func (m *MockService) SetEpochEnd(endBlock uint64) {
}
