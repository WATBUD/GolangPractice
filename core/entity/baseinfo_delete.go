// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"mai.today/core/entity/baseinfo"
	"mai.today/core/entity/predicate"
)

// BaseInfoDelete is the builder for deleting a BaseInfo entity.
type BaseInfoDelete struct {
	config
	hooks    []Hook
	mutation *BaseInfoMutation
}

// Where appends a list predicates to the BaseInfoDelete builder.
func (bid *BaseInfoDelete) Where(ps ...predicate.BaseInfo) *BaseInfoDelete {
	bid.mutation.Where(ps...)
	return bid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bid *BaseInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bid.sqlExec, bid.mutation, bid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bid *BaseInfoDelete) ExecX(ctx context.Context) int {
	n, err := bid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bid *BaseInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(baseinfo.Table, sqlgraph.NewFieldSpec(baseinfo.FieldID, field.TypeString))
	if ps := bid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bid.mutation.done = true
	return affected, err
}

// BaseInfoDeleteOne is the builder for deleting a single BaseInfo entity.
type BaseInfoDeleteOne struct {
	bid *BaseInfoDelete
}

// Where appends a list predicates to the BaseInfoDelete builder.
func (bido *BaseInfoDeleteOne) Where(ps ...predicate.BaseInfo) *BaseInfoDeleteOne {
	bido.bid.mutation.Where(ps...)
	return bido
}

// Exec executes the deletion query.
func (bido *BaseInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := bido.bid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baseinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bido *BaseInfoDeleteOne) ExecX(ctx context.Context) {
	if err := bido.Exec(ctx); err != nil {
		panic(err)
	}
}
