// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"mai.today/core/entity/basenavstate"
	"mai.today/core/entity/predicate"
)

// BaseNavStateUpdate is the builder for updating BaseNavState entities.
type BaseNavStateUpdate struct {
	config
	hooks    []Hook
	mutation *BaseNavStateMutation
}

// Where appends a list predicates to the BaseNavStateUpdate builder.
func (bnsu *BaseNavStateUpdate) Where(ps ...predicate.BaseNavState) *BaseNavStateUpdate {
	bnsu.mutation.Where(ps...)
	return bnsu
}

// SetCreatedAt sets the "created_at" field.
func (bnsu *BaseNavStateUpdate) SetCreatedAt(t time.Time) *BaseNavStateUpdate {
	bnsu.mutation.SetCreatedAt(t)
	return bnsu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bnsu *BaseNavStateUpdate) SetNillableCreatedAt(t *time.Time) *BaseNavStateUpdate {
	if t != nil {
		bnsu.SetCreatedAt(*t)
	}
	return bnsu
}

// SetDeletedAt sets the "deleted_at" field.
func (bnsu *BaseNavStateUpdate) SetDeletedAt(t time.Time) *BaseNavStateUpdate {
	bnsu.mutation.SetDeletedAt(t)
	return bnsu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bnsu *BaseNavStateUpdate) SetNillableDeletedAt(t *time.Time) *BaseNavStateUpdate {
	if t != nil {
		bnsu.SetDeletedAt(*t)
	}
	return bnsu
}

// SetUpdatedAt sets the "updated_at" field.
func (bnsu *BaseNavStateUpdate) SetUpdatedAt(t time.Time) *BaseNavStateUpdate {
	bnsu.mutation.SetUpdatedAt(t)
	return bnsu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bnsu *BaseNavStateUpdate) SetNillableUpdatedAt(t *time.Time) *BaseNavStateUpdate {
	if t != nil {
		bnsu.SetUpdatedAt(*t)
	}
	return bnsu
}

// SetBaseID sets the "base_id" field.
func (bnsu *BaseNavStateUpdate) SetBaseID(s string) *BaseNavStateUpdate {
	bnsu.mutation.SetBaseID(s)
	return bnsu
}

// SetNillableBaseID sets the "base_id" field if the given value is not nil.
func (bnsu *BaseNavStateUpdate) SetNillableBaseID(s *string) *BaseNavStateUpdate {
	if s != nil {
		bnsu.SetBaseID(*s)
	}
	return bnsu
}

