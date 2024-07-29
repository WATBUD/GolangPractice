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
	"mai.today/core/entity/basemember"
	"mai.today/core/entity/predicate"
)

// BaseMemberUpdate is the builder for updating BaseMember entities.
type BaseMemberUpdate struct {
	config
	hooks    []Hook
	mutation *BaseMemberMutation
}

// Where appends a list predicates to the BaseMemberUpdate builder.
func (bmu *BaseMemberUpdate) Where(ps ...predicate.BaseMember) *BaseMemberUpdate {
	bmu.mutation.Where(ps...)
	return bmu
}

// SetCreatedAt sets the "created_at" field.
func (bmu *BaseMemberUpdate) SetCreatedAt(t time.Time) *BaseMemberUpdate {
	bmu.mutation.SetCreatedAt(t)
	return bmu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bmu *BaseMemberUpdate) SetNillableCreatedAt(t *time.Time) *BaseMemberUpdate {
	if t != nil {
		bmu.SetCreatedAt(*t)
	}
	return bmu
}

// SetDeletedAt sets the "deleted_at" field.
func (bmu *BaseMemberUpdate) SetDeletedAt(t time.Time) *BaseMemberUpdate {
	bmu.mutation.SetDeletedAt(t)
	return bmu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bmu *BaseMemberUpdate) SetNillableDeletedAt(t *time.Time) *BaseMemberUpdate {
	if t != nil {
		bmu.SetDeletedAt(*t)
	}
	return bmu
}

// SetUpdatedAt sets the "updated_at" field.
func (bmu *BaseMemberUpdate) SetUpdatedAt(t time.Time) *BaseMemberUpdate {
	bmu.mutation.SetUpdatedAt(t)
	return bmu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bmu *BaseMemberUpdate) SetNillableUpdatedAt(t *time.Time) *BaseMemberUpdate {
	if t != nil {
		bmu.SetUpdatedAt(*t)
	}
	return bmu
}

// SetBaseID sets the "base_id" field.
func (bmu *BaseMemberUpdate) SetBaseID(s string) *BaseMemberUpdate {
	bmu.mutation.SetBaseID(s)
	return bmu
}

// SetNillableBaseID sets the "base_id" field if the given value is not nil.
func (bmu *BaseMemberUpdate) SetNillableBaseID(s *string) *BaseMemberUpdate {
	if s != nil {
		bmu.SetBaseID(*s)
	}
	return bmu
}

