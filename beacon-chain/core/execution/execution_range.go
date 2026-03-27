package execution

import (
	"context"

	"github.com/OffchainLabs/prysm/v7/beacon-chain/db"
	"github.com/OffchainLabs/prysm/v7/config/params"
	"github.com/OffchainLabs/prysm/v7/consensus-types/primitives"
)

func GetEpochExecutionRange(
	ctx context.Context,
	beaconDB db.ReadOnlyDatabase,
	epoch primitives.Epoch,
) (*uint64, *uint64) {

	startSlot := primitives.Slot(uint64(epoch) * uint64(params.BeaconConfig().SlotsPerEpoch))
	endSlot := startSlot + primitives.Slot(params.BeaconConfig().SlotsPerEpoch-1)

	var firstBlock *uint64
	var lastBlock *uint64

	for slot := startSlot; slot <= endSlot; slot++ {
		blocks, err := beaconDB.BlocksBySlot(ctx, slot)
		if err != nil || len(blocks) == 0 {
			continue
		}

		for _, block := range blocks {
			payload, err := block.Block().Body().Execution()
			if err != nil || payload == nil {
				continue
			}

			bn := payload.BlockNumber()

			if firstBlock == nil {
				firstBlock = &bn
			}

			lastBlock = &bn
		}
	}

	return firstBlock, lastBlock
}