// SetUserID sets the "user_id" field.
func (bnsu *BaseNavStateUpdate) SetUserID(s string) *BaseNavStateUpdate {
	bnsu.mutation.SetUserID(s)
	return bnsu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (bnsu *BaseNavStateUpdate) SetNillableUserID(s *string) *BaseNavStateUpdate {
	if s != nil {
		bnsu.SetUserID(*s)
	}
	return bnsu
}

// SetIndex sets the "index" field.
func (bnsu *BaseNavStateUpdate) SetIndex(i int) *BaseNavStateUpdate {
	bnsu.mutation.ResetIndex()
	bnsu.mutation.SetIndex(i)
	return bnsu
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (bnsu *BaseNavStateUpdate) SetNillableIndex(i *int) *BaseNavStateUpdate {
	if i != nil {
		bnsu.SetIndex(*i)
	}
	return bnsu
}

// AddIndex adds i to the "index" field.
func (bnsu *BaseNavStateUpdate) AddIndex(i int) *BaseNavStateUpdate {
	bnsu.mutation.AddIndex(i)
	return bnsu
}

// Mutation returns the BaseNavStateMutation object of the builder.
func (bnsu *BaseNavStateUpdate) Mutation() *BaseNavStateMutation {
	return bnsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bnsu *BaseNavStateUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bnsu.sqlSave, bnsu.mutation, bnsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bnsu *BaseNavStateUpdate) SaveX(ctx context.Context) int {
	affected, err := bnsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bnsu *BaseNavStateUpdate) Exec(ctx context.Context) error {
	_, err := bnsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bnsu *BaseNavStateUpdate) ExecX(ctx context.Context) {
	if err := bnsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bnsu *BaseNavStateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(basenavstate.Table, basenavstate.Columns, sqlgraph.NewFieldSpec(basenavstate.FieldID, field.TypeString))
	if ps := bnsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bnsu.mutation.CreatedAt(); ok {
		_spec.SetField(basenavstate.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bnsu.mutation.DeletedAt(); ok {
		_spec.SetField(basenavstate.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := bnsu.mutation.UpdatedAt(); ok {
		_spec.SetField(basenavstate.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bnsu.mutation.BaseID(); ok {
		_spec.SetField(basenavstate.FieldBaseID, field.TypeString, value)
	}
	if value, ok := bnsu.mutation.UserID(); ok {
		_spec.SetField(basenavstate.FieldUserID, field.TypeString, value)
	}
	if value, ok := bnsu.mutation.Index(); ok {
		_spec.SetField(basenavstate.FieldIndex, field.TypeInt, value)
	}
	if value, ok := bnsu.mutation.AddedIndex(); ok {
		_spec.AddField(basenavstate.FieldIndex, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bnsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{basenavstate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bnsu.mutation.done = true
	return n, nil
}

// BaseNavStateUpdateOne is the builder for updating a single BaseNavState entity.
type BaseNavStateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaseNavStateMutation
}

// SetCreatedAt sets the "created_at" field.
func (bnsuo *BaseNavStateUpdateOne) SetCreatedAt(t time.Time) *BaseNavStateUpdateOne {
	bnsuo.mutation.SetCreatedAt(t)
	return bnsuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bnsuo *BaseNavStateUpdateOne) SetNillableCreatedAt(t *time.Time) *BaseNavStateUpdateOne {
	if t != nil {
		bnsuo.SetCreatedAt(*t)
	}
	return bnsuo
}

// SetDeletedAt sets the "deleted_at" field.
func (bnsuo *BaseNavStateUpdateOne) SetDeletedAt(t time.Time) *BaseNavStateUpdateOne {
	bnsuo.mutation.SetDeletedAt(t)
	return bnsuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bnsuo *BaseNavStateUpdateOne) SetNillableDeletedAt(t *time.Time) *BaseNavStateUpdateOne {
	if t != nil {
		bnsuo.SetDeletedAt(*t)
	}
	return bnsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (bnsuo *BaseNavStateUpdateOne) SetUpdatedAt(t time.Time) *BaseNavStateUpdateOne {
	bnsuo.mutation.SetUpdatedAt(t)
	return bnsuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bnsuo *BaseNavStateUpdateOne) SetNillableUpdatedAt(t *time.Time) *BaseNavStateUpdateOne {
	if t != nil {
		bnsuo.SetUpdatedAt(*t)
	}
	return bnsuo
}

// SetBaseID sets the "base_id" field.
func (bnsuo *BaseNavStateUpdateOne) SetBaseID(s string) *BaseNavStateUpdateOne {
	bnsuo.mutation.SetBaseID(s)
	return bnsuo
}

// SetNillableBaseID sets the "base_id" field if the given value is not nil.
func (bnsuo *BaseNavStateUpdateOne) SetNillableBaseID(s *string) *BaseNavStateUpdateOne {
	if s != nil {
		bnsuo.SetBaseID(*s)
	}
	return bnsuo
}

// SetUserID sets the "user_id" field.
func (bnsuo *BaseNavStateUpdateOne) SetUserID(s string) *BaseNavStateUpdateOne {
	bnsuo.mutation.SetUserID(s)
	return bnsuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (bnsuo *BaseNavStateUpdateOne) SetNillableUserID(s *string) *BaseNavStateUpdateOne {
	if s != nil {
		bnsuo.SetUserID(*s)
	}
	return bnsuo
}

// SetIndex sets the "index" field.
func (bnsuo *BaseNavStateUpdateOne) SetIndex(i int) *BaseNavStateUpdateOne {
	bnsuo.mutation.ResetIndex()
	bnsuo.mutation.SetIndex(i)
	return bnsuo
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (bnsuo *BaseNavStateUpdateOne) SetNillableIndex(i *int) *BaseNavStateUpdateOne {
	if i != nil {
		bnsuo.SetIndex(*i)
	}
	return bnsuo
}

// AddIndex adds i to the "index" field.
func (bnsuo *BaseNavStateUpdateOne) AddIndex(i int) *BaseNavStateUpdateOne {
	bnsuo.mutation.AddIndex(i)
	return bnsuo
}

// Mutation returns the BaseNavStateMutation object of the builder.
func (bnsuo *BaseNavStateUpdateOne) Mutation() *BaseNavStateMutation {
	return bnsuo.mutation
}

// Where appends a list predicates to the BaseNavStateUpdate builder.
func (bnsuo *BaseNavStateUpdateOne) Where(ps ...predicate.BaseNavState) *BaseNavStateUpdateOne {
	bnsuo.mutation.Where(ps...)
	return bnsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bnsuo *BaseNavStateUpdateOne) Select(field string, fields ...string) *BaseNavStateUpdateOne {
	bnsuo.fields = append([]string{field}, fields...)
	return bnsuo
}

// Save executes the query and returns the updated BaseNavState entity.
func (bnsuo *BaseNavStateUpdateOne) Save(ctx context.Context) (*BaseNavState, error) {
	return withHooks(ctx, bnsuo.sqlSave, bnsuo.mutation, bnsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bnsuo *BaseNavStateUpdateOne) SaveX(ctx context.Context) *BaseNavState {
	node, err := bnsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bnsuo *BaseNavStateUpdateOne) Exec(ctx context.Context) error {
	_, err := bnsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bnsuo *BaseNavStateUpdateOne) ExecX(ctx context.Context) {
	if err := bnsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bnsuo *BaseNavStateUpdateOne) sqlSave(ctx context.Context) (_node *BaseNavState, err error) {
	_spec := sqlgraph.NewUpdateSpec(basenavstate.Table, basenavstate.Columns, sqlgraph.NewFieldSpec(basenavstate.FieldID, field.TypeString))
	id, ok := bnsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entity: missing "BaseNavState.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bnsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, basenavstate.FieldID)
		for _, f := range fields {
			if !basenavstate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entity: invalid field %q for query", f)}
			}
			if f != basenavstate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bnsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bnsuo.mutation.CreatedAt(); ok {
		_spec.SetField(basenavstate.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bnsuo.mutation.DeletedAt(); ok {
		_spec.SetField(basenavstate.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := bnsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(basenavstate.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bnsuo.mutation.BaseID(); ok {
		_spec.SetField(basenavstate.FieldBaseID, field.TypeString, value)
	}
	if value, ok := bnsuo.mutation.UserID(); ok {
		_spec.SetField(basenavstate.FieldUserID, field.TypeString, value)
	}
	if value, ok := bnsuo.mutation.Index(); ok {
		_spec.SetField(basenavstate.FieldIndex, field.TypeInt, value)
	}
	if value, ok := bnsuo.mutation.AddedIndex(); ok {
		_spec.AddField(basenavstate.FieldIndex, field.TypeInt, value)
	}
	_node = &BaseNavState{config: bnsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bnsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{basenavstate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bnsuo.mutation.done = true
	return _node, nil
}
