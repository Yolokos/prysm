package score

type Service interface {
	GetScore(pubkey [48]byte) uint64
	RegisterValidator(pubkey [48]byte) error
	TargetValidatorsCount() uint64
	SetEpochStart(startBlock uint64)
	SetEpochEnd(endBlock uint64)
	GetEpochRangeScore() uint64
}
