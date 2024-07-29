// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"mai.today/core/entity/basenavstate"
	"mai.today/core/entity/predicate"
)

// BaseNavStateDelete is the builder for deleting a BaseNavState entity.
type BaseNavStateDelete struct {
	config
	hooks    []Hook
	mutation *BaseNavStateMutation
}

// Where appends a list predicates to the BaseNavStateDelete builder.
func (bnsd *BaseNavStateDelete) Where(ps ...predicate.BaseNavState) *BaseNavStateDelete {
	bnsd.mutation.Where(ps...)
	return bnsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bnsd *BaseNavStateDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bnsd.sqlExec, bnsd.mutation, bnsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bnsd *BaseNavStateDelete) ExecX(ctx context.Context) int {
	n, err := bnsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bnsd *BaseNavStateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(basenavstate.Table, sqlgraph.NewFieldSpec(basenavstate.FieldID, field.TypeString))
	if ps := bnsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bnsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bnsd.mutation.done = true
	return affected, err
}

// BaseNavStateDeleteOne is the builder for deleting a single BaseNavState entity.
type BaseNavStateDeleteOne struct {
	bnsd *BaseNavStateDelete
}

// Where appends a list predicates to the BaseNavStateDelete builder.
func (bnsdo *BaseNavStateDeleteOne) Where(ps ...predicate.BaseNavState) *BaseNavStateDeleteOne {
	bnsdo.bnsd.mutation.Where(ps...)
	return bnsdo
}

// Exec executes the deletion query.
func (bnsdo *BaseNavStateDeleteOne) Exec(ctx context.Context) error {
	n, err := bnsdo.bnsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{basenavstate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bnsdo *BaseNavStateDeleteOne) ExecX(ctx context.Context) {
	if err := bnsdo.Exec(ctx); err != nil {
		panic(err)
	}
}