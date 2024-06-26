// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ddr4869/ganache-explorer/ent/block"
)

// BlockCreate is the builder for creating a Block entity.
type BlockCreate struct {
	config
	mutation *BlockMutation
	hooks    []Hook
}

// SetNumber sets the "number" field.
func (bc *BlockCreate) SetNumber(i int) *BlockCreate {
	bc.mutation.SetNumber(i)
	return bc
}

// SetGasLimit sets the "gas_limit" field.
func (bc *BlockCreate) SetGasLimit(i int) *BlockCreate {
	bc.mutation.SetGasLimit(i)
	return bc
}

// SetGasUsed sets the "gas_used" field.
func (bc *BlockCreate) SetGasUsed(i int) *BlockCreate {
	bc.mutation.SetGasUsed(i)
	return bc
}

// SetDifficulty sets the "difficulty" field.
func (bc *BlockCreate) SetDifficulty(i int) *BlockCreate {
	bc.mutation.SetDifficulty(i)
	return bc
}

// SetTime sets the "time" field.
func (bc *BlockCreate) SetTime(u uint64) *BlockCreate {
	bc.mutation.SetTime(u)
	return bc
}

// SetNumberU64 sets the "number_u64" field.
func (bc *BlockCreate) SetNumberU64(u uint64) *BlockCreate {
	bc.mutation.SetNumberU64(u)
	return bc
}

// SetMixDigest sets the "mix_digest" field.
func (bc *BlockCreate) SetMixDigest(s string) *BlockCreate {
	bc.mutation.SetMixDigest(s)
	return bc
}

// SetNonce sets the "nonce" field.
func (bc *BlockCreate) SetNonce(i int) *BlockCreate {
	bc.mutation.SetNonce(i)
	return bc
}

// SetCoinbase sets the "coinbase" field.
func (bc *BlockCreate) SetCoinbase(s string) *BlockCreate {
	bc.mutation.SetCoinbase(s)
	return bc
}

// SetRoot sets the "root" field.
func (bc *BlockCreate) SetRoot(s string) *BlockCreate {
	bc.mutation.SetRoot(s)
	return bc
}

// SetParentHash sets the "parent_hash" field.
func (bc *BlockCreate) SetParentHash(s string) *BlockCreate {
	bc.mutation.SetParentHash(s)
	return bc
}

// SetTxHash sets the "tx_hash" field.
func (bc *BlockCreate) SetTxHash(s string) *BlockCreate {
	bc.mutation.SetTxHash(s)
	return bc
}

// SetReceiptHash sets the "receipt_hash" field.
func (bc *BlockCreate) SetReceiptHash(s string) *BlockCreate {
	bc.mutation.SetReceiptHash(s)
	return bc
}

// SetUncleHash sets the "uncle_hash" field.
func (bc *BlockCreate) SetUncleHash(s string) *BlockCreate {
	bc.mutation.SetUncleHash(s)
	return bc
}

// SetTxCount sets the "tx_count" field.
func (bc *BlockCreate) SetTxCount(i int) *BlockCreate {
	bc.mutation.SetTxCount(i)
	return bc
}

// Mutation returns the BlockMutation object of the builder.
func (bc *BlockCreate) Mutation() *BlockMutation {
	return bc.mutation
}

