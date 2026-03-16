package score

type MockServiceMixed struct {
	RegisterCalls int
	TargetCount   int
}

func (m *MockServiceMixed) GetScore(pubkey [48]byte) uint64 {
	switch pubkey[0] {

	case 0:
		return 900

	case 1:
		return 700

	case 2:
		return 400

	default:
		return 200
	}
}

func (m *MockServiceMixed) RegisterValidator(pubkey [48]byte) error {
	m.RegisterCalls++
	return nil
}

func (m *MockServiceMixed) TargetValidatorsCount() int {
	return m.TargetCount
}
