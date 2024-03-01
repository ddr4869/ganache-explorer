// Code generated by ent, DO NOT EDIT.

package transaction

import (
	"entgo.io/ent/dialect/sql"
	"github.com/ddr4869/ether-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldType, v))
}

// ChainID applies equality check predicate on the "chain_id" field. It's identical to ChainIDEQ.
func ChainID(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldChainID, v))
}

// Nonce applies equality check predicate on the "nonce" field. It's identical to NonceEQ.
func Nonce(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldNonce, v))
}

// To applies equality check predicate on the "to" field. It's identical to ToEQ.
func To(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldTo, v))
}

// Gas applies equality check predicate on the "gas" field. It's identical to GasEQ.
func Gas(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldGas, v))
}

// GasPrice applies equality check predicate on the "gasPrice" field. It's identical to GasPriceEQ.
func GasPrice(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldGasPrice, v))
}

// GasTipCap applies equality check predicate on the "gasTipCap" field. It's identical to GasTipCapEQ.
func GasTipCap(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldGasTipCap, v))
}

// GasFeeCap applies equality check predicate on the "gasFeeCap" field. It's identical to GasFeeCapEQ.
func GasFeeCap(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldGasFeeCap, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldValue, v))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldData, v))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldHash, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldType, v))
}

// ChainIDEQ applies the EQ predicate on the "chain_id" field.
func ChainIDEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldChainID, v))
}

// ChainIDNEQ applies the NEQ predicate on the "chain_id" field.
func ChainIDNEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldChainID, v))
}

// ChainIDIn applies the In predicate on the "chain_id" field.
func ChainIDIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldChainID, vs...))
}

// ChainIDNotIn applies the NotIn predicate on the "chain_id" field.
func ChainIDNotIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldChainID, vs...))
}

// ChainIDGT applies the GT predicate on the "chain_id" field.
func ChainIDGT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldChainID, v))
}

// ChainIDGTE applies the GTE predicate on the "chain_id" field.
func ChainIDGTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldChainID, v))
}

// ChainIDLT applies the LT predicate on the "chain_id" field.
func ChainIDLT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldChainID, v))
}

// ChainIDLTE applies the LTE predicate on the "chain_id" field.
func ChainIDLTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldChainID, v))
}

// NonceEQ applies the EQ predicate on the "nonce" field.
func NonceEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldNonce, v))
}

// NonceNEQ applies the NEQ predicate on the "nonce" field.
func NonceNEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldNonce, v))
}

// NonceIn applies the In predicate on the "nonce" field.
func NonceIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldNonce, vs...))
}

// NonceNotIn applies the NotIn predicate on the "nonce" field.
func NonceNotIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldNonce, vs...))
}

// NonceGT applies the GT predicate on the "nonce" field.
func NonceGT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldNonce, v))
}

// NonceGTE applies the GTE predicate on the "nonce" field.
func NonceGTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldNonce, v))
}

// NonceLT applies the LT predicate on the "nonce" field.
func NonceLT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldNonce, v))
}

// NonceLTE applies the LTE predicate on the "nonce" field.
func NonceLTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldNonce, v))
}

// ToEQ applies the EQ predicate on the "to" field.
func ToEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldTo, v))
}

// ToNEQ applies the NEQ predicate on the "to" field.
func ToNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldTo, v))
}

// ToIn applies the In predicate on the "to" field.
func ToIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldTo, vs...))
}

// ToNotIn applies the NotIn predicate on the "to" field.
func ToNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldTo, vs...))
}

// ToGT applies the GT predicate on the "to" field.
func ToGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldTo, v))
}

// ToGTE applies the GTE predicate on the "to" field.
func ToGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldTo, v))
}

// ToLT applies the LT predicate on the "to" field.
func ToLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldTo, v))
}

// ToLTE applies the LTE predicate on the "to" field.
func ToLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldTo, v))
}

// ToContains applies the Contains predicate on the "to" field.
func ToContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldTo, v))
}

// ToHasPrefix applies the HasPrefix predicate on the "to" field.
func ToHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldTo, v))
}

// ToHasSuffix applies the HasSuffix predicate on the "to" field.
func ToHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldTo, v))
}

// ToEqualFold applies the EqualFold predicate on the "to" field.
func ToEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldTo, v))
}

// ToContainsFold applies the ContainsFold predicate on the "to" field.
func ToContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldTo, v))
}

// GasEQ applies the EQ predicate on the "gas" field.
func GasEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldGas, v))
}

// GasNEQ applies the NEQ predicate on the "gas" field.
func GasNEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldGas, v))
}

// GasIn applies the In predicate on the "gas" field.
func GasIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldGas, vs...))
}

// GasNotIn applies the NotIn predicate on the "gas" field.
func GasNotIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldGas, vs...))
}

// GasGT applies the GT predicate on the "gas" field.
func GasGT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldGas, v))
}

// GasGTE applies the GTE predicate on the "gas" field.
func GasGTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldGas, v))
}

// GasLT applies the LT predicate on the "gas" field.
func GasLT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldGas, v))
}

// GasLTE applies the LTE predicate on the "gas" field.
func GasLTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldGas, v))
}

// GasPriceEQ applies the EQ predicate on the "gasPrice" field.
func GasPriceEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldGasPrice, v))
}

// GasPriceNEQ applies the NEQ predicate on the "gasPrice" field.
func GasPriceNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldGasPrice, v))
}

// GasPriceIn applies the In predicate on the "gasPrice" field.
func GasPriceIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldGasPrice, vs...))
}

// GasPriceNotIn applies the NotIn predicate on the "gasPrice" field.
func GasPriceNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldGasPrice, vs...))
}

// GasPriceGT applies the GT predicate on the "gasPrice" field.
func GasPriceGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldGasPrice, v))
}

// GasPriceGTE applies the GTE predicate on the "gasPrice" field.
func GasPriceGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldGasPrice, v))
}

// GasPriceLT applies the LT predicate on the "gasPrice" field.
func GasPriceLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldGasPrice, v))
}

// GasPriceLTE applies the LTE predicate on the "gasPrice" field.
func GasPriceLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldGasPrice, v))
}

// GasPriceContains applies the Contains predicate on the "gasPrice" field.
func GasPriceContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldGasPrice, v))
}

// GasPriceHasPrefix applies the HasPrefix predicate on the "gasPrice" field.
func GasPriceHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldGasPrice, v))
}

// GasPriceHasSuffix applies the HasSuffix predicate on the "gasPrice" field.
func GasPriceHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldGasPrice, v))
}

// GasPriceEqualFold applies the EqualFold predicate on the "gasPrice" field.
func GasPriceEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldGasPrice, v))
}

// GasPriceContainsFold applies the ContainsFold predicate on the "gasPrice" field.
func GasPriceContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldGasPrice, v))
}

// GasTipCapEQ applies the EQ predicate on the "gasTipCap" field.
func GasTipCapEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldGasTipCap, v))
}

// GasTipCapNEQ applies the NEQ predicate on the "gasTipCap" field.
func GasTipCapNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldGasTipCap, v))
}

// GasTipCapIn applies the In predicate on the "gasTipCap" field.
func GasTipCapIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldGasTipCap, vs...))
}

// GasTipCapNotIn applies the NotIn predicate on the "gasTipCap" field.
func GasTipCapNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldGasTipCap, vs...))
}

// GasTipCapGT applies the GT predicate on the "gasTipCap" field.
func GasTipCapGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldGasTipCap, v))
}

// GasTipCapGTE applies the GTE predicate on the "gasTipCap" field.
func GasTipCapGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldGasTipCap, v))
}

// GasTipCapLT applies the LT predicate on the "gasTipCap" field.
func GasTipCapLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldGasTipCap, v))
}

// GasTipCapLTE applies the LTE predicate on the "gasTipCap" field.
func GasTipCapLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldGasTipCap, v))
}

// GasTipCapContains applies the Contains predicate on the "gasTipCap" field.
func GasTipCapContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldGasTipCap, v))
}

// GasTipCapHasPrefix applies the HasPrefix predicate on the "gasTipCap" field.
func GasTipCapHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldGasTipCap, v))
}

// GasTipCapHasSuffix applies the HasSuffix predicate on the "gasTipCap" field.
func GasTipCapHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldGasTipCap, v))
}

// GasTipCapEqualFold applies the EqualFold predicate on the "gasTipCap" field.
func GasTipCapEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldGasTipCap, v))
}

// GasTipCapContainsFold applies the ContainsFold predicate on the "gasTipCap" field.
func GasTipCapContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldGasTipCap, v))
}

// GasFeeCapEQ applies the EQ predicate on the "gasFeeCap" field.
func GasFeeCapEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldGasFeeCap, v))
}

// GasFeeCapNEQ applies the NEQ predicate on the "gasFeeCap" field.
func GasFeeCapNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldGasFeeCap, v))
}

// GasFeeCapIn applies the In predicate on the "gasFeeCap" field.
func GasFeeCapIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldGasFeeCap, vs...))
}

// GasFeeCapNotIn applies the NotIn predicate on the "gasFeeCap" field.
func GasFeeCapNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldGasFeeCap, vs...))
}

// GasFeeCapGT applies the GT predicate on the "gasFeeCap" field.
func GasFeeCapGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldGasFeeCap, v))
}

// GasFeeCapGTE applies the GTE predicate on the "gasFeeCap" field.
func GasFeeCapGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldGasFeeCap, v))
}

// GasFeeCapLT applies the LT predicate on the "gasFeeCap" field.
func GasFeeCapLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldGasFeeCap, v))
}

// GasFeeCapLTE applies the LTE predicate on the "gasFeeCap" field.
func GasFeeCapLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldGasFeeCap, v))
}

// GasFeeCapContains applies the Contains predicate on the "gasFeeCap" field.
func GasFeeCapContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldGasFeeCap, v))
}

// GasFeeCapHasPrefix applies the HasPrefix predicate on the "gasFeeCap" field.
func GasFeeCapHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldGasFeeCap, v))
}

// GasFeeCapHasSuffix applies the HasSuffix predicate on the "gasFeeCap" field.
func GasFeeCapHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldGasFeeCap, v))
}

// GasFeeCapEqualFold applies the EqualFold predicate on the "gasFeeCap" field.
func GasFeeCapEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldGasFeeCap, v))
}

// GasFeeCapContainsFold applies the ContainsFold predicate on the "gasFeeCap" field.
func GasFeeCapContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldGasFeeCap, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldValue, v))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldData, v))
}

// DataContains applies the Contains predicate on the "data" field.
func DataContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldData, v))
}

// DataHasPrefix applies the HasPrefix predicate on the "data" field.
func DataHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldData, v))
}

// DataHasSuffix applies the HasSuffix predicate on the "data" field.
func DataHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldData, v))
}

// DataIsNil applies the IsNil predicate on the "data" field.
func DataIsNil() predicate.Transaction {
	return predicate.Transaction(sql.FieldIsNull(FieldData))
}

// DataNotNil applies the NotNil predicate on the "data" field.
func DataNotNil() predicate.Transaction {
	return predicate.Transaction(sql.FieldNotNull(FieldData))
}

// DataEqualFold applies the EqualFold predicate on the "data" field.
func DataEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldData, v))
}

// DataContainsFold applies the ContainsFold predicate on the "data" field.
func DataContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldData, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldHash, v))
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldHash, v))
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldHash, v))
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldHash, v))
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldHash, v))
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldHash, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.NotPredicates(p))
}
