package score

type MockServiceLow struct{}

func (m *MockServiceLow) GetScore(pubkey [48]byte) uint64 {
	return 100
}
