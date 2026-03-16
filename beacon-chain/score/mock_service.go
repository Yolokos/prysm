package score

type MockService struct {
	RegisterCalls int
}

func (m *MockService) GetScore(pubkey [48]byte) uint64 {
	return 700
}

func (m *MockService) RegisterValidator(pubkey [48]byte) error {
	m.RegisterCalls++
	return nil
}

func (m *MockService) TargetValidatorsCount() int {
	return 1
}
