// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ddr4869/ether-go/ent/block"
	"github.com/ddr4869/ether-go/ent/predicate"
)

// BlockUpdate is the builder for updating Block entities.
type BlockUpdate struct {
	config
	hooks    []Hook
	mutation *BlockMutation
}

// Where appends a list predicates to the BlockUpdate builder.
func (bu *BlockUpdate) Where(ps ...predicate.Block) *BlockUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetNumber sets the "number" field.
func (bu *BlockUpdate) SetNumber(i int) *BlockUpdate {
	bu.mutation.ResetNumber()
	bu.mutation.SetNumber(i)
	return bu
}

// AddNumber adds i to the "number" field.
func (bu *BlockUpdate) AddNumber(i int) *BlockUpdate {
	bu.mutation.AddNumber(i)
	return bu
}

// SetGasLimit sets the "gas_limit" field.
func (bu *BlockUpdate) SetGasLimit(i int) *BlockUpdate {
	bu.mutation.ResetGasLimit()
	bu.mutation.SetGasLimit(i)
	return bu
}

// AddGasLimit adds i to the "gas_limit" field.
func (bu *BlockUpdate) AddGasLimit(i int) *BlockUpdate {
	bu.mutation.AddGasLimit(i)
	return bu
}

// SetGasUsed sets the "gas_used" field.
func (bu *BlockUpdate) SetGasUsed(i int) *BlockUpdate {
	bu.mutation.ResetGasUsed()
	bu.mutation.SetGasUsed(i)
	return bu
}

// AddGasUsed adds i to the "gas_used" field.
func (bu *BlockUpdate) AddGasUsed(i int) *BlockUpdate {
	bu.mutation.AddGasUsed(i)
	return bu
}

// SetDifficulty sets the "difficulty" field.
func (bu *BlockUpdate) SetDifficulty(i int) *BlockUpdate {
	bu.mutation.ResetDifficulty()
	bu.mutation.SetDifficulty(i)
	return bu
}

// AddDifficulty adds i to the "difficulty" field.
func (bu *BlockUpdate) AddDifficulty(i int) *BlockUpdate {
	bu.mutation.AddDifficulty(i)
	return bu
}

// SetTime sets the "time" field.
func (bu *BlockUpdate) SetTime(t time.Time) *BlockUpdate {
	bu.mutation.SetTime(t)
	return bu
}

// SetNumberU64 sets the "number_u64" field.
func (bu *BlockUpdate) SetNumberU64(u uint64) *BlockUpdate {
	bu.mutation.ResetNumberU64()
	bu.mutation.SetNumberU64(u)
	return bu
}

// AddNumberU64 adds u to the "number_u64" field.
func (bu *BlockUpdate) AddNumberU64(u int64) *BlockUpdate {
	bu.mutation.AddNumberU64(u)
	return bu
}

// SetMixDigest sets the "mix_digest" field.
func (bu *BlockUpdate) SetMixDigest(s string) *BlockUpdate {
	bu.mutation.SetMixDigest(s)
	return bu
}

// SetNonce sets the "nonce" field.
func (bu *BlockUpdate) SetNonce(i int) *BlockUpdate {
	bu.mutation.ResetNonce()
	bu.mutation.SetNonce(i)
	return bu
}

// AddNonce adds i to the "nonce" field.
func (bu *BlockUpdate) AddNonce(i int) *BlockUpdate {
	bu.mutation.AddNonce(i)
	return bu
}

// SetCoinbase sets the "coinbase" field.
func (bu *BlockUpdate) SetCoinbase(s string) *BlockUpdate {
	bu.mutation.SetCoinbase(s)
	return bu
}

// SetRoot sets the "root" field.
func (bu *BlockUpdate) SetRoot(s string) *BlockUpdate {
	bu.mutation.SetRoot(s)
	return bu
}

// SetParentHash sets the "parent_hash" field.
func (bu *BlockUpdate) SetParentHash(s string) *BlockUpdate {
	bu.mutation.SetParentHash(s)
	return bu
}

// SetTxHash sets the "tx_hash" field.
func (bu *BlockUpdate) SetTxHash(s string) *BlockUpdate {
	bu.mutation.SetTxHash(s)
	return bu
}

// SetReceiptHash sets the "receipt_hash" field.
func (bu *BlockUpdate) SetReceiptHash(s string) *BlockUpdate {
	bu.mutation.SetReceiptHash(s)
	return bu
}

// SetUncleHash sets the "uncle_hash" field.
func (bu *BlockUpdate) SetUncleHash(s string) *BlockUpdate {
	bu.mutation.SetUncleHash(s)
	return bu
}

// SetTxCount sets the "tx_count" field.
func (bu *BlockUpdate) SetTxCount(i int) *BlockUpdate {
	bu.mutation.ResetTxCount()
	bu.mutation.SetTxCount(i)
	return bu
}

// AddTxCount adds i to the "tx_count" field.
func (bu *BlockUpdate) AddTxCount(i int) *BlockUpdate {
	bu.mutation.AddTxCount(i)
	return bu
}

// Mutation returns the BlockMutation object of the builder.
func (bu *BlockUpdate) Mutation() *BlockMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlockUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, BlockMutation](ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlockUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlockUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlockUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BlockUpdate) check() error {
	if v, ok := bu.mutation.Number(); ok {
		if err := block.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Block.number": %w`, err)}
		}
	}
	return nil
}

func (bu *BlockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(block.Table, block.Columns, sqlgraph.NewFieldSpec(block.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Number(); ok {
		_spec.SetField(block.FieldNumber, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedNumber(); ok {
		_spec.AddField(block.FieldNumber, field.TypeInt, value)
	}
	if value, ok := bu.mutation.GasLimit(); ok {
		_spec.SetField(block.FieldGasLimit, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedGasLimit(); ok {
		_spec.AddField(block.FieldGasLimit, field.TypeInt, value)
	}
	if value, ok := bu.mutation.GasUsed(); ok {
		_spec.SetField(block.FieldGasUsed, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedGasUsed(); ok {
		_spec.AddField(block.FieldGasUsed, field.TypeInt, value)
	}
	if value, ok := bu.mutation.Difficulty(); ok {
		_spec.SetField(block.FieldDifficulty, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedDifficulty(); ok {
		_spec.AddField(block.FieldDifficulty, field.TypeInt, value)
	}
	if value, ok := bu.mutation.Time(); ok {
		_spec.SetField(block.FieldTime, field.TypeTime, value)
	}
	if value, ok := bu.mutation.NumberU64(); ok {
		_spec.SetField(block.FieldNumberU64, field.TypeUint64, value)
	}
	if value, ok := bu.mutation.AddedNumberU64(); ok {
		_spec.AddField(block.FieldNumberU64, field.TypeUint64, value)
	}
	if value, ok := bu.mutation.MixDigest(); ok {
		_spec.SetField(block.FieldMixDigest, field.TypeString, value)
	}
	if value, ok := bu.mutation.Nonce(); ok {
		_spec.SetField(block.FieldNonce, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedNonce(); ok {
		_spec.AddField(block.FieldNonce, field.TypeInt, value)
	}
	if value, ok := bu.mutation.Coinbase(); ok {
		_spec.SetField(block.FieldCoinbase, field.TypeString, value)
	}
	if value, ok := bu.mutation.Root(); ok {
		_spec.SetField(block.FieldRoot, field.TypeString, value)
	}
	if value, ok := bu.mutation.ParentHash(); ok {
		_spec.SetField(block.FieldParentHash, field.TypeString, value)
	}
	if value, ok := bu.mutation.TxHash(); ok {
		_spec.SetField(block.FieldTxHash, field.TypeString, value)
	}
	if value, ok := bu.mutation.ReceiptHash(); ok {
		_spec.SetField(block.FieldReceiptHash, field.TypeString, value)
	}
	if value, ok := bu.mutation.UncleHash(); ok {
		_spec.SetField(block.FieldUncleHash, field.TypeString, value)
	}
	if value, ok := bu.mutation.TxCount(); ok {
		_spec.SetField(block.FieldTxCount, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedTxCount(); ok {
		_spec.AddField(block.FieldTxCount, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{block.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BlockUpdateOne is the builder for updating a single Block entity.
type BlockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlockMutation
}

// SetNumber sets the "number" field.
func (buo *BlockUpdateOne) SetNumber(i int) *BlockUpdateOne {
	buo.mutation.ResetNumber()
	buo.mutation.SetNumber(i)
	return buo
}

// AddNumber adds i to the "number" field.
func (buo *BlockUpdateOne) AddNumber(i int) *BlockUpdateOne {
	buo.mutation.AddNumber(i)
	return buo
}

// SetGasLimit sets the "gas_limit" field.
func (buo *BlockUpdateOne) SetGasLimit(i int) *BlockUpdateOne {
	buo.mutation.ResetGasLimit()
	buo.mutation.SetGasLimit(i)
	return buo
}

// AddGasLimit adds i to the "gas_limit" field.
func (buo *BlockUpdateOne) AddGasLimit(i int) *BlockUpdateOne {
	buo.mutation.AddGasLimit(i)
	return buo
}

// SetGasUsed sets the "gas_used" field.
func (buo *BlockUpdateOne) SetGasUsed(i int) *BlockUpdateOne {
	buo.mutation.ResetGasUsed()
	buo.mutation.SetGasUsed(i)
	return buo
}

// AddGasUsed adds i to the "gas_used" field.
func (buo *BlockUpdateOne) AddGasUsed(i int) *BlockUpdateOne {
	buo.mutation.AddGasUsed(i)
	return buo
}

// SetDifficulty sets the "difficulty" field.
func (buo *BlockUpdateOne) SetDifficulty(i int) *BlockUpdateOne {
	buo.mutation.ResetDifficulty()
	buo.mutation.SetDifficulty(i)
	return buo
}

// AddDifficulty adds i to the "difficulty" field.
func (buo *BlockUpdateOne) AddDifficulty(i int) *BlockUpdateOne {
	buo.mutation.AddDifficulty(i)
	return buo
}

// SetTime sets the "time" field.
func (buo *BlockUpdateOne) SetTime(t time.Time) *BlockUpdateOne {
	buo.mutation.SetTime(t)
	return buo
}

// SetNumberU64 sets the "number_u64" field.
func (buo *BlockUpdateOne) SetNumberU64(u uint64) *BlockUpdateOne {
	buo.mutation.ResetNumberU64()
	buo.mutation.SetNumberU64(u)
	return buo
}

// AddNumberU64 adds u to the "number_u64" field.
func (buo *BlockUpdateOne) AddNumberU64(u int64) *BlockUpdateOne {
	buo.mutation.AddNumberU64(u)
	return buo
}

// SetMixDigest sets the "mix_digest" field.
func (buo *BlockUpdateOne) SetMixDigest(s string) *BlockUpdateOne {
	buo.mutation.SetMixDigest(s)
	return buo
}

// SetNonce sets the "nonce" field.
func (buo *BlockUpdateOne) SetNonce(i int) *BlockUpdateOne {
	buo.mutation.ResetNonce()
	buo.mutation.SetNonce(i)
	return buo
}

// AddNonce adds i to the "nonce" field.
func (buo *BlockUpdateOne) AddNonce(i int) *BlockUpdateOne {
	buo.mutation.AddNonce(i)
	return buo
}

// SetCoinbase sets the "coinbase" field.
func (buo *BlockUpdateOne) SetCoinbase(s string) *BlockUpdateOne {
	buo.mutation.SetCoinbase(s)
	return buo
}

// SetRoot sets the "root" field.
func (buo *BlockUpdateOne) SetRoot(s string) *BlockUpdateOne {
	buo.mutation.SetRoot(s)
	return buo
}

// SetParentHash sets the "parent_hash" field.
func (buo *BlockUpdateOne) SetParentHash(s string) *BlockUpdateOne {
	buo.mutation.SetParentHash(s)
	return buo
}

// SetTxHash sets the "tx_hash" field.
func (buo *BlockUpdateOne) SetTxHash(s string) *BlockUpdateOne {
	buo.mutation.SetTxHash(s)
	return buo
}

// SetReceiptHash sets the "receipt_hash" field.
func (buo *BlockUpdateOne) SetReceiptHash(s string) *BlockUpdateOne {
	buo.mutation.SetReceiptHash(s)
	return buo
}

// SetUncleHash sets the "uncle_hash" field.
func (buo *BlockUpdateOne) SetUncleHash(s string) *BlockUpdateOne {
	buo.mutation.SetUncleHash(s)
	return buo
}

// SetTxCount sets the "tx_count" field.
func (buo *BlockUpdateOne) SetTxCount(i int) *BlockUpdateOne {
	buo.mutation.ResetTxCount()
	buo.mutation.SetTxCount(i)
	return buo
}

// AddTxCount adds i to the "tx_count" field.
func (buo *BlockUpdateOne) AddTxCount(i int) *BlockUpdateOne {
	buo.mutation.AddTxCount(i)
	return buo
}

// Mutation returns the BlockMutation object of the builder.
func (buo *BlockUpdateOne) Mutation() *BlockMutation {
	return buo.mutation
}

// Where appends a list predicates to the BlockUpdate builder.
func (buo *BlockUpdateOne) Where(ps ...predicate.Block) *BlockUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlockUpdateOne) Select(field string, fields ...string) *BlockUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Block entity.
func (buo *BlockUpdateOne) Save(ctx context.Context) (*Block, error) {
	return withHooks[*Block, BlockMutation](ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlockUpdateOne) SaveX(ctx context.Context) *Block {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlockUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlockUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BlockUpdateOne) check() error {
	if v, ok := buo.mutation.Number(); ok {
		if err := block.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Block.number": %w`, err)}
		}
	}
	return nil
}

func (buo *BlockUpdateOne) sqlSave(ctx context.Context) (_node *Block, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(block.Table, block.Columns, sqlgraph.NewFieldSpec(block.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Block.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, block.FieldID)
		for _, f := range fields {
			if !block.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != block.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Number(); ok {
		_spec.SetField(block.FieldNumber, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedNumber(); ok {
		_spec.AddField(block.FieldNumber, field.TypeInt, value)
	}
	if value, ok := buo.mutation.GasLimit(); ok {
		_spec.SetField(block.FieldGasLimit, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedGasLimit(); ok {
		_spec.AddField(block.FieldGasLimit, field.TypeInt, value)
	}
	if value, ok := buo.mutation.GasUsed(); ok {
		_spec.SetField(block.FieldGasUsed, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedGasUsed(); ok {
		_spec.AddField(block.FieldGasUsed, field.TypeInt, value)
	}
	if value, ok := buo.mutation.Difficulty(); ok {
		_spec.SetField(block.FieldDifficulty, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedDifficulty(); ok {
		_spec.AddField(block.FieldDifficulty, field.TypeInt, value)
	}
	if value, ok := buo.mutation.Time(); ok {
		_spec.SetField(block.FieldTime, field.TypeTime, value)
	}
	if value, ok := buo.mutation.NumberU64(); ok {
		_spec.SetField(block.FieldNumberU64, field.TypeUint64, value)
	}
	if value, ok := buo.mutation.AddedNumberU64(); ok {
		_spec.AddField(block.FieldNumberU64, field.TypeUint64, value)
	}
	if value, ok := buo.mutation.MixDigest(); ok {
		_spec.SetField(block.FieldMixDigest, field.TypeString, value)
	}
	if value, ok := buo.mutation.Nonce(); ok {
		_spec.SetField(block.FieldNonce, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedNonce(); ok {
		_spec.AddField(block.FieldNonce, field.TypeInt, value)
	}
	if value, ok := buo.mutation.Coinbase(); ok {
		_spec.SetField(block.FieldCoinbase, field.TypeString, value)
	}
	if value, ok := buo.mutation.Root(); ok {
		_spec.SetField(block.FieldRoot, field.TypeString, value)
	}
	if value, ok := buo.mutation.ParentHash(); ok {
		_spec.SetField(block.FieldParentHash, field.TypeString, value)
	}
	if value, ok := buo.mutation.TxHash(); ok {
		_spec.SetField(block.FieldTxHash, field.TypeString, value)
	}
	if value, ok := buo.mutation.ReceiptHash(); ok {
		_spec.SetField(block.FieldReceiptHash, field.TypeString, value)
	}
	if value, ok := buo.mutation.UncleHash(); ok {
		_spec.SetField(block.FieldUncleHash, field.TypeString, value)
	}
	if value, ok := buo.mutation.TxCount(); ok {
		_spec.SetField(block.FieldTxCount, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedTxCount(); ok {
		_spec.AddField(block.FieldTxCount, field.TypeInt, value)
	}
	_node = &Block{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{block.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}