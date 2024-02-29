package dto

// func (b *Block) Number() *big.Int     { return new(big.Int).Set(b.header.Number) }
// func (b *Block) GasLimit() uint64     { return b.header.GasLimit }
// func (b *Block) GasUsed() uint64      { return b.header.GasUsed }
// func (b *Block) Difficulty() *big.Int { return new(big.Int).Set(b.header.Difficulty) }
// func (b *Block) Time() uint64         { return b.header.Time }

// func (b *Block) NumberU64() uint64        { return b.header.Number.Uint64() }
// func (b *Block) MixDigest() common.Hash   { return b.header.MixDigest }
// func (b *Block) Nonce() uint64            { return binary.BigEndian.Uint64(b.header.Nonce[:]) }
// func (b *Block) Bloom() Bloom             { return b.header.Bloom }
// func (b *Block) Coinbase() common.Address { return b.header.Coinbase }
// func (b *Block) Root() common.Hash        { return b.header.Root }
// func (b *Block) ParentHash() common.Hash  { return b.header.ParentHash }
// func (b *Block) TxHash() common.Hash      { return b.header.TxHash }
// func (b *Block) ReceiptHash() common.Hash { return b.header.ReceiptHash }
// func (b *Block) UncleHash() common.Hash   { return b.header.UncleHash }
// func (b *Block) Extra() []byte            { return common.CopyBytes(b.header.Extra) }

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
