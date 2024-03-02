package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("block_number").
			Positive(),
		field.Int("type"),
		field.Int("chain_id"),
		field.Int("nonce"),
		field.String("to"),
		field.Int("gas"),
		field.String("gasPrice"),
		field.String("gasTipCap"),
		field.String("gasFeeCap"),
		field.String("value"),
		field.Text("data").Optional(),
		field.String("hash"),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return nil
}
