package score

type MockServiceLow struct {
	RegisterCalls int
	TargetCount   int
}

func (m *MockServiceLow) GetScore(pubkey [48]byte) uint64 {
	if pubkey[0] == 0 {
		return 100
	}
	return 900
}

func (m *MockServiceLow) RegisterValidator(pubkey [48]byte) error {
	m.RegisterCalls++
	return nil
}

func (m *MockServiceLow) TargetValidatorsCount() uint64 {
	return uint64(m.TargetCount)
}

func (m *MockServiceLow) GetEpochRangeScore() uint64 {
	return 200
}

func (m *MockServiceLow) SetEpochStart(startBlock uint64) {
}

func (m *MockServiceLow) SetEpochEnd(endBlock uint64) {
}
