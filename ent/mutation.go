// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/nidhey27/cqrs-go/ent/account"
	"github.com/nidhey27/cqrs-go/ent/predicate"
	"github.com/nidhey27/cqrs-go/ent/transaction"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAccount     = "Account"
	TypeTransaction = "Transaction"
)

// AccountMutation represents an operation that mutates the Account nodes in the graph.
type AccountMutation struct {
	config
	op                  Op
	typ                 string
	id                  *int
	account_number      *string
	account_holder_name *string
	account_balance     *float64
	addaccount_balance  *float64
	clearedFields       map[string]struct{}
	done                bool
	oldValue            func(context.Context) (*Account, error)
	predicates          []predicate.Account
}

var _ ent.Mutation = (*AccountMutation)(nil)

// accountOption allows management of the mutation configuration using functional options.
type accountOption func(*AccountMutation)

// newAccountMutation creates new mutation for the Account entity.
func newAccountMutation(c config, op Op, opts ...accountOption) *AccountMutation {
	m := &AccountMutation{
		config:        c,
		op:            op,
		typ:           TypeAccount,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAccountID sets the ID field of the mutation.
func withAccountID(id int) accountOption {
	return func(m *AccountMutation) {
		var (
			err   error
			once  sync.Once
			value *Account
		)
		m.oldValue = func(ctx context.Context) (*Account, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Account.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAccount sets the old Account of the mutation.
func withAccount(node *Account) accountOption {
	return func(m *AccountMutation) {
		m.oldValue = func(context.Context) (*Account, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AccountMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AccountMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AccountMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AccountMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Account.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAccountNumber sets the "account_number" field.
func (m *AccountMutation) SetAccountNumber(s string) {
	m.account_number = &s
}

// AccountNumber returns the value of the "account_number" field in the mutation.
func (m *AccountMutation) AccountNumber() (r string, exists bool) {
	v := m.account_number
	if v == nil {
		return
	}
	return *v, true
}

// OldAccountNumber returns the old "account_number" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldAccountNumber(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAccountNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAccountNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAccountNumber: %w", err)
	}
	return oldValue.AccountNumber, nil
}

// ResetAccountNumber resets all changes to the "account_number" field.
func (m *AccountMutation) ResetAccountNumber() {
	m.account_number = nil
}

// SetAccountHolderName sets the "account_holder_name" field.
func (m *AccountMutation) SetAccountHolderName(s string) {
	m.account_holder_name = &s
}

// AccountHolderName returns the value of the "account_holder_name" field in the mutation.
func (m *AccountMutation) AccountHolderName() (r string, exists bool) {
	v := m.account_holder_name
	if v == nil {
		return
	}
	return *v, true
}

// OldAccountHolderName returns the old "account_holder_name" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldAccountHolderName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAccountHolderName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAccountHolderName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAccountHolderName: %w", err)
	}
	return oldValue.AccountHolderName, nil
}

// ResetAccountHolderName resets all changes to the "account_holder_name" field.
func (m *AccountMutation) ResetAccountHolderName() {
	m.account_holder_name = nil
}

// SetAccountBalance sets the "account_balance" field.
func (m *AccountMutation) SetAccountBalance(f float64) {
	m.account_balance = &f
	m.addaccount_balance = nil
}

// AccountBalance returns the value of the "account_balance" field in the mutation.
func (m *AccountMutation) AccountBalance() (r float64, exists bool) {
	v := m.account_balance
	if v == nil {
		return
	}
	return *v, true
}

// OldAccountBalance returns the old "account_balance" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldAccountBalance(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAccountBalance is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAccountBalance requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAccountBalance: %w", err)
	}
	return oldValue.AccountBalance, nil
}

// AddAccountBalance adds f to the "account_balance" field.
func (m *AccountMutation) AddAccountBalance(f float64) {
	if m.addaccount_balance != nil {
		*m.addaccount_balance += f
	} else {
		m.addaccount_balance = &f
	}
}

// AddedAccountBalance returns the value that was added to the "account_balance" field in this mutation.
func (m *AccountMutation) AddedAccountBalance() (r float64, exists bool) {
	v := m.addaccount_balance
	if v == nil {
		return
	}
	return *v, true
}

// ResetAccountBalance resets all changes to the "account_balance" field.
func (m *AccountMutation) ResetAccountBalance() {
	m.account_balance = nil
	m.addaccount_balance = nil
}

// Where appends a list predicates to the AccountMutation builder.
func (m *AccountMutation) Where(ps ...predicate.Account) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the AccountMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *AccountMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Account, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *AccountMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *AccountMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Account).
func (m *AccountMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AccountMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.account_number != nil {
		fields = append(fields, account.FieldAccountNumber)
	}
	if m.account_holder_name != nil {
		fields = append(fields, account.FieldAccountHolderName)
	}
	if m.account_balance != nil {
		fields = append(fields, account.FieldAccountBalance)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AccountMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case account.FieldAccountNumber:
		return m.AccountNumber()
	case account.FieldAccountHolderName:
		return m.AccountHolderName()
	case account.FieldAccountBalance:
		return m.AccountBalance()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AccountMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case account.FieldAccountNumber:
		return m.OldAccountNumber(ctx)
	case account.FieldAccountHolderName:
		return m.OldAccountHolderName(ctx)
	case account.FieldAccountBalance:
		return m.OldAccountBalance(ctx)
	}
	return nil, fmt.Errorf("unknown Account field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) SetField(name string, value ent.Value) error {
	switch name {
	case account.FieldAccountNumber:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAccountNumber(v)
		return nil
	case account.FieldAccountHolderName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAccountHolderName(v)
		return nil
	case account.FieldAccountBalance:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAccountBalance(v)
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AccountMutation) AddedFields() []string {
	var fields []string
	if m.addaccount_balance != nil {
		fields = append(fields, account.FieldAccountBalance)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AccountMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case account.FieldAccountBalance:
		return m.AddedAccountBalance()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) AddField(name string, value ent.Value) error {
	switch name {
	case account.FieldAccountBalance:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAccountBalance(v)
		return nil
	}
	return fmt.Errorf("unknown Account numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AccountMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AccountMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AccountMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Account nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AccountMutation) ResetField(name string) error {
	switch name {
	case account.FieldAccountNumber:
		m.ResetAccountNumber()
		return nil
	case account.FieldAccountHolderName:
		m.ResetAccountHolderName()
		return nil
	case account.FieldAccountBalance:
		m.ResetAccountBalance()
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AccountMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AccountMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AccountMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AccountMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AccountMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AccountMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AccountMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Account unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AccountMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Account edge %s", name)
}

// TransactionMutation represents an operation that mutates the Transaction nodes in the graph.
type TransactionMutation struct {
	config
	op             Op
	typ            string
	id             *int
	transaction_id *string
	account_number *string
	_type          *transaction.Type
	amount         *float64
	addamount      *float64
	timestamp      *time.Time
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*Transaction, error)
	predicates     []predicate.Transaction
}

var _ ent.Mutation = (*TransactionMutation)(nil)

// transactionOption allows management of the mutation configuration using functional options.
type transactionOption func(*TransactionMutation)

// newTransactionMutation creates new mutation for the Transaction entity.
func newTransactionMutation(c config, op Op, opts ...transactionOption) *TransactionMutation {
	m := &TransactionMutation{
		config:        c,
		op:            op,
		typ:           TypeTransaction,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTransactionID sets the ID field of the mutation.
func withTransactionID(id int) transactionOption {
	return func(m *TransactionMutation) {
		var (
			err   error
			once  sync.Once
			value *Transaction
		)
		m.oldValue = func(ctx context.Context) (*Transaction, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Transaction.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTransaction sets the old Transaction of the mutation.
func withTransaction(node *Transaction) transactionOption {
	return func(m *TransactionMutation) {
		m.oldValue = func(context.Context) (*Transaction, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TransactionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TransactionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TransactionMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TransactionMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Transaction.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTransactionID sets the "transaction_id" field.
func (m *TransactionMutation) SetTransactionID(s string) {
	m.transaction_id = &s
}

// TransactionID returns the value of the "transaction_id" field in the mutation.
func (m *TransactionMutation) TransactionID() (r string, exists bool) {
	v := m.transaction_id
	if v == nil {
		return
	}
	return *v, true
}

// OldTransactionID returns the old "transaction_id" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldTransactionID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTransactionID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTransactionID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTransactionID: %w", err)
	}
	return oldValue.TransactionID, nil
}

// ResetTransactionID resets all changes to the "transaction_id" field.
func (m *TransactionMutation) ResetTransactionID() {
	m.transaction_id = nil
}

// SetAccountNumber sets the "account_number" field.
func (m *TransactionMutation) SetAccountNumber(s string) {
	m.account_number = &s
}

// AccountNumber returns the value of the "account_number" field in the mutation.
func (m *TransactionMutation) AccountNumber() (r string, exists bool) {
	v := m.account_number
	if v == nil {
		return
	}
	return *v, true
}

// OldAccountNumber returns the old "account_number" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldAccountNumber(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAccountNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAccountNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAccountNumber: %w", err)
	}
	return oldValue.AccountNumber, nil
}

// ResetAccountNumber resets all changes to the "account_number" field.
func (m *TransactionMutation) ResetAccountNumber() {
	m.account_number = nil
}

// SetType sets the "type" field.
func (m *TransactionMutation) SetType(t transaction.Type) {
	m._type = &t
}

// GetType returns the value of the "type" field in the mutation.
func (m *TransactionMutation) GetType() (r transaction.Type, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old "type" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldType(ctx context.Context) (v transaction.Type, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// ResetType resets all changes to the "type" field.
func (m *TransactionMutation) ResetType() {
	m._type = nil
}

// SetAmount sets the "amount" field.
func (m *TransactionMutation) SetAmount(f float64) {
	m.amount = &f
	m.addamount = nil
}

// Amount returns the value of the "amount" field in the mutation.
func (m *TransactionMutation) Amount() (r float64, exists bool) {
	v := m.amount
	if v == nil {
		return
	}
	return *v, true
}

// OldAmount returns the old "amount" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldAmount(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAmount: %w", err)
	}
	return oldValue.Amount, nil
}

// AddAmount adds f to the "amount" field.
func (m *TransactionMutation) AddAmount(f float64) {
	if m.addamount != nil {
		*m.addamount += f
	} else {
		m.addamount = &f
	}
}

// AddedAmount returns the value that was added to the "amount" field in this mutation.
func (m *TransactionMutation) AddedAmount() (r float64, exists bool) {
	v := m.addamount
	if v == nil {
		return
	}
	return *v, true
}

// ResetAmount resets all changes to the "amount" field.
func (m *TransactionMutation) ResetAmount() {
	m.amount = nil
	m.addamount = nil
}

// SetTimestamp sets the "timestamp" field.
func (m *TransactionMutation) SetTimestamp(t time.Time) {
	m.timestamp = &t
}

// Timestamp returns the value of the "timestamp" field in the mutation.
func (m *TransactionMutation) Timestamp() (r time.Time, exists bool) {
	v := m.timestamp
	if v == nil {
		return
	}
	return *v, true
}

// OldTimestamp returns the old "timestamp" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldTimestamp(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTimestamp is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTimestamp requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTimestamp: %w", err)
	}
	return oldValue.Timestamp, nil
}

// ResetTimestamp resets all changes to the "timestamp" field.
func (m *TransactionMutation) ResetTimestamp() {
	m.timestamp = nil
}

// Where appends a list predicates to the TransactionMutation builder.
func (m *TransactionMutation) Where(ps ...predicate.Transaction) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TransactionMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TransactionMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Transaction, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TransactionMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TransactionMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Transaction).
func (m *TransactionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TransactionMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.transaction_id != nil {
		fields = append(fields, transaction.FieldTransactionID)
	}
	if m.account_number != nil {
		fields = append(fields, transaction.FieldAccountNumber)
	}
	if m._type != nil {
		fields = append(fields, transaction.FieldType)
	}
	if m.amount != nil {
		fields = append(fields, transaction.FieldAmount)
	}
	if m.timestamp != nil {
		fields = append(fields, transaction.FieldTimestamp)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TransactionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case transaction.FieldTransactionID:
		return m.TransactionID()
	case transaction.FieldAccountNumber:
		return m.AccountNumber()
	case transaction.FieldType:
		return m.GetType()
	case transaction.FieldAmount:
		return m.Amount()
	case transaction.FieldTimestamp:
		return m.Timestamp()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TransactionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case transaction.FieldTransactionID:
		return m.OldTransactionID(ctx)
	case transaction.FieldAccountNumber:
		return m.OldAccountNumber(ctx)
	case transaction.FieldType:
		return m.OldType(ctx)
	case transaction.FieldAmount:
		return m.OldAmount(ctx)
	case transaction.FieldTimestamp:
		return m.OldTimestamp(ctx)
	}
	return nil, fmt.Errorf("unknown Transaction field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TransactionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case transaction.FieldTransactionID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTransactionID(v)
		return nil
	case transaction.FieldAccountNumber:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAccountNumber(v)
		return nil
	case transaction.FieldType:
		v, ok := value.(transaction.Type)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	case transaction.FieldAmount:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAmount(v)
		return nil
	case transaction.FieldTimestamp:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTimestamp(v)
		return nil
	}
	return fmt.Errorf("unknown Transaction field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TransactionMutation) AddedFields() []string {
	var fields []string
	if m.addamount != nil {
		fields = append(fields, transaction.FieldAmount)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TransactionMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case transaction.FieldAmount:
		return m.AddedAmount()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TransactionMutation) AddField(name string, value ent.Value) error {
	switch name {
	case transaction.FieldAmount:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAmount(v)
		return nil
	}
	return fmt.Errorf("unknown Transaction numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TransactionMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TransactionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TransactionMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Transaction nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TransactionMutation) ResetField(name string) error {
	switch name {
	case transaction.FieldTransactionID:
		m.ResetTransactionID()
		return nil
	case transaction.FieldAccountNumber:
		m.ResetAccountNumber()
		return nil
	case transaction.FieldType:
		m.ResetType()
		return nil
	case transaction.FieldAmount:
		m.ResetAmount()
		return nil
	case transaction.FieldTimestamp:
		m.ResetTimestamp()
		return nil
	}
	return fmt.Errorf("unknown Transaction field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TransactionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TransactionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TransactionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TransactionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TransactionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TransactionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TransactionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Transaction unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TransactionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Transaction edge %s", name)
}