package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Block holds the schema definition for the Block entity.
type Block struct {
	ent.Schema
}

// Fields of the Block.
func (Block) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number").
			Positive(),
		field.Int("gas_limit"),
		field.Int("gas_used"),
		field.Int("difficulty"),
		field.Uint64("time"),
		field.Uint64("number_u64"),
		field.String("mix_digest"),
		field.Int("nonce"),
		field.String("coinbase"),
		field.String("root"),
		field.String("parent_hash"),
		field.String("tx_hash"),
		field.String("receipt_hash"),
		field.String("uncle_hash"),
		field.Int("tx_count"),
	}
}

// Edges of the Block.
func (Block) Edges() []ent.Edge {
	return nil
}
