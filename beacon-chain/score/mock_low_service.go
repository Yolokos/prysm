package score

type MockServiceLow struct {
	RegisterCalls int
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

func (m *MockServiceLow) TargetValidatorsCount() int {
	return 1
}
