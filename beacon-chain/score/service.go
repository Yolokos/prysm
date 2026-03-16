package score

type Service interface {
	GetScore(pubkey [48]byte) uint64
	RegisterValidator(pubkey [48]byte) error
	TargetValidatorsCount() int
}
