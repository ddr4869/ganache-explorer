package dto

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

// 144
type GetBlockHeaderByNumberResponse struct {
	BlockNumber int   `json:"block_number"`
	Block       Block `json:"block"`
}
