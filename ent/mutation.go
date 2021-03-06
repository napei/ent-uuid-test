// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"test/ent/testmodel"
	"time"

	"github.com/facebook/ent"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTestModel = "TestModel"
)

// TestModelMutation represents an operation that mutate the TestModels
// nodes in the graph.
type TestModelMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	create_time   *time.Time
	update_time   *time.Time
	test          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*TestModel, error)
}

var _ ent.Mutation = (*TestModelMutation)(nil)

// testmodelOption allows to manage the mutation configuration using functional options.
type testmodelOption func(*TestModelMutation)

// newTestModelMutation creates new mutation for $n.Name.
func newTestModelMutation(c config, op Op, opts ...testmodelOption) *TestModelMutation {
	m := &TestModelMutation{
		config:        c,
		op:            op,
		typ:           TypeTestModel,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTestModelID sets the id field of the mutation.
func withTestModelID(id uuid.UUID) testmodelOption {
	return func(m *TestModelMutation) {
		var (
			err   error
			once  sync.Once
			value *TestModel
		)
		m.oldValue = func(ctx context.Context) (*TestModel, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().TestModel.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTestModel sets the old TestModel of the mutation.
func withTestModel(node *TestModel) testmodelOption {
	return func(m *TestModelMutation) {
		m.oldValue = func(context.Context) (*TestModel, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TestModelMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TestModelMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on TestModel creation.
func (m *TestModelMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *TestModelMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCreateTime sets the create_time field.
func (m *TestModelMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the create_time value in the mutation.
func (m *TestModelMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old create_time value of the TestModel.
// If the TestModel object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *TestModelMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreateTime is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime reset all changes of the "create_time" field.
func (m *TestModelMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the update_time field.
func (m *TestModelMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the update_time value in the mutation.
func (m *TestModelMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old update_time value of the TestModel.
// If the TestModel object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *TestModelMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdateTime is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime reset all changes of the "update_time" field.
func (m *TestModelMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetTest sets the test field.
func (m *TestModelMutation) SetTest(s string) {
	m.test = &s
}

// Test returns the test value in the mutation.
func (m *TestModelMutation) Test() (r string, exists bool) {
	v := m.test
	if v == nil {
		return
	}
	return *v, true
}

// OldTest returns the old test value of the TestModel.
// If the TestModel object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *TestModelMutation) OldTest(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTest is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTest requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTest: %w", err)
	}
	return oldValue.Test, nil
}

// ResetTest reset all changes of the "test" field.
func (m *TestModelMutation) ResetTest() {
	m.test = nil
}

// Op returns the operation name.
func (m *TestModelMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (TestModel).
func (m *TestModelMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *TestModelMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.create_time != nil {
		fields = append(fields, testmodel.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, testmodel.FieldUpdateTime)
	}
	if m.test != nil {
		fields = append(fields, testmodel.FieldTest)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *TestModelMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case testmodel.FieldCreateTime:
		return m.CreateTime()
	case testmodel.FieldUpdateTime:
		return m.UpdateTime()
	case testmodel.FieldTest:
		return m.Test()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *TestModelMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case testmodel.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case testmodel.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case testmodel.FieldTest:
		return m.OldTest(ctx)
	}
	return nil, fmt.Errorf("unknown TestModel field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TestModelMutation) SetField(name string, value ent.Value) error {
	switch name {
	case testmodel.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case testmodel.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case testmodel.FieldTest:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTest(v)
		return nil
	}
	return fmt.Errorf("unknown TestModel field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *TestModelMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *TestModelMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TestModelMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown TestModel numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *TestModelMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *TestModelMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *TestModelMutation) ClearField(name string) error {
	return fmt.Errorf("unknown TestModel nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *TestModelMutation) ResetField(name string) error {
	switch name {
	case testmodel.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case testmodel.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case testmodel.FieldTest:
		m.ResetTest()
		return nil
	}
	return fmt.Errorf("unknown TestModel field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *TestModelMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *TestModelMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *TestModelMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *TestModelMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *TestModelMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *TestModelMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *TestModelMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown TestModel unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *TestModelMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown TestModel edge %s", name)
}