// Save creates the Block in the database.
func (bc *BlockCreate) Save(ctx context.Context) (*Block, error) {
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlockCreate) SaveX(ctx context.Context) *Block {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlockCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlockCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlockCreate) check() error {
	if _, ok := bc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Block.number"`)}
	}
	if v, ok := bc.mutation.Number(); ok {
		if err := block.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Block.number": %w`, err)}
		}
	}
	if _, ok := bc.mutation.GasLimit(); !ok {
		return &ValidationError{Name: "gas_limit", err: errors.New(`ent: missing required field "Block.gas_limit"`)}
	}
	if _, ok := bc.mutation.GasUsed(); !ok {
		return &ValidationError{Name: "gas_used", err: errors.New(`ent: missing required field "Block.gas_used"`)}
	}
	if _, ok := bc.mutation.Difficulty(); !ok {
		return &ValidationError{Name: "difficulty", err: errors.New(`ent: missing required field "Block.difficulty"`)}
	}
	if _, ok := bc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "Block.time"`)}
	}
	if _, ok := bc.mutation.NumberU64(); !ok {
		return &ValidationError{Name: "number_u64", err: errors.New(`ent: missing required field "Block.number_u64"`)}
	}
	if _, ok := bc.mutation.MixDigest(); !ok {
		return &ValidationError{Name: "mix_digest", err: errors.New(`ent: missing required field "Block.mix_digest"`)}
	}
	if _, ok := bc.mutation.Nonce(); !ok {
		return &ValidationError{Name: "nonce", err: errors.New(`ent: missing required field "Block.nonce"`)}
	}
	if _, ok := bc.mutation.Coinbase(); !ok {
		return &ValidationError{Name: "coinbase", err: errors.New(`ent: missing required field "Block.coinbase"`)}
	}
	if _, ok := bc.mutation.Root(); !ok {
		return &ValidationError{Name: "root", err: errors.New(`ent: missing required field "Block.root"`)}
	}
	if _, ok := bc.mutation.ParentHash(); !ok {
		return &ValidationError{Name: "parent_hash", err: errors.New(`ent: missing required field "Block.parent_hash"`)}
	}
	if _, ok := bc.mutation.TxHash(); !ok {
		return &ValidationError{Name: "tx_hash", err: errors.New(`ent: missing required field "Block.tx_hash"`)}
	}
	if _, ok := bc.mutation.ReceiptHash(); !ok {
		return &ValidationError{Name: "receipt_hash", err: errors.New(`ent: missing required field "Block.receipt_hash"`)}
	}
	if _, ok := bc.mutation.UncleHash(); !ok {
		return &ValidationError{Name: "uncle_hash", err: errors.New(`ent: missing required field "Block.uncle_hash"`)}
	}
	if _, ok := bc.mutation.TxCount(); !ok {
		return &ValidationError{Name: "tx_count", err: errors.New(`ent: missing required field "Block.tx_count"`)}
	}
	return nil
}

func (bc *BlockCreate) sqlSave(ctx context.Context) (*Block, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BlockCreate) createSpec() (*Block, *sqlgraph.CreateSpec) {
	var (
		_node = &Block{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(block.Table, sqlgraph.NewFieldSpec(block.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.Number(); ok {
		_spec.SetField(block.FieldNumber, field.TypeInt, value)
		_node.Number = value
	}
	if value, ok := bc.mutation.GasLimit(); ok {
		_spec.SetField(block.FieldGasLimit, field.TypeInt, value)
		_node.GasLimit = value
	}
	if value, ok := bc.mutation.GasUsed(); ok {
		_spec.SetField(block.FieldGasUsed, field.TypeInt, value)
		_node.GasUsed = value
	}
	if value, ok := bc.mutation.Difficulty(); ok {
		_spec.SetField(block.FieldDifficulty, field.TypeInt, value)
		_node.Difficulty = value
	}
	if value, ok := bc.mutation.Time(); ok {
		_spec.SetField(block.FieldTime, field.TypeUint64, value)
		_node.Time = value
	}
	if value, ok := bc.mutation.NumberU64(); ok {
		_spec.SetField(block.FieldNumberU64, field.TypeUint64, value)
		_node.NumberU64 = value
	}
	if value, ok := bc.mutation.MixDigest(); ok {
		_spec.SetField(block.FieldMixDigest, field.TypeString, value)
		_node.MixDigest = value
	}
	if value, ok := bc.mutation.Nonce(); ok {
		_spec.SetField(block.FieldNonce, field.TypeInt, value)
		_node.Nonce = value
	}
	if value, ok := bc.mutation.Coinbase(); ok {
		_spec.SetField(block.FieldCoinbase, field.TypeString, value)
		_node.Coinbase = value
	}
	if value, ok := bc.mutation.Root(); ok {
		_spec.SetField(block.FieldRoot, field.TypeString, value)
		_node.Root = value
	}
	if value, ok := bc.mutation.ParentHash(); ok {
		_spec.SetField(block.FieldParentHash, field.TypeString, value)
		_node.ParentHash = value
	}
	if value, ok := bc.mutation.TxHash(); ok {
		_spec.SetField(block.FieldTxHash, field.TypeString, value)
		_node.TxHash = value
	}
	if value, ok := bc.mutation.ReceiptHash(); ok {
		_spec.SetField(block.FieldReceiptHash, field.TypeString, value)
		_node.ReceiptHash = value
	}
	if value, ok := bc.mutation.UncleHash(); ok {
		_spec.SetField(block.FieldUncleHash, field.TypeString, value)
		_node.UncleHash = value
	}
	if value, ok := bc.mutation.TxCount(); ok {
		_spec.SetField(block.FieldTxCount, field.TypeInt, value)
		_node.TxCount = value
	}
	return _node, _spec
}

// BlockCreateBulk is the builder for creating many Block entities in bulk.
type BlockCreateBulk struct {
	config
	err      error
	builders []*BlockCreate
}

// Save creates the Block entities in the database.
func (bcb *BlockCreateBulk) Save(ctx context.Context) ([]*Block, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Block, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlockMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlockCreateBulk) SaveX(ctx context.Context) []*Block {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlockCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlockCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
