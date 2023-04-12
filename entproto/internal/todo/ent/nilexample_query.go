// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/contrib/entproto/internal/todo/ent/nilexample"
	"entgo.io/contrib/entproto/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NilExampleQuery is the builder for querying NilExample entities.
type NilExampleQuery struct {
	config
	ctx        *QueryContext
	order      []nilexample.Order
	inters     []Interceptor
	predicates []predicate.NilExample
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NilExampleQuery builder.
func (neq *NilExampleQuery) Where(ps ...predicate.NilExample) *NilExampleQuery {
	neq.predicates = append(neq.predicates, ps...)
	return neq
}

// Limit the number of records to be returned by this query.
func (neq *NilExampleQuery) Limit(limit int) *NilExampleQuery {
	neq.ctx.Limit = &limit
	return neq
}

// Offset to start from.
func (neq *NilExampleQuery) Offset(offset int) *NilExampleQuery {
	neq.ctx.Offset = &offset
	return neq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (neq *NilExampleQuery) Unique(unique bool) *NilExampleQuery {
	neq.ctx.Unique = &unique
	return neq
}

// Order specifies how the records should be ordered.
func (neq *NilExampleQuery) Order(o ...nilexample.Order) *NilExampleQuery {
	neq.order = append(neq.order, o...)
	return neq
}

// First returns the first NilExample entity from the query.
// Returns a *NotFoundError when no NilExample was found.
func (neq *NilExampleQuery) First(ctx context.Context) (*NilExample, error) {
	nodes, err := neq.Limit(1).All(setContextOp(ctx, neq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{nilexample.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (neq *NilExampleQuery) FirstX(ctx context.Context) *NilExample {
	node, err := neq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NilExample ID from the query.
// Returns a *NotFoundError when no NilExample ID was found.
func (neq *NilExampleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = neq.Limit(1).IDs(setContextOp(ctx, neq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nilexample.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (neq *NilExampleQuery) FirstIDX(ctx context.Context) int {
	id, err := neq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NilExample entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NilExample entity is found.
// Returns a *NotFoundError when no NilExample entities are found.
func (neq *NilExampleQuery) Only(ctx context.Context) (*NilExample, error) {
	nodes, err := neq.Limit(2).All(setContextOp(ctx, neq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{nilexample.Label}
	default:
		return nil, &NotSingularError{nilexample.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (neq *NilExampleQuery) OnlyX(ctx context.Context) *NilExample {
	node, err := neq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NilExample ID in the query.
// Returns a *NotSingularError when more than one NilExample ID is found.
// Returns a *NotFoundError when no entities are found.
func (neq *NilExampleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = neq.Limit(2).IDs(setContextOp(ctx, neq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nilexample.Label}
	default:
		err = &NotSingularError{nilexample.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (neq *NilExampleQuery) OnlyIDX(ctx context.Context) int {
	id, err := neq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NilExamples.
func (neq *NilExampleQuery) All(ctx context.Context) ([]*NilExample, error) {
	ctx = setContextOp(ctx, neq.ctx, "All")
	if err := neq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NilExample, *NilExampleQuery]()
	return withInterceptors[[]*NilExample](ctx, neq, qr, neq.inters)
}

// AllX is like All, but panics if an error occurs.
func (neq *NilExampleQuery) AllX(ctx context.Context) []*NilExample {
	nodes, err := neq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NilExample IDs.
func (neq *NilExampleQuery) IDs(ctx context.Context) (ids []int, err error) {
	if neq.ctx.Unique == nil && neq.path != nil {
		neq.Unique(true)
	}
	ctx = setContextOp(ctx, neq.ctx, "IDs")
	if err = neq.Select(nilexample.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (neq *NilExampleQuery) IDsX(ctx context.Context) []int {
	ids, err := neq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (neq *NilExampleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, neq.ctx, "Count")
	if err := neq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, neq, querierCount[*NilExampleQuery](), neq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (neq *NilExampleQuery) CountX(ctx context.Context) int {
	count, err := neq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (neq *NilExampleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, neq.ctx, "Exist")
	switch _, err := neq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (neq *NilExampleQuery) ExistX(ctx context.Context) bool {
	exist, err := neq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NilExampleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (neq *NilExampleQuery) Clone() *NilExampleQuery {
	if neq == nil {
		return nil
	}
	return &NilExampleQuery{
		config:     neq.config,
		ctx:        neq.ctx.Clone(),
		order:      append([]nilexample.Order{}, neq.order...),
		inters:     append([]Interceptor{}, neq.inters...),
		predicates: append([]predicate.NilExample{}, neq.predicates...),
		// clone intermediate query.
		sql:  neq.sql.Clone(),
		path: neq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StrNil string `json:"str_nil,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NilExample.Query().
//		GroupBy(nilexample.FieldStrNil).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (neq *NilExampleQuery) GroupBy(field string, fields ...string) *NilExampleGroupBy {
	neq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NilExampleGroupBy{build: neq}
	grbuild.flds = &neq.ctx.Fields
	grbuild.label = nilexample.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		StrNil string `json:"str_nil,omitempty"`
//	}
//
//	client.NilExample.Query().
//		Select(nilexample.FieldStrNil).
//		Scan(ctx, &v)
func (neq *NilExampleQuery) Select(fields ...string) *NilExampleSelect {
	neq.ctx.Fields = append(neq.ctx.Fields, fields...)
	sbuild := &NilExampleSelect{NilExampleQuery: neq}
	sbuild.label = nilexample.Label
	sbuild.flds, sbuild.scan = &neq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NilExampleSelect configured with the given aggregations.
func (neq *NilExampleQuery) Aggregate(fns ...AggregateFunc) *NilExampleSelect {
	return neq.Select().Aggregate(fns...)
}

func (neq *NilExampleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range neq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, neq); err != nil {
				return err
			}
		}
	}
	for _, f := range neq.ctx.Fields {
		if !nilexample.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if neq.path != nil {
		prev, err := neq.path(ctx)
		if err != nil {
			return err
		}
		neq.sql = prev
	}
	return nil
}

func (neq *NilExampleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NilExample, error) {
	var (
		nodes = []*NilExample{}
		_spec = neq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NilExample).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NilExample{config: neq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, neq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (neq *NilExampleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := neq.querySpec()
	_spec.Node.Columns = neq.ctx.Fields
	if len(neq.ctx.Fields) > 0 {
		_spec.Unique = neq.ctx.Unique != nil && *neq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, neq.driver, _spec)
}

func (neq *NilExampleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(nilexample.Table, nilexample.Columns, sqlgraph.NewFieldSpec(nilexample.FieldID, field.TypeInt))
	_spec.From = neq.sql
	if unique := neq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if neq.path != nil {
		_spec.Unique = true
	}
	if fields := neq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nilexample.FieldID)
		for i := range fields {
			if fields[i] != nilexample.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := neq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := neq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := neq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := neq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (neq *NilExampleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(neq.driver.Dialect())
	t1 := builder.Table(nilexample.Table)
	columns := neq.ctx.Fields
	if len(columns) == 0 {
		columns = nilexample.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if neq.sql != nil {
		selector = neq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if neq.ctx.Unique != nil && *neq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range neq.predicates {
		p(selector)
	}
	for _, p := range neq.order {
		p(selector)
	}
	if offset := neq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := neq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NilExampleGroupBy is the group-by builder for NilExample entities.
type NilExampleGroupBy struct {
	selector
	build *NilExampleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (negb *NilExampleGroupBy) Aggregate(fns ...AggregateFunc) *NilExampleGroupBy {
	negb.fns = append(negb.fns, fns...)
	return negb
}

// Scan applies the selector query and scans the result into the given value.
func (negb *NilExampleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, negb.build.ctx, "GroupBy")
	if err := negb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NilExampleQuery, *NilExampleGroupBy](ctx, negb.build, negb, negb.build.inters, v)
}

func (negb *NilExampleGroupBy) sqlScan(ctx context.Context, root *NilExampleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(negb.fns))
	for _, fn := range negb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*negb.flds)+len(negb.fns))
		for _, f := range *negb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*negb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := negb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NilExampleSelect is the builder for selecting fields of NilExample entities.
type NilExampleSelect struct {
	*NilExampleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nes *NilExampleSelect) Aggregate(fns ...AggregateFunc) *NilExampleSelect {
	nes.fns = append(nes.fns, fns...)
	return nes
}

// Scan applies the selector query and scans the result into the given value.
func (nes *NilExampleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nes.ctx, "Select")
	if err := nes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NilExampleQuery, *NilExampleSelect](ctx, nes.NilExampleQuery, nes, nes.inters, v)
}

func (nes *NilExampleSelect) sqlScan(ctx context.Context, root *NilExampleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nes.fns))
	for _, fn := range nes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
