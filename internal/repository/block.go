package repository

import (
	"context"

	"log"

	"github.com/ddr4869/ganache-explorer/ent"
	"github.com/ethereum/go-ethereum/core/types"
)

func (r Repository) CreateBlock(block *types.Block) (*ent.Block, error) {
	b, err := r.entClient.Block.Create().
		SetNumber(int(block.Number().Int64())).
		SetGasLimit(int(block.GasLimit())).
		SetGasUsed(int(block.GasUsed())).
		SetDifficulty(int(block.Difficulty().Int64())).
		SetTime(block.Time()).
		SetNumberU64(block.NumberU64()).
		SetMixDigest(block.MixDigest().Hex()).
		SetNonce(int(block.Nonce())).
		SetCoinbase(block.Coinbase().Hex()).
		SetRoot(block.Root().Hex()).
		SetParentHash(block.ParentHash().Hex()).
		SetTxHash(block.TxHash().Hex()).
		SetReceiptHash(block.ReceiptHash().Hex()).
		SetUncleHash(block.UncleHash().Hex()).
		SetTxCount(block.Transactions().Len()).
		Save(context.Background())
	if err != nil {
		log.Printf("Failed to save block: %v", err)
		return nil, err
	}
	return b, nil
}
