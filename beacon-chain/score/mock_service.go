package score

type MockService struct{}

func (m *MockService) GetScore(pubkey [48]byte) uint64 {
	return 700
}
