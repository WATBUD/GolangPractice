// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"mai.today/core/entity/base"
	"mai.today/core/entity/predicate"
)

// BaseQuery is the builder for querying Base entities.
type BaseQuery struct {
	config
	ctx        *QueryContext
	order      []base.OrderOption
	inters     []Interceptor
	predicates []predicate.Base
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaseQuery builder.
func (bq *BaseQuery) Where(ps ...predicate.Base) *BaseQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BaseQuery) Limit(limit int) *BaseQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BaseQuery) Offset(offset int) *BaseQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BaseQuery) Unique(unique bool) *BaseQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BaseQuery) Order(o ...base.OrderOption) *BaseQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// First returns the first Base entity from the query.
// Returns a *NotFoundError when no Base was found.
func (bq *BaseQuery) First(ctx context.Context) (*Base, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{base.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BaseQuery) FirstX(ctx context.Context) *Base {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Base ID from the query.
// Returns a *NotFoundError when no Base ID was found.
func (bq *BaseQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{base.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BaseQuery) FirstIDX(ctx context.Context) string {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Base entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Base entity is found.
// Returns a *NotFoundError when no Base entities are found.
func (bq *BaseQuery) Only(ctx context.Context) (*Base, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{base.Label}
	default:
		return nil, &NotSingularError{base.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BaseQuery) OnlyX(ctx context.Context) *Base {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Base ID in the query.
// Returns a *NotSingularError when more than one Base ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BaseQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{base.Label}
	default:
		err = &NotSingularError{base.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BaseQuery) OnlyIDX(ctx context.Context) string {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Bases.
func (bq *BaseQuery) All(ctx context.Context) ([]*Base, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Base, *BaseQuery]()
	return withInterceptors[[]*Base](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BaseQuery) AllX(ctx context.Context) []*Base {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Base IDs.
func (bq *BaseQuery) IDs(ctx context.Context) (ids []string, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(base.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BaseQuery) IDsX(ctx context.Context) []string {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BaseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BaseQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BaseQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BaseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entity: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BaseQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BaseQuery) Clone() *BaseQuery {
	if bq == nil {
		return nil
	}
	return &BaseQuery{
		config:     bq.config,
		ctx:        bq.ctx.Clone(),
		order:      append([]base.OrderOption{}, bq.order...),
		inters:     append([]Interceptor{}, bq.inters...),
		predicates: append([]predicate.Base{}, bq.predicates...),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Base.Query().
//		GroupBy(base.FieldCreatedAt).
//		Aggregate(entity.Count()).
//		Scan(ctx, &v)
func (bq *BaseQuery) GroupBy(field string, fields ...string) *BaseGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BaseGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = base.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Base.Query().
//		Select(base.FieldCreatedAt).
//		Scan(ctx, &v)
func (bq *BaseQuery) Select(fields ...string) *BaseSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BaseSelect{BaseQuery: bq}
	sbuild.label = base.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BaseSelect configured with the given aggregations.
func (bq *BaseQuery) Aggregate(fns ...AggregateFunc) *BaseSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BaseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("entity: uninitialized interceptor (forgotten import entity/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !base.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entity: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Base, error) {
	var (
		nodes = []*Base{}
		_spec = bq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Base).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Base{config: bq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bq *BaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(base.Table, base.Columns, sqlgraph.NewFieldSpec(base.FieldID, field.TypeString))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, base.FieldID)
		for i := range fields {
			if fields[i] != base.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(base.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = base.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaseGroupBy is the group-by builder for Base entities.
type BaseGroupBy struct {
	selector
	build *BaseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BaseGroupBy) Aggregate(fns ...AggregateFunc) *BaseGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BaseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BaseQuery, *BaseGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BaseGroupBy) sqlScan(ctx context.Context, root *BaseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BaseSelect is the builder for selecting fields of Base entities.
type BaseSelect struct {
	*BaseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BaseSelect) Aggregate(fns ...AggregateFunc) *BaseSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BaseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BaseQuery, *BaseSelect](ctx, bs.BaseQuery, bs, bs.inters, v)
}

func (bs *BaseSelect) sqlScan(ctx context.Context, root *BaseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}