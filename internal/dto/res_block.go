package dto

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type Block struct {
	Number      uint64 `json:"number"`
	GasLimit    uint64 `json:"gas_limit"`
	GasUsed     uint64 `json:"gas_used"`
	Difficulty  uint64 `json:"difficulty"`
	Time        uint64 `json:"time"`
	NumberU64   uint64 `json:"number_u64"`
	MixDigest   string `json:"mix_digest"`
	Nonce       uint64 `json:"nonce"`
	Coinbase    string `json:"coinbase"`
	Root        string `json:"root"`
	ParentHash  string `json:"parent_hash"`
	TxHash      string `json:"tx_hash"`
	ReceiptHash string `json:"receipt_hash"`
	UncleHash   string `json:"uncle_hash"`

	TransactionCount int `json:"tx_count"`
}

type GetBlockHeaderByNumberResponse struct {
	Block Block `json:"block"`
}

func ConvertBlockToResponse(block types.Block) GetBlockHeaderByNumberResponse {
	return GetBlockHeaderByNumberResponse{
		Block: Block{
			Number:      block.Number().Uint64(),
			GasLimit:    block.GasLimit(),
			GasUsed:     block.GasUsed(),
			Time:        block.Time(),
			Difficulty:  block.Difficulty().Uint64(),
			NumberU64:   block.NumberU64(),
			MixDigest:   block.MixDigest().Hex(),
			Nonce:       block.Nonce(),
			Coinbase:    block.Coinbase().Hex(),
			Root:        block.Root().Hex(),
			ParentHash:  block.ParentHash().Hex(),
			TxHash:      block.TxHash().Hex(),
			ReceiptHash: block.ReceiptHash().Hex(),
			UncleHash:   block.UncleHash().Hex(),

			TransactionCount: len(block.Transactions()),
		},
	}
}
