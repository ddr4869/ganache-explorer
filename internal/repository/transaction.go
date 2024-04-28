package repository

import (
	"context"
	"log"

	"github.com/ddr4869/ganache-explorer/ent"
	"github.com/ethereum/go-ethereum/core/types"
)

func (r Repository) CreateTranasction(block_number int, tx *types.Transaction) (*ent.Transaction, error) {
	query := r.entClient.Transaction.Create().
		SetBlockNumber(block_number).
		SetType(int(tx.Type())).
		SetChainID(int(tx.ChainId().Int64())).
		SetNonce(int(tx.Nonce())).
		SetGas(int(tx.Gas())).
		SetGasPrice(tx.GasPrice().String()).
		SetGasTipCap(tx.GasTipCap().String()).
		SetGasFeeCap(tx.GasFeeCap().String()).
		SetValue(tx.Value().String()).
		//SetData(string(tx.Data())).
		SetHash(tx.Hash().Hex())
	if tx.To() == nil {
		query.SetTo("")
	} else {
		query.SetTo(tx.To().Hex())
	}
	t, err := query.Save(context.Background())
	if err != nil {
		log.Printf("Failed to save transaction: %v", err)
		return nil, err
	}
	return t, nil
}
