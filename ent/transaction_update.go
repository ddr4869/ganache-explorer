// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ddr4869/ganache-explorer/ent/predicate"
	"github.com/ddr4869/ganache-explorer/ent/transaction"
)

// TransactionUpdate is the builder for updating Transaction entities.
type TransactionUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionMutation
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tu *TransactionUpdate) Where(ps ...predicate.Transaction) *TransactionUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetBlockNumber sets the "block_number" field.
func (tu *TransactionUpdate) SetBlockNumber(i int) *TransactionUpdate {
	tu.mutation.ResetBlockNumber()
	tu.mutation.SetBlockNumber(i)
	return tu
}

// SetNillableBlockNumber sets the "block_number" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableBlockNumber(i *int) *TransactionUpdate {
	if i != nil {
		tu.SetBlockNumber(*i)
	}
	return tu
}

// AddBlockNumber adds i to the "block_number" field.
func (tu *TransactionUpdate) AddBlockNumber(i int) *TransactionUpdate {
	tu.mutation.AddBlockNumber(i)
	return tu
}

// SetType sets the "type" field.
func (tu *TransactionUpdate) SetType(i int) *TransactionUpdate {
	tu.mutation.ResetType()
	tu.mutation.SetType(i)
	return tu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableType(i *int) *TransactionUpdate {
	if i != nil {
		tu.SetType(*i)
	}
	return tu
}

// AddType adds i to the "type" field.
func (tu *TransactionUpdate) AddType(i int) *TransactionUpdate {
	tu.mutation.AddType(i)
	return tu
}

// SetChainID sets the "chain_id" field.
func (tu *TransactionUpdate) SetChainID(i int) *TransactionUpdate {
	tu.mutation.ResetChainID()
	tu.mutation.SetChainID(i)
	return tu
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableChainID(i *int) *TransactionUpdate {
	if i != nil {
		tu.SetChainID(*i)
	}
	return tu
}

// AddChainID adds i to the "chain_id" field.
func (tu *TransactionUpdate) AddChainID(i int) *TransactionUpdate {
	tu.mutation.AddChainID(i)
	return tu
}

// SetNonce sets the "nonce" field.
func (tu *TransactionUpdate) SetNonce(i int) *TransactionUpdate {
	tu.mutation.ResetNonce()
	tu.mutation.SetNonce(i)
	return tu
}

// SetNillableNonce sets the "nonce" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableNonce(i *int) *TransactionUpdate {
	if i != nil {
		tu.SetNonce(*i)
	}
	return tu
}

// AddNonce adds i to the "nonce" field.
func (tu *TransactionUpdate) AddNonce(i int) *TransactionUpdate {
	tu.mutation.AddNonce(i)
	return tu
}

// SetTo sets the "to" field.
func (tu *TransactionUpdate) SetTo(s string) *TransactionUpdate {
	tu.mutation.SetTo(s)
	return tu
}

// SetNillableTo sets the "to" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableTo(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetTo(*s)
	}
	return tu
}

// SetGas sets the "gas" field.
func (tu *TransactionUpdate) SetGas(i int) *TransactionUpdate {
	tu.mutation.ResetGas()
	tu.mutation.SetGas(i)
	return tu
}

// SetNillableGas sets the "gas" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableGas(i *int) *TransactionUpdate {
	if i != nil {
		tu.SetGas(*i)
	}
	return tu
}

// AddGas adds i to the "gas" field.
func (tu *TransactionUpdate) AddGas(i int) *TransactionUpdate {
	tu.mutation.AddGas(i)
	return tu
}

// SetGasPrice sets the "gasPrice" field.
func (tu *TransactionUpdate) SetGasPrice(s string) *TransactionUpdate {
	tu.mutation.SetGasPrice(s)
	return tu
}

// SetNillableGasPrice sets the "gasPrice" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableGasPrice(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetGasPrice(*s)
	}
	return tu
}

// SetGasTipCap sets the "gasTipCap" field.
func (tu *TransactionUpdate) SetGasTipCap(s string) *TransactionUpdate {
	tu.mutation.SetGasTipCap(s)
	return tu
}

