// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ddr4869/ether-go/ent/transaction"
)

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type int `json:"type,omitempty"`
	// ChainID holds the value of the "chain_id" field.
	ChainID int `json:"chain_id,omitempty"`
	// Nonce holds the value of the "nonce" field.
	Nonce int `json:"nonce,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
	// To holds the value of the "to" field.
	To string `json:"to,omitempty"`
	// Gas holds the value of the "gas" field.
	Gas int `json:"gas,omitempty"`
	// GasPrice holds the value of the "gasPrice" field.
	GasPrice string `json:"gasPrice,omitempty"`
	// GasTipCap holds the value of the "gasTipCap" field.
	GasTipCap string `json:"gasTipCap,omitempty"`
	// GasFeeCap holds the value of the "gasFeeCap" field.
	GasFeeCap string `json:"gasFeeCap,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Data holds the value of the "data" field.
	Data string `json:"data,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash         string `json:"hash,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID, transaction.FieldType, transaction.FieldChainID, transaction.FieldNonce, transaction.FieldGas:
			values[i] = new(sql.NullInt64)
		case transaction.FieldFrom, transaction.FieldTo, transaction.FieldGasPrice, transaction.FieldGasTipCap, transaction.FieldGasFeeCap, transaction.FieldValue, transaction.FieldData, transaction.FieldHash:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transaction fields.
func (t *Transaction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case transaction.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = int(value.Int64)
			}
		case transaction.FieldChainID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chain_id", values[i])
			} else if value.Valid {
				t.ChainID = int(value.Int64)
			}
		case transaction.FieldNonce:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nonce", values[i])
			} else if value.Valid {
				t.Nonce = int(value.Int64)
			}
		case transaction.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				t.From = value.String
			}
		case transaction.FieldTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				t.To = value.String
			}
		case transaction.FieldGas:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gas", values[i])
			} else if value.Valid {
				t.Gas = int(value.Int64)
			}
		case transaction.FieldGasPrice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gasPrice", values[i])
			} else if value.Valid {
				t.GasPrice = value.String
			}
		case transaction.FieldGasTipCap:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gasTipCap", values[i])
			} else if value.Valid {
				t.GasTipCap = value.String
			}
		case transaction.FieldGasFeeCap:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gasFeeCap", values[i])
			} else if value.Valid {
				t.GasFeeCap = value.String
			}
		case transaction.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				t.Value = value.String
			}
		case transaction.FieldData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value.Valid {
				t.Data = value.String
			}
		case transaction.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				t.Hash = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Transaction.
// This includes values selected through modifiers, order, etc.
func (t *Transaction) GetValue(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Transaction.
// Note that you need to call Transaction.Unwrap() before calling this method if this Transaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transaction) Update() *TransactionUpdateOne {
	return NewTransactionClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Transaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transaction) Unwrap() *Transaction {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transaction is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteString(", ")
	builder.WriteString("chain_id=")
	builder.WriteString(fmt.Sprintf("%v", t.ChainID))
	builder.WriteString(", ")
	builder.WriteString("nonce=")
	builder.WriteString(fmt.Sprintf("%v", t.Nonce))
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(t.From)
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(t.To)
	builder.WriteString(", ")
	builder.WriteString("gas=")
	builder.WriteString(fmt.Sprintf("%v", t.Gas))
	builder.WriteString(", ")
	builder.WriteString("gasPrice=")
	builder.WriteString(t.GasPrice)
	builder.WriteString(", ")
	builder.WriteString("gasTipCap=")
	builder.WriteString(t.GasTipCap)
	builder.WriteString(", ")
	builder.WriteString("gasFeeCap=")
	builder.WriteString(t.GasFeeCap)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(t.Value)
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(t.Data)
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(t.Hash)
	builder.WriteByte(')')
	return builder.String()
}

// Transactions is a parsable slice of Transaction.
type Transactions []*Transaction
