// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"mai.today/core/entity/basenavstate"
	"mai.today/core/entity/predicate"
)

// BaseNavStateQuery is the builder for querying BaseNavState entities.
type BaseNavStateQuery struct {
	config
	ctx        *QueryContext
	order      []basenavstate.OrderOption
	inters     []Interceptor
	predicates []predicate.BaseNavState
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaseNavStateQuery builder.
func (bnsq *BaseNavStateQuery) Where(ps ...predicate.BaseNavState) *BaseNavStateQuery {
	bnsq.predicates = append(bnsq.predicates, ps...)
	return bnsq
}

// Limit the number of records to be returned by this query.
func (bnsq *BaseNavStateQuery) Limit(limit int) *BaseNavStateQuery {
	bnsq.ctx.Limit = &limit
	return bnsq
}

// Offset to start from.
func (bnsq *BaseNavStateQuery) Offset(offset int) *BaseNavStateQuery {
	bnsq.ctx.Offset = &offset
	return bnsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bnsq *BaseNavStateQuery) Unique(unique bool) *BaseNavStateQuery {
	bnsq.ctx.Unique = &unique
	return bnsq
}

// Order specifies how the records should be ordered.
func (bnsq *BaseNavStateQuery) Order(o ...basenavstate.OrderOption) *BaseNavStateQuery {
	bnsq.order = append(bnsq.order, o...)
	return bnsq
}

// First returns the first BaseNavState entity from the query.
// Returns a *NotFoundError when no BaseNavState was found.
func (bnsq *BaseNavStateQuery) First(ctx context.Context) (*BaseNavState, error) {
	nodes, err := bnsq.Limit(1).All(setContextOp(ctx, bnsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{basenavstate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bnsq *BaseNavStateQuery) FirstX(ctx context.Context) *BaseNavState {
	node, err := bnsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaseNavState ID from the query.
// Returns a *NotFoundError when no BaseNavState ID was found.
func (bnsq *BaseNavStateQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bnsq.Limit(1).IDs(setContextOp(ctx, bnsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{basenavstate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bnsq *BaseNavStateQuery) FirstIDX(ctx context.Context) string {
	id, err := bnsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaseNavState entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaseNavState entity is found.
// Returns a *NotFoundError when no BaseNavState entities are found.
func (bnsq *BaseNavStateQuery) Only(ctx context.Context) (*BaseNavState, error) {
	nodes, err := bnsq.Limit(2).All(setContextOp(ctx, bnsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{basenavstate.Label}
	default:
		return nil, &NotSingularError{basenavstate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bnsq *BaseNavStateQuery) OnlyX(ctx context.Context) *BaseNavState {
	node, err := bnsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaseNavState ID in the query.
// Returns a *NotSingularError when more than one BaseNavState ID is found.
// Returns a *NotFoundError when no entities are found.
func (bnsq *BaseNavStateQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bnsq.Limit(2).IDs(setContextOp(ctx, bnsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{basenavstate.Label}
	default:
		err = &NotSingularError{basenavstate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bnsq *BaseNavStateQuery) OnlyIDX(ctx context.Context) string {
	id, err := bnsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaseNavStates.
func (bnsq *BaseNavStateQuery) All(ctx context.Context) ([]*BaseNavState, error) {
	ctx = setContextOp(ctx, bnsq.ctx, "All")
	if err := bnsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BaseNavState, *BaseNavStateQuery]()
	return withInterceptors[[]*BaseNavState](ctx, bnsq, qr, bnsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bnsq *BaseNavStateQuery) AllX(ctx context.Context) []*BaseNavState {
	nodes, err := bnsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaseNavState IDs.
func (bnsq *BaseNavStateQuery) IDs(ctx context.Context) (ids []string, err error) {
	if bnsq.ctx.Unique == nil && bnsq.path != nil {
		bnsq.Unique(true)
	}
	ctx = setContextOp(ctx, bnsq.ctx, "IDs")
	if err = bnsq.Select(basenavstate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bnsq *BaseNavStateQuery) IDsX(ctx context.Context) []string {
	ids, err := bnsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bnsq *BaseNavStateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bnsq.ctx, "Count")
	if err := bnsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bnsq, querierCount[*BaseNavStateQuery](), bnsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bnsq *BaseNavStateQuery) CountX(ctx context.Context) int {
	count, err := bnsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bnsq *BaseNavStateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bnsq.ctx, "Exist")
	switch _, err := bnsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entity: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bnsq *BaseNavStateQuery) ExistX(ctx context.Context) bool {
	exist, err := bnsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaseNavStateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bnsq *BaseNavStateQuery) Clone() *BaseNavStateQuery {
	if bnsq == nil {
		return nil
	}
	return &BaseNavStateQuery{
		config:     bnsq.config,
		ctx:        bnsq.ctx.Clone(),
		order:      append([]basenavstate.OrderOption{}, bnsq.order...),
		inters:     append([]Interceptor{}, bnsq.inters...),
		predicates: append([]predicate.BaseNavState{}, bnsq.predicates...),
		// clone intermediate query.
		sql:  bnsq.sql.Clone(),
		path: bnsq.path,
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
//	client.BaseNavState.Query().
//		GroupBy(basenavstate.FieldCreatedAt).
//		Aggregate(entity.Count()).
//		Scan(ctx, &v)
func (bnsq *BaseNavStateQuery) GroupBy(field string, fields ...string) *BaseNavStateGroupBy {
	bnsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BaseNavStateGroupBy{build: bnsq}
	grbuild.flds = &bnsq.ctx.Fields
	grbuild.label = basenavstate.Label
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
//	client.BaseNavState.Query().
//		Select(basenavstate.FieldCreatedAt).
//		Scan(ctx, &v)
func (bnsq *BaseNavStateQuery) Select(fields ...string) *BaseNavStateSelect {
	bnsq.ctx.Fields = append(bnsq.ctx.Fields, fields...)
	sbuild := &BaseNavStateSelect{BaseNavStateQuery: bnsq}
	sbuild.label = basenavstate.Label
	sbuild.flds, sbuild.scan = &bnsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BaseNavStateSelect configured with the given aggregations.
func (bnsq *BaseNavStateQuery) Aggregate(fns ...AggregateFunc) *BaseNavStateSelect {
	return bnsq.Select().Aggregate(fns...)
}

func (bnsq *BaseNavStateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bnsq.inters {
		if inter == nil {
			return fmt.Errorf("entity: uninitialized interceptor (forgotten import entity/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bnsq); err != nil {
				return err
			}
		}
	}
	for _, f := range bnsq.ctx.Fields {
		if !basenavstate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entity: invalid field %q for query", f)}
		}
	}
	if bnsq.path != nil {
		prev, err := bnsq.path(ctx)
		if err != nil {
			return err
		}
		bnsq.sql = prev
	}
	return nil
}

func (bnsq *BaseNavStateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaseNavState, error) {
	var (
		nodes = []*BaseNavState{}
		_spec = bnsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BaseNavState).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BaseNavState{config: bnsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bnsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bnsq *BaseNavStateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bnsq.querySpec()
	_spec.Node.Columns = bnsq.ctx.Fields
	if len(bnsq.ctx.Fields) > 0 {
		_spec.Unique = bnsq.ctx.Unique != nil && *bnsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bnsq.driver, _spec)
}

func (bnsq *BaseNavStateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(basenavstate.Table, basenavstate.Columns, sqlgraph.NewFieldSpec(basenavstate.FieldID, field.TypeString))
	_spec.From = bnsq.sql
	if unique := bnsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bnsq.path != nil {
		_spec.Unique = true
	}
	if fields := bnsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, basenavstate.FieldID)
		for i := range fields {
			if fields[i] != basenavstate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bnsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bnsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bnsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bnsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bnsq *BaseNavStateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bnsq.driver.Dialect())
	t1 := builder.Table(basenavstate.Table)
	columns := bnsq.ctx.Fields
	if len(columns) == 0 {
		columns = basenavstate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bnsq.sql != nil {
		selector = bnsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bnsq.ctx.Unique != nil && *bnsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bnsq.predicates {
		p(selector)
	}
	for _, p := range bnsq.order {
		p(selector)
	}
	if offset := bnsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bnsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaseNavStateGroupBy is the group-by builder for BaseNavState entities.
type BaseNavStateGroupBy struct {
	selector
	build *BaseNavStateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bnsgb *BaseNavStateGroupBy) Aggregate(fns ...AggregateFunc) *BaseNavStateGroupBy {
	bnsgb.fns = append(bnsgb.fns, fns...)
	return bnsgb
}

// Scan applies the selector query and scans the result into the given value.
func (bnsgb *BaseNavStateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bnsgb.build.ctx, "GroupBy")
	if err := bnsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BaseNavStateQuery, *BaseNavStateGroupBy](ctx, bnsgb.build, bnsgb, bnsgb.build.inters, v)
}

func (bnsgb *BaseNavStateGroupBy) sqlScan(ctx context.Context, root *BaseNavStateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bnsgb.fns))
	for _, fn := range bnsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bnsgb.flds)+len(bnsgb.fns))
		for _, f := range *bnsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bnsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bnsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BaseNavStateSelect is the builder for selecting fields of BaseNavState entities.
type BaseNavStateSelect struct {
	*BaseNavStateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bnss *BaseNavStateSelect) Aggregate(fns ...AggregateFunc) *BaseNavStateSelect {
	bnss.fns = append(bnss.fns, fns...)
	return bnss
}

// Scan applies the selector query and scans the result into the given value.
func (bnss *BaseNavStateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bnss.ctx, "Select")
	if err := bnss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BaseNavStateQuery, *BaseNavStateSelect](ctx, bnss.BaseNavStateQuery, bnss, bnss.inters, v)
}

func (bnss *BaseNavStateSelect) sqlScan(ctx context.Context, root *BaseNavStateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bnss.fns))
	for _, fn := range bnss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bnss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bnss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