// SetNillableGasTipCap sets the "gasTipCap" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableGasTipCap(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetGasTipCap(*s)
	}
	return tu
}

// SetGasFeeCap sets the "gasFeeCap" field.
func (tu *TransactionUpdate) SetGasFeeCap(s string) *TransactionUpdate {
	tu.mutation.SetGasFeeCap(s)
	return tu
}

// SetNillableGasFeeCap sets the "gasFeeCap" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableGasFeeCap(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetGasFeeCap(*s)
	}
	return tu
}

// SetValue sets the "value" field.
func (tu *TransactionUpdate) SetValue(s string) *TransactionUpdate {
	tu.mutation.SetValue(s)
	return tu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableValue(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetValue(*s)
	}
	return tu
}

// SetData sets the "data" field.
func (tu *TransactionUpdate) SetData(s string) *TransactionUpdate {
	tu.mutation.SetData(s)
	return tu
}

// SetNillableData sets the "data" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableData(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetData(*s)
	}
	return tu
}

// ClearData clears the value of the "data" field.
func (tu *TransactionUpdate) ClearData() *TransactionUpdate {
	tu.mutation.ClearData()
	return tu
}

// SetHash sets the "hash" field.
func (tu *TransactionUpdate) SetHash(s string) *TransactionUpdate {
	tu.mutation.SetHash(s)
	return tu
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableHash(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetHash(*s)
	}
	return tu
}

// Mutation returns the TransactionMutation object of the builder.
func (tu *TransactionUpdate) Mutation() *TransactionMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TransactionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TransactionUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TransactionUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TransactionUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TransactionUpdate) check() error {
	if v, ok := tu.mutation.BlockNumber(); ok {
		if err := transaction.BlockNumberValidator(v); err != nil {
			return &ValidationError{Name: "block_number", err: fmt.Errorf(`ent: validator failed for field "Transaction.block_number": %w`, err)}
		}
	}
	return nil
}

func (tu *TransactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(transaction.Table, transaction.Columns, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.BlockNumber(); ok {
		_spec.SetField(transaction.FieldBlockNumber, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedBlockNumber(); ok {
		_spec.AddField(transaction.FieldBlockNumber, field.TypeInt, value)
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.SetField(transaction.FieldType, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedType(); ok {
		_spec.AddField(transaction.FieldType, field.TypeInt, value)
	}
	if value, ok := tu.mutation.ChainID(); ok {
		_spec.SetField(transaction.FieldChainID, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedChainID(); ok {
		_spec.AddField(transaction.FieldChainID, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Nonce(); ok {
		_spec.SetField(transaction.FieldNonce, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedNonce(); ok {
		_spec.AddField(transaction.FieldNonce, field.TypeInt, value)
	}
	if value, ok := tu.mutation.To(); ok {
		_spec.SetField(transaction.FieldTo, field.TypeString, value)
	}
	if value, ok := tu.mutation.Gas(); ok {
		_spec.SetField(transaction.FieldGas, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedGas(); ok {
		_spec.AddField(transaction.FieldGas, field.TypeInt, value)
	}
	if value, ok := tu.mutation.GasPrice(); ok {
		_spec.SetField(transaction.FieldGasPrice, field.TypeString, value)
	}
	if value, ok := tu.mutation.GasTipCap(); ok {
		_spec.SetField(transaction.FieldGasTipCap, field.TypeString, value)
	}
	if value, ok := tu.mutation.GasFeeCap(); ok {
		_spec.SetField(transaction.FieldGasFeeCap, field.TypeString, value)
	}
	if value, ok := tu.mutation.Value(); ok {
		_spec.SetField(transaction.FieldValue, field.TypeString, value)
	}
	if value, ok := tu.mutation.Data(); ok {
		_spec.SetField(transaction.FieldData, field.TypeString, value)
	}
	if tu.mutation.DataCleared() {
		_spec.ClearField(transaction.FieldData, field.TypeString)
	}
	if value, ok := tu.mutation.Hash(); ok {
		_spec.SetField(transaction.FieldHash, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TransactionUpdateOne is the builder for updating a single Transaction entity.
type TransactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionMutation
}

// SetBlockNumber sets the "block_number" field.
func (tuo *TransactionUpdateOne) SetBlockNumber(i int) *TransactionUpdateOne {
	tuo.mutation.ResetBlockNumber()
	tuo.mutation.SetBlockNumber(i)
	return tuo
}

// SetNillableBlockNumber sets the "block_number" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableBlockNumber(i *int) *TransactionUpdateOne {
	if i != nil {
		tuo.SetBlockNumber(*i)
	}
	return tuo
}

// AddBlockNumber adds i to the "block_number" field.
func (tuo *TransactionUpdateOne) AddBlockNumber(i int) *TransactionUpdateOne {
	tuo.mutation.AddBlockNumber(i)
	return tuo
}

// SetType sets the "type" field.
func (tuo *TransactionUpdateOne) SetType(i int) *TransactionUpdateOne {
	tuo.mutation.ResetType()
	tuo.mutation.SetType(i)
	return tuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableType(i *int) *TransactionUpdateOne {
	if i != nil {
		tuo.SetType(*i)
	}
	return tuo
}

// AddType adds i to the "type" field.
func (tuo *TransactionUpdateOne) AddType(i int) *TransactionUpdateOne {
	tuo.mutation.AddType(i)
	return tuo
}

// SetChainID sets the "chain_id" field.
func (tuo *TransactionUpdateOne) SetChainID(i int) *TransactionUpdateOne {
	tuo.mutation.ResetChainID()
	tuo.mutation.SetChainID(i)
	return tuo
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableChainID(i *int) *TransactionUpdateOne {
	if i != nil {
		tuo.SetChainID(*i)
	}
	return tuo
}

// AddChainID adds i to the "chain_id" field.
func (tuo *TransactionUpdateOne) AddChainID(i int) *TransactionUpdateOne {
	tuo.mutation.AddChainID(i)
	return tuo
}

// SetNonce sets the "nonce" field.
func (tuo *TransactionUpdateOne) SetNonce(i int) *TransactionUpdateOne {
	tuo.mutation.ResetNonce()
	tuo.mutation.SetNonce(i)
	return tuo
}

// SetNillableNonce sets the "nonce" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableNonce(i *int) *TransactionUpdateOne {
	if i != nil {
		tuo.SetNonce(*i)
	}
	return tuo
}

// AddNonce adds i to the "nonce" field.
func (tuo *TransactionUpdateOne) AddNonce(i int) *TransactionUpdateOne {
	tuo.mutation.AddNonce(i)
	return tuo
}

// SetTo sets the "to" field.
func (tuo *TransactionUpdateOne) SetTo(s string) *TransactionUpdateOne {
	tuo.mutation.SetTo(s)
	return tuo
}

// SetNillableTo sets the "to" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableTo(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetTo(*s)
	}
	return tuo
}

// SetGas sets the "gas" field.
func (tuo *TransactionUpdateOne) SetGas(i int) *TransactionUpdateOne {
	tuo.mutation.ResetGas()
	tuo.mutation.SetGas(i)
	return tuo
}

// SetNillableGas sets the "gas" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableGas(i *int) *TransactionUpdateOne {
	if i != nil {
		tuo.SetGas(*i)
	}
	return tuo
}

// AddGas adds i to the "gas" field.
func (tuo *TransactionUpdateOne) AddGas(i int) *TransactionUpdateOne {
	tuo.mutation.AddGas(i)
	return tuo
}

// SetGasPrice sets the "gasPrice" field.
func (tuo *TransactionUpdateOne) SetGasPrice(s string) *TransactionUpdateOne {
	tuo.mutation.SetGasPrice(s)
	return tuo
}

// SetNillableGasPrice sets the "gasPrice" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableGasPrice(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetGasPrice(*s)
	}
	return tuo
}

// SetGasTipCap sets the "gasTipCap" field.
func (tuo *TransactionUpdateOne) SetGasTipCap(s string) *TransactionUpdateOne {
	tuo.mutation.SetGasTipCap(s)
	return tuo
}

// SetNillableGasTipCap sets the "gasTipCap" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableGasTipCap(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetGasTipCap(*s)
	}
	return tuo
}

// SetGasFeeCap sets the "gasFeeCap" field.
func (tuo *TransactionUpdateOne) SetGasFeeCap(s string) *TransactionUpdateOne {
	tuo.mutation.SetGasFeeCap(s)
	return tuo
}

// SetNillableGasFeeCap sets the "gasFeeCap" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableGasFeeCap(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetGasFeeCap(*s)
	}
	return tuo
}

// SetValue sets the "value" field.
func (tuo *TransactionUpdateOne) SetValue(s string) *TransactionUpdateOne {
	tuo.mutation.SetValue(s)
	return tuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableValue(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetValue(*s)
	}
	return tuo
}

// SetData sets the "data" field.
func (tuo *TransactionUpdateOne) SetData(s string) *TransactionUpdateOne {
	tuo.mutation.SetData(s)
	return tuo
}

// SetNillableData sets the "data" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableData(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetData(*s)
	}
	return tuo
}

// ClearData clears the value of the "data" field.
func (tuo *TransactionUpdateOne) ClearData() *TransactionUpdateOne {
	tuo.mutation.ClearData()
	return tuo
}

// SetHash sets the "hash" field.
func (tuo *TransactionUpdateOne) SetHash(s string) *TransactionUpdateOne {
	tuo.mutation.SetHash(s)
	return tuo
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableHash(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetHash(*s)
	}
	return tuo
}

// Mutation returns the TransactionMutation object of the builder.
func (tuo *TransactionUpdateOne) Mutation() *TransactionMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tuo *TransactionUpdateOne) Where(ps ...predicate.Transaction) *TransactionUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TransactionUpdateOne) Select(field string, fields ...string) *TransactionUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Transaction entity.
func (tuo *TransactionUpdateOne) Save(ctx context.Context) (*Transaction, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TransactionUpdateOne) SaveX(ctx context.Context) *Transaction {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TransactionUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TransactionUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TransactionUpdateOne) check() error {
	if v, ok := tuo.mutation.BlockNumber(); ok {
		if err := transaction.BlockNumberValidator(v); err != nil {
			return &ValidationError{Name: "block_number", err: fmt.Errorf(`ent: validator failed for field "Transaction.block_number": %w`, err)}
		}
	}
	return nil
}

func (tuo *TransactionUpdateOne) sqlSave(ctx context.Context) (_node *Transaction, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(transaction.Table, transaction.Columns, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Transaction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transaction.FieldID)
		for _, f := range fields {
			if !transaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.BlockNumber(); ok {
		_spec.SetField(transaction.FieldBlockNumber, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedBlockNumber(); ok {
		_spec.AddField(transaction.FieldBlockNumber, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.SetField(transaction.FieldType, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedType(); ok {
		_spec.AddField(transaction.FieldType, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.ChainID(); ok {
		_spec.SetField(transaction.FieldChainID, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedChainID(); ok {
		_spec.AddField(transaction.FieldChainID, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Nonce(); ok {
		_spec.SetField(transaction.FieldNonce, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedNonce(); ok {
		_spec.AddField(transaction.FieldNonce, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.To(); ok {
		_spec.SetField(transaction.FieldTo, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Gas(); ok {
		_spec.SetField(transaction.FieldGas, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedGas(); ok {
		_spec.AddField(transaction.FieldGas, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.GasPrice(); ok {
		_spec.SetField(transaction.FieldGasPrice, field.TypeString, value)
	}
	if value, ok := tuo.mutation.GasTipCap(); ok {
		_spec.SetField(transaction.FieldGasTipCap, field.TypeString, value)
	}
	if value, ok := tuo.mutation.GasFeeCap(); ok {
		_spec.SetField(transaction.FieldGasFeeCap, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Value(); ok {
		_spec.SetField(transaction.FieldValue, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Data(); ok {
		_spec.SetField(transaction.FieldData, field.TypeString, value)
	}
	if tuo.mutation.DataCleared() {
		_spec.ClearField(transaction.FieldData, field.TypeString)
	}
	if value, ok := tuo.mutation.Hash(); ok {
		_spec.SetField(transaction.FieldHash, field.TypeString, value)
	}
	_node = &Transaction{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
