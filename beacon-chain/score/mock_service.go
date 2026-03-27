package score

type MockService struct {
	Scores map[[48]byte]uint64

	EpochScore  uint64
	TargetCount uint64

	RegisterCalls int

	BlockNumber uint64
	StartBlock  uint64
	EndBlock    uint64
}

func (m *MockService) GetScore(pubkey [48]byte) uint64 {
	if m.Scores == nil {
		return 0
	}
	return m.Scores[pubkey]
}

func (m *MockService) GetEpochRangeScore() uint64 {
	return m.EpochScore
}

func (m *MockService) TargetValidatorsCount() uint64 {
	return m.TargetCount
}

func (m *MockService) SetBlockNumber(block uint64) {
	m.BlockNumber = block
}

func (m *MockService) SetBlockRange(start, end uint64) {
	m.StartBlock = start
	m.EndBlock = end
}
