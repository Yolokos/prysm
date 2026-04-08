package score

import (
	"context"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
)

type ValidatorAIService struct {
	ethClient *ethclient.Client
	contract  *ScoreCaller

	cache       map[uint64]map[[32]byte]uint64
	blockNumber uint64
	mu          sync.RWMutex

	startBlock uint64
	endBlock   uint64
}

func (s *ValidatorAIService) Start() {}

func (s *ValidatorAIService) Stop() error {
	s.ethClient.Close()
	return nil
}

func (s *ValidatorAIService) Status() error {
	return nil
}

func (s *ValidatorAIService) SetBlockNumber(block uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.blockNumber = block
}

func (s *ValidatorAIService) SetBlockRange(start, end uint64) {
	s.mu.Lock()
	s.startBlock = start
	s.endBlock = end
	s.mu.Unlock()
}

const maxBlocks = 100

func (s *ValidatorAIService) GetEpochRangeScore() (uint64, error) {
	s.mu.RLock()
	start := s.startBlock
	end := s.endBlock
	s.mu.RUnlock()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	callOpts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	score, err := s.contract.GetEpochRangeScore(callOpts,
		new(big.Int).SetUint64(start), new(big.Int).SetUint64(end))
	if err != nil {
		return 0, err
	}

	return score.Uint64(), nil
}

func (s *ValidatorAIService) TargetValidatorsCount() (uint64, error) {
	s.mu.RLock()
	block := s.blockNumber
	s.mu.RUnlock()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := s.contract.TargetValidatorsCount(&bind.CallOpts{
		Pending:     false,
		BlockNumber: new(big.Int).SetUint64(block),
		Context:     ctx,
	})
	if err != nil {
		return 0, err
	}

	return count.Uint64(), nil
}

func (s *ValidatorAIService) GetScore(pubkey [48]byte) (uint64, error) {
	s.mu.RLock()
	block := s.blockNumber
	if blockCache, ok := s.cache[block]; ok {
		key := crypto.Keccak256Hash(pubkey[:])
		if val, ok := blockCache[key]; ok {
			s.mu.RUnlock()
			return val, nil
		}
	}
	s.mu.RUnlock()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := crypto.Keccak256Hash(pubkey[:])
	scoreBig, err := s.contract.GetScore(&bind.CallOpts{
		Pending: false,
		//BlockNumber: new(big.Int).SetUint64(block),
		BlockNumber: nil,
		Context:     ctx,
	}, key)
	if err != nil {
		return 0, err
	}

	score := scoreBig.Uint64()

	s.mu.Lock()
	if _, ok := s.cache[block]; !ok {
		s.cache[block] = make(map[[32]byte]uint64)
	}
	s.cache[block][key] = score

	if len(s.cache) > maxBlocks {
		var oldest uint64 = math.MaxUint64
		for k := range s.cache {
			if k < oldest {
				oldest = k
			}
		}
		delete(s.cache, oldest)
	}
	s.mu.Unlock()

	return score, nil
}

func (s *ValidatorAIService) IsValidatorRegistered(pubkey [48]byte) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := crypto.Keccak256Hash(pubkey[:])

	res, err := s.contract.IsValidatorRegistered(&bind.CallOpts{
		Pending:     false,
		BlockNumber: new(big.Int).SetUint64(s.blockNumber),
		Context:     ctx,
	}, key)
	if err != nil {
		return false, err
	}

	return res, nil
}

func NewAIService(rpcURL string, contractAddr string) (*ValidatorAIService, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	contract, err := NewScoreCaller(common.HexToAddress(contractAddr), client)
	if err != nil {
		return nil, err
	}

	return &ValidatorAIService{
		ethClient: client,
		contract:  contract,
		cache:     make(map[uint64]map[[32]byte]uint64),
	}, nil
}
