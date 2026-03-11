package score

type Service interface {
	GetScore(pubkey [48]byte) uint64
}
