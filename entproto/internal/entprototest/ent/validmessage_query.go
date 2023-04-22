// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/contrib/entproto/internal/entprototest/ent/validmessage"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ValidMessageQuery is the builder for querying ValidMessage entities.
type ValidMessageQuery struct {
	config
	ctx        *QueryContext
	order      []validmessage.OrderOption
	inters     []Interceptor
	predicates []predicate.ValidMessage
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ValidMessageQuery builder.
func (vmq *ValidMessageQuery) Where(ps ...predicate.ValidMessage) *ValidMessageQuery {
	vmq.predicates = append(vmq.predicates, ps...)
	return vmq
}

// Limit the number of records to be returned by this query.
func (vmq *ValidMessageQuery) Limit(limit int) *ValidMessageQuery {
	vmq.ctx.Limit = &limit
	return vmq
}

// Offset to start from.
func (vmq *ValidMessageQuery) Offset(offset int) *ValidMessageQuery {
	vmq.ctx.Offset = &offset
	return vmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vmq *ValidMessageQuery) Unique(unique bool) *ValidMessageQuery {
	vmq.ctx.Unique = &unique
	return vmq
}

// Order specifies how the records should be ordered.
func (vmq *ValidMessageQuery) Order(o ...validmessage.OrderOption) *ValidMessageQuery {
	vmq.order = append(vmq.order, o...)
	return vmq
}

// First returns the first ValidMessage entity from the query.
// Returns a *NotFoundError when no ValidMessage was found.
func (vmq *ValidMessageQuery) First(ctx context.Context) (*ValidMessage, error) {
	nodes, err := vmq.Limit(1).All(setContextOp(ctx, vmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{validmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vmq *ValidMessageQuery) FirstX(ctx context.Context) *ValidMessage {
	node, err := vmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ValidMessage ID from the query.
// Returns a *NotFoundError when no ValidMessage ID was found.
func (vmq *ValidMessageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vmq.Limit(1).IDs(setContextOp(ctx, vmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{validmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vmq *ValidMessageQuery) FirstIDX(ctx context.Context) int {
	id, err := vmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ValidMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ValidMessage entity is found.
// Returns a *NotFoundError when no ValidMessage entities are found.
func (vmq *ValidMessageQuery) Only(ctx context.Context) (*ValidMessage, error) {
	nodes, err := vmq.Limit(2).All(setContextOp(ctx, vmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{validmessage.Label}
	default:
		return nil, &NotSingularError{validmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vmq *ValidMessageQuery) OnlyX(ctx context.Context) *ValidMessage {
	node, err := vmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ValidMessage ID in the query.
// Returns a *NotSingularError when more than one ValidMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (vmq *ValidMessageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vmq.Limit(2).IDs(setContextOp(ctx, vmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{validmessage.Label}
	default:
		err = &NotSingularError{validmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vmq *ValidMessageQuery) OnlyIDX(ctx context.Context) int {
	id, err := vmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ValidMessages.
func (vmq *ValidMessageQuery) All(ctx context.Context) ([]*ValidMessage, error) {
	ctx = setContextOp(ctx, vmq.ctx, "All")
	if err := vmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ValidMessage, *ValidMessageQuery]()
	return withInterceptors[[]*ValidMessage](ctx, vmq, qr, vmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vmq *ValidMessageQuery) AllX(ctx context.Context) []*ValidMessage {
	nodes, err := vmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ValidMessage IDs.
func (vmq *ValidMessageQuery) IDs(ctx context.Context) (ids []int, err error) {
	if vmq.ctx.Unique == nil && vmq.path != nil {
		vmq.Unique(true)
	}
	ctx = setContextOp(ctx, vmq.ctx, "IDs")
	if err = vmq.Select(validmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vmq *ValidMessageQuery) IDsX(ctx context.Context) []int {
	ids, err := vmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vmq *ValidMessageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vmq.ctx, "Count")
	if err := vmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vmq, querierCount[*ValidMessageQuery](), vmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vmq *ValidMessageQuery) CountX(ctx context.Context) int {
	count, err := vmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vmq *ValidMessageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vmq.ctx, "Exist")
	switch _, err := vmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vmq *ValidMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := vmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ValidMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vmq *ValidMessageQuery) Clone() *ValidMessageQuery {
	if vmq == nil {
		return nil
	}
	return &ValidMessageQuery{
		config:     vmq.config,
		ctx:        vmq.ctx.Clone(),
		order:      append([]validmessage.OrderOption{}, vmq.order...),
		inters:     append([]Interceptor{}, vmq.inters...),
		predicates: append([]predicate.ValidMessage{}, vmq.predicates...),
		// clone intermediate query.
		sql:  vmq.sql.Clone(),
		path: vmq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ValidMessage.Query().
//		GroupBy(validmessage.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vmq *ValidMessageQuery) GroupBy(field string, fields ...string) *ValidMessageGroupBy {
	vmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ValidMessageGroupBy{build: vmq}
	grbuild.flds = &vmq.ctx.Fields
	grbuild.label = validmessage.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ValidMessage.Query().
//		Select(validmessage.FieldName).
//		Scan(ctx, &v)
func (vmq *ValidMessageQuery) Select(fields ...string) *ValidMessageSelect {
	vmq.ctx.Fields = append(vmq.ctx.Fields, fields...)
	sbuild := &ValidMessageSelect{ValidMessageQuery: vmq}
	sbuild.label = validmessage.Label
	sbuild.flds, sbuild.scan = &vmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ValidMessageSelect configured with the given aggregations.
func (vmq *ValidMessageQuery) Aggregate(fns ...AggregateFunc) *ValidMessageSelect {
	return vmq.Select().Aggregate(fns...)
}

func (vmq *ValidMessageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vmq); err != nil {
				return err
			}
		}
	}
	for _, f := range vmq.ctx.Fields {
		if !validmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vmq.path != nil {
		prev, err := vmq.path(ctx)
		if err != nil {
			return err
		}
		vmq.sql = prev
	}
	return nil
}

func (vmq *ValidMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ValidMessage, error) {
	var (
		nodes = []*ValidMessage{}
		_spec = vmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ValidMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ValidMessage{config: vmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vmq *ValidMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vmq.querySpec()
	_spec.Node.Columns = vmq.ctx.Fields
	if len(vmq.ctx.Fields) > 0 {
		_spec.Unique = vmq.ctx.Unique != nil && *vmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vmq.driver, _spec)
}

func (vmq *ValidMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(validmessage.Table, validmessage.Columns, sqlgraph.NewFieldSpec(validmessage.FieldID, field.TypeInt))
	_spec.From = vmq.sql
	if unique := vmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vmq.path != nil {
		_spec.Unique = true
	}
	if fields := vmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, validmessage.FieldID)
		for i := range fields {
			if fields[i] != validmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vmq *ValidMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vmq.driver.Dialect())
	t1 := builder.Table(validmessage.Table)
	columns := vmq.ctx.Fields
	if len(columns) == 0 {
		columns = validmessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vmq.sql != nil {
		selector = vmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vmq.ctx.Unique != nil && *vmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vmq.predicates {
		p(selector)
	}
	for _, p := range vmq.order {
		p(selector)
	}
	if offset := vmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ValidMessageGroupBy is the group-by builder for ValidMessage entities.
type ValidMessageGroupBy struct {
	selector
	build *ValidMessageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vmgb *ValidMessageGroupBy) Aggregate(fns ...AggregateFunc) *ValidMessageGroupBy {
	vmgb.fns = append(vmgb.fns, fns...)
	return vmgb
}

// Scan applies the selector query and scans the result into the given value.
func (vmgb *ValidMessageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vmgb.build.ctx, "GroupBy")
	if err := vmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ValidMessageQuery, *ValidMessageGroupBy](ctx, vmgb.build, vmgb, vmgb.build.inters, v)
}

func (vmgb *ValidMessageGroupBy) sqlScan(ctx context.Context, root *ValidMessageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vmgb.fns))
	for _, fn := range vmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vmgb.flds)+len(vmgb.fns))
		for _, f := range *vmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ValidMessageSelect is the builder for selecting fields of ValidMessage entities.
type ValidMessageSelect struct {
	*ValidMessageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vms *ValidMessageSelect) Aggregate(fns ...AggregateFunc) *ValidMessageSelect {
	vms.fns = append(vms.fns, fns...)
	return vms
}

// Scan applies the selector query and scans the result into the given value.
func (vms *ValidMessageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vms.ctx, "Select")
	if err := vms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ValidMessageQuery, *ValidMessageSelect](ctx, vms.ValidMessageQuery, vms, vms.inters, v)
}

func (vms *ValidMessageSelect) sqlScan(ctx context.Context, root *ValidMessageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vms.fns))
	for _, fn := range vms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
