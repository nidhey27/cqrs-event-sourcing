// Code generated by ent, DO NOT EDIT.

package account

import (
	"entgo.io/ent/dialect/sql"
	"github.com/nidhey27/cqrs-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldID, id))
}

// AccountNumber applies equality check predicate on the "account_number" field. It's identical to AccountNumberEQ.
func AccountNumber(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldAccountNumber, v))
}

// AccountHolderName applies equality check predicate on the "account_holder_name" field. It's identical to AccountHolderNameEQ.
func AccountHolderName(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldAccountHolderName, v))
}

// AccountBalance applies equality check predicate on the "account_balance" field. It's identical to AccountBalanceEQ.
func AccountBalance(v float64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldAccountBalance, v))
}

// AccountNumberEQ applies the EQ predicate on the "account_number" field.
func AccountNumberEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldAccountNumber, v))
}

// AccountNumberNEQ applies the NEQ predicate on the "account_number" field.
func AccountNumberNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldAccountNumber, v))
}

// AccountNumberIn applies the In predicate on the "account_number" field.
func AccountNumberIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldAccountNumber, vs...))
}

// AccountNumberNotIn applies the NotIn predicate on the "account_number" field.
func AccountNumberNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldAccountNumber, vs...))
}

// AccountNumberGT applies the GT predicate on the "account_number" field.
func AccountNumberGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldAccountNumber, v))
}

// AccountNumberGTE applies the GTE predicate on the "account_number" field.
func AccountNumberGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldAccountNumber, v))
}

// AccountNumberLT applies the LT predicate on the "account_number" field.
func AccountNumberLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldAccountNumber, v))
}

// AccountNumberLTE applies the LTE predicate on the "account_number" field.
func AccountNumberLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldAccountNumber, v))
}

// AccountNumberContains applies the Contains predicate on the "account_number" field.
func AccountNumberContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldAccountNumber, v))
}

// AccountNumberHasPrefix applies the HasPrefix predicate on the "account_number" field.
func AccountNumberHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldAccountNumber, v))
}

// AccountNumberHasSuffix applies the HasSuffix predicate on the "account_number" field.
func AccountNumberHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldAccountNumber, v))
}

// AccountNumberEqualFold applies the EqualFold predicate on the "account_number" field.
func AccountNumberEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldAccountNumber, v))
}

// AccountNumberContainsFold applies the ContainsFold predicate on the "account_number" field.
func AccountNumberContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldAccountNumber, v))
}

// AccountHolderNameEQ applies the EQ predicate on the "account_holder_name" field.
func AccountHolderNameEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldAccountHolderName, v))
}

// AccountHolderNameNEQ applies the NEQ predicate on the "account_holder_name" field.
func AccountHolderNameNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldAccountHolderName, v))
}

// AccountHolderNameIn applies the In predicate on the "account_holder_name" field.
func AccountHolderNameIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldAccountHolderName, vs...))
}

// AccountHolderNameNotIn applies the NotIn predicate on the "account_holder_name" field.
func AccountHolderNameNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldAccountHolderName, vs...))
}

// AccountHolderNameGT applies the GT predicate on the "account_holder_name" field.
func AccountHolderNameGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldAccountHolderName, v))
}

// AccountHolderNameGTE applies the GTE predicate on the "account_holder_name" field.
func AccountHolderNameGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldAccountHolderName, v))
}

// AccountHolderNameLT applies the LT predicate on the "account_holder_name" field.
func AccountHolderNameLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldAccountHolderName, v))
}

// AccountHolderNameLTE applies the LTE predicate on the "account_holder_name" field.
func AccountHolderNameLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldAccountHolderName, v))
}

// AccountHolderNameContains applies the Contains predicate on the "account_holder_name" field.
func AccountHolderNameContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldAccountHolderName, v))
}

// AccountHolderNameHasPrefix applies the HasPrefix predicate on the "account_holder_name" field.
func AccountHolderNameHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldAccountHolderName, v))
}

// AccountHolderNameHasSuffix applies the HasSuffix predicate on the "account_holder_name" field.
func AccountHolderNameHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldAccountHolderName, v))
}

// AccountHolderNameEqualFold applies the EqualFold predicate on the "account_holder_name" field.
func AccountHolderNameEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldAccountHolderName, v))
}

// AccountHolderNameContainsFold applies the ContainsFold predicate on the "account_holder_name" field.
func AccountHolderNameContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldAccountHolderName, v))
}

// AccountBalanceEQ applies the EQ predicate on the "account_balance" field.
func AccountBalanceEQ(v float64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldAccountBalance, v))
}

// AccountBalanceNEQ applies the NEQ predicate on the "account_balance" field.
func AccountBalanceNEQ(v float64) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldAccountBalance, v))
}

// AccountBalanceIn applies the In predicate on the "account_balance" field.
func AccountBalanceIn(vs ...float64) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldAccountBalance, vs...))
}

// AccountBalanceNotIn applies the NotIn predicate on the "account_balance" field.
func AccountBalanceNotIn(vs ...float64) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldAccountBalance, vs...))
}

// AccountBalanceGT applies the GT predicate on the "account_balance" field.
func AccountBalanceGT(v float64) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldAccountBalance, v))
}

// AccountBalanceGTE applies the GTE predicate on the "account_balance" field.
func AccountBalanceGTE(v float64) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldAccountBalance, v))
}

// AccountBalanceLT applies the LT predicate on the "account_balance" field.
func AccountBalanceLT(v float64) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldAccountBalance, v))
}

// AccountBalanceLTE applies the LTE predicate on the "account_balance" field.
func AccountBalanceLTE(v float64) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldAccountBalance, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(sql.NotPredicates(p))
}
