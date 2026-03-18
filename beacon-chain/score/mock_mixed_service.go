package score

type MockServiceMixed struct {
	RegisterCalls int
	TargetCount   uint64
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

func (m *MockServiceMixed) TargetValidatorsCount() uint64 {
	return uint64(m.TargetCount)
}

func (m *MockServiceMixed) GetEpochRangeScore() uint64 {
	return 200
}

func (m *MockServiceMixed) SetEpochStart(startBlock uint64) {
}

func (m *MockServiceMixed) SetEpochEnd(endBlock uint64) {
}