// SetUserID sets the "user_id" field.
func (bmu *BaseMemberUpdate) SetUserID(s string) *BaseMemberUpdate {
	bmu.mutation.SetUserID(s)
	return bmu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (bmu *BaseMemberUpdate) SetNillableUserID(s *string) *BaseMemberUpdate {
	if s != nil {
		bmu.SetUserID(*s)
	}
	return bmu
}

// Mutation returns the BaseMemberMutation object of the builder.
func (bmu *BaseMemberUpdate) Mutation() *BaseMemberMutation {
	return bmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bmu *BaseMemberUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bmu.sqlSave, bmu.mutation, bmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bmu *BaseMemberUpdate) SaveX(ctx context.Context) int {
	affected, err := bmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bmu *BaseMemberUpdate) Exec(ctx context.Context) error {
	_, err := bmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmu *BaseMemberUpdate) ExecX(ctx context.Context) {
	if err := bmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bmu *BaseMemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(basemember.Table, basemember.Columns, sqlgraph.NewFieldSpec(basemember.FieldID, field.TypeString))
	if ps := bmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bmu.mutation.CreatedAt(); ok {
		_spec.SetField(basemember.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bmu.mutation.DeletedAt(); ok {
		_spec.SetField(basemember.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := bmu.mutation.UpdatedAt(); ok {
		_spec.SetField(basemember.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bmu.mutation.BaseID(); ok {
		_spec.SetField(basemember.FieldBaseID, field.TypeString, value)
	}
	if value, ok := bmu.mutation.UserID(); ok {
		_spec.SetField(basemember.FieldUserID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{basemember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bmu.mutation.done = true
	return n, nil
}

// BaseMemberUpdateOne is the builder for updating a single BaseMember entity.
type BaseMemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BaseMemberMutation
}

// SetCreatedAt sets the "created_at" field.
func (bmuo *BaseMemberUpdateOne) SetCreatedAt(t time.Time) *BaseMemberUpdateOne {
	bmuo.mutation.SetCreatedAt(t)
	return bmuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bmuo *BaseMemberUpdateOne) SetNillableCreatedAt(t *time.Time) *BaseMemberUpdateOne {
	if t != nil {
		bmuo.SetCreatedAt(*t)
	}
	return bmuo
}

// SetDeletedAt sets the "deleted_at" field.
func (bmuo *BaseMemberUpdateOne) SetDeletedAt(t time.Time) *BaseMemberUpdateOne {
	bmuo.mutation.SetDeletedAt(t)
	return bmuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bmuo *BaseMemberUpdateOne) SetNillableDeletedAt(t *time.Time) *BaseMemberUpdateOne {
	if t != nil {
		bmuo.SetDeletedAt(*t)
	}
	return bmuo
}

// SetUpdatedAt sets the "updated_at" field.
func (bmuo *BaseMemberUpdateOne) SetUpdatedAt(t time.Time) *BaseMemberUpdateOne {
	bmuo.mutation.SetUpdatedAt(t)
	return bmuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bmuo *BaseMemberUpdateOne) SetNillableUpdatedAt(t *time.Time) *BaseMemberUpdateOne {
	if t != nil {
		bmuo.SetUpdatedAt(*t)
	}
	return bmuo
}

// SetBaseID sets the "base_id" field.
func (bmuo *BaseMemberUpdateOne) SetBaseID(s string) *BaseMemberUpdateOne {
	bmuo.mutation.SetBaseID(s)
	return bmuo
}

// SetNillableBaseID sets the "base_id" field if the given value is not nil.
func (bmuo *BaseMemberUpdateOne) SetNillableBaseID(s *string) *BaseMemberUpdateOne {
	if s != nil {
		bmuo.SetBaseID(*s)
	}
	return bmuo
}

// SetUserID sets the "user_id" field.
func (bmuo *BaseMemberUpdateOne) SetUserID(s string) *BaseMemberUpdateOne {
	bmuo.mutation.SetUserID(s)
	return bmuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (bmuo *BaseMemberUpdateOne) SetNillableUserID(s *string) *BaseMemberUpdateOne {
	if s != nil {
		bmuo.SetUserID(*s)
	}
	return bmuo
}

// Mutation returns the BaseMemberMutation object of the builder.
func (bmuo *BaseMemberUpdateOne) Mutation() *BaseMemberMutation {
	return bmuo.mutation
}

// Where appends a list predicates to the BaseMemberUpdate builder.
func (bmuo *BaseMemberUpdateOne) Where(ps ...predicate.BaseMember) *BaseMemberUpdateOne {
	bmuo.mutation.Where(ps...)
	return bmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bmuo *BaseMemberUpdateOne) Select(field string, fields ...string) *BaseMemberUpdateOne {
	bmuo.fields = append([]string{field}, fields...)
	return bmuo
}

// Save executes the query and returns the updated BaseMember entity.
func (bmuo *BaseMemberUpdateOne) Save(ctx context.Context) (*BaseMember, error) {
	return withHooks(ctx, bmuo.sqlSave, bmuo.mutation, bmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bmuo *BaseMemberUpdateOne) SaveX(ctx context.Context) *BaseMember {
	node, err := bmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bmuo *BaseMemberUpdateOne) Exec(ctx context.Context) error {
	_, err := bmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmuo *BaseMemberUpdateOne) ExecX(ctx context.Context) {
	if err := bmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bmuo *BaseMemberUpdateOne) sqlSave(ctx context.Context) (_node *BaseMember, err error) {
	_spec := sqlgraph.NewUpdateSpec(basemember.Table, basemember.Columns, sqlgraph.NewFieldSpec(basemember.FieldID, field.TypeString))
	id, ok := bmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entity: missing "BaseMember.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, basemember.FieldID)
		for _, f := range fields {
			if !basemember.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entity: invalid field %q for query", f)}
			}
			if f != basemember.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bmuo.mutation.CreatedAt(); ok {
		_spec.SetField(basemember.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bmuo.mutation.DeletedAt(); ok {
		_spec.SetField(basemember.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := bmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(basemember.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bmuo.mutation.BaseID(); ok {
		_spec.SetField(basemember.FieldBaseID, field.TypeString, value)
	}
	if value, ok := bmuo.mutation.UserID(); ok {
		_spec.SetField(basemember.FieldUserID, field.TypeString, value)
	}
	_node = &BaseMember{config: bmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{basemember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bmuo.mutation.done = true
	return _node, nil
}