// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nidhey27/cqrs-go/ent/predicate"
	"github.com/nidhey27/cqrs-go/ent/transaction"
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

// SetTransactionID sets the "transaction_id" field.
func (tu *TransactionUpdate) SetTransactionID(s string) *TransactionUpdate {
	tu.mutation.SetTransactionID(s)
	return tu
}

// SetNillableTransactionID sets the "transaction_id" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableTransactionID(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetTransactionID(*s)
	}
	return tu
}

// SetAccountNumber sets the "account_number" field.
func (tu *TransactionUpdate) SetAccountNumber(s string) *TransactionUpdate {
	tu.mutation.SetAccountNumber(s)
	return tu
}

// SetNillableAccountNumber sets the "account_number" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableAccountNumber(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetAccountNumber(*s)
	}
	return tu
}

// SetType sets the "type" field.
func (tu *TransactionUpdate) SetType(t transaction.Type) *TransactionUpdate {
	tu.mutation.SetType(t)
	return tu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableType(t *transaction.Type) *TransactionUpdate {
	if t != nil {
		tu.SetType(*t)
	}
	return tu
}

// SetAmount sets the "amount" field.
func (tu *TransactionUpdate) SetAmount(f float64) *TransactionUpdate {
	tu.mutation.ResetAmount()
	tu.mutation.SetAmount(f)
	return tu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableAmount(f *float64) *TransactionUpdate {
	if f != nil {
		tu.SetAmount(*f)
	}
	return tu
}

// AddAmount adds f to the "amount" field.
func (tu *TransactionUpdate) AddAmount(f float64) *TransactionUpdate {
	tu.mutation.AddAmount(f)
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
	if v, ok := tu.mutation.GetType(); ok {
		if err := transaction.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Transaction.type": %w`, err)}
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
	if value, ok := tu.mutation.TransactionID(); ok {
		_spec.SetField(transaction.FieldTransactionID, field.TypeString, value)
	}
	if value, ok := tu.mutation.AccountNumber(); ok {
		_spec.SetField(transaction.FieldAccountNumber, field.TypeString, value)
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.SetField(transaction.FieldType, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.Amount(); ok {
		_spec.SetField(transaction.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedAmount(); ok {
		_spec.AddField(transaction.FieldAmount, field.TypeFloat64, value)
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

// SetTransactionID sets the "transaction_id" field.
func (tuo *TransactionUpdateOne) SetTransactionID(s string) *TransactionUpdateOne {
	tuo.mutation.SetTransactionID(s)
	return tuo
}

// SetNillableTransactionID sets the "transaction_id" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableTransactionID(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetTransactionID(*s)
	}
	return tuo
}

// SetAccountNumber sets the "account_number" field.
func (tuo *TransactionUpdateOne) SetAccountNumber(s string) *TransactionUpdateOne {
	tuo.mutation.SetAccountNumber(s)
	return tuo
}

// SetNillableAccountNumber sets the "account_number" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableAccountNumber(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetAccountNumber(*s)
	}
	return tuo
}

// SetType sets the "type" field.
func (tuo *TransactionUpdateOne) SetType(t transaction.Type) *TransactionUpdateOne {
	tuo.mutation.SetType(t)
	return tuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableType(t *transaction.Type) *TransactionUpdateOne {
	if t != nil {
		tuo.SetType(*t)
	}
	return tuo
}

// SetAmount sets the "amount" field.
func (tuo *TransactionUpdateOne) SetAmount(f float64) *TransactionUpdateOne {
	tuo.mutation.ResetAmount()
	tuo.mutation.SetAmount(f)
	return tuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableAmount(f *float64) *TransactionUpdateOne {
	if f != nil {
		tuo.SetAmount(*f)
	}
	return tuo
}

// AddAmount adds f to the "amount" field.
func (tuo *TransactionUpdateOne) AddAmount(f float64) *TransactionUpdateOne {
	tuo.mutation.AddAmount(f)
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
	if v, ok := tuo.mutation.GetType(); ok {
		if err := transaction.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Transaction.type": %w`, err)}
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
	if value, ok := tuo.mutation.TransactionID(); ok {
		_spec.SetField(transaction.FieldTransactionID, field.TypeString, value)
	}
	if value, ok := tuo.mutation.AccountNumber(); ok {
		_spec.SetField(transaction.FieldAccountNumber, field.TypeString, value)
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.SetField(transaction.FieldType, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.Amount(); ok {
		_spec.SetField(transaction.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedAmount(); ok {
		_spec.AddField(transaction.FieldAmount, field.TypeFloat64, value)
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