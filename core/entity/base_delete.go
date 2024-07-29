// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"mai.today/core/entity/base"
	"mai.today/core/entity/predicate"
)

// BaseDelete is the builder for deleting a Base entity.
type BaseDelete struct {
	config
	hooks    []Hook
	mutation *BaseMutation
}

// Where appends a list predicates to the BaseDelete builder.
func (bd *BaseDelete) Where(ps ...predicate.Base) *BaseDelete {
	bd.mutation.Where(ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BaseDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bd.sqlExec, bd.mutation, bd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BaseDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BaseDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(base.Table, sqlgraph.NewFieldSpec(base.FieldID, field.TypeString))
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bd.mutation.done = true
	return affected, err
}

// BaseDeleteOne is the builder for deleting a single Base entity.
type BaseDeleteOne struct {
	bd *BaseDelete
}

// Where appends a list predicates to the BaseDelete builder.
func (bdo *BaseDeleteOne) Where(ps ...predicate.Base) *BaseDeleteOne {
	bdo.bd.mutation.Where(ps...)
	return bdo
}

// Exec executes the deletion query.
func (bdo *BaseDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{base.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BaseDeleteOne) ExecX(ctx context.Context) {
	if err := bdo.Exec(ctx); err != nil {
		panic(err)
	}
}