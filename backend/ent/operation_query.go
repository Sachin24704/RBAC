// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/model"
	"backend/ent/operation"
	"backend/ent/predicate"
	"backend/ent/roleoperation"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OperationQuery is the builder for querying Operation entities.
type OperationQuery struct {
	config
	ctx                *QueryContext
	order              []operation.OrderOption
	inters             []Interceptor
	predicates         []predicate.Operation
	withModel          *ModelQuery
	withRoleOperations *RoleOperationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OperationQuery builder.
func (oq *OperationQuery) Where(ps ...predicate.Operation) *OperationQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit the number of records to be returned by this query.
func (oq *OperationQuery) Limit(limit int) *OperationQuery {
	oq.ctx.Limit = &limit
	return oq
}

// Offset to start from.
func (oq *OperationQuery) Offset(offset int) *OperationQuery {
	oq.ctx.Offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OperationQuery) Unique(unique bool) *OperationQuery {
	oq.ctx.Unique = &unique
	return oq
}

// Order specifies how the records should be ordered.
func (oq *OperationQuery) Order(o ...operation.OrderOption) *OperationQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryModel chains the current query on the "model" edge.
func (oq *OperationQuery) QueryModel() *ModelQuery {
	query := (&ModelClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(operation.Table, operation.FieldID, selector),
			sqlgraph.To(model.Table, model.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, operation.ModelTable, operation.ModelColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoleOperations chains the current query on the "role_operations" edge.
func (oq *OperationQuery) QueryRoleOperations() *RoleOperationQuery {
	query := (&RoleOperationClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(operation.Table, operation.FieldID, selector),
			sqlgraph.To(roleoperation.Table, roleoperation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, operation.RoleOperationsTable, operation.RoleOperationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Operation entity from the query.
// Returns a *NotFoundError when no Operation was found.
func (oq *OperationQuery) First(ctx context.Context) (*Operation, error) {
	nodes, err := oq.Limit(1).All(setContextOp(ctx, oq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{operation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OperationQuery) FirstX(ctx context.Context) *Operation {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Operation ID from the query.
// Returns a *NotFoundError when no Operation ID was found.
func (oq *OperationQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(setContextOp(ctx, oq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{operation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OperationQuery) FirstIDX(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Operation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Operation entity is found.
// Returns a *NotFoundError when no Operation entities are found.
func (oq *OperationQuery) Only(ctx context.Context) (*Operation, error) {
	nodes, err := oq.Limit(2).All(setContextOp(ctx, oq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{operation.Label}
	default:
		return nil, &NotSingularError{operation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OperationQuery) OnlyX(ctx context.Context) *Operation {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Operation ID in the query.
// Returns a *NotSingularError when more than one Operation ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *OperationQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(setContextOp(ctx, oq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{operation.Label}
	default:
		err = &NotSingularError{operation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OperationQuery) OnlyIDX(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Operations.
func (oq *OperationQuery) All(ctx context.Context) ([]*Operation, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryAll)
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Operation, *OperationQuery]()
	return withInterceptors[[]*Operation](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *OperationQuery) AllX(ctx context.Context) []*Operation {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Operation IDs.
func (oq *OperationQuery) IDs(ctx context.Context) (ids []int, err error) {
	if oq.ctx.Unique == nil && oq.path != nil {
		oq.Unique(true)
	}
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryIDs)
	if err = oq.Select(operation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OperationQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OperationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryCount)
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*OperationQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OperationQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OperationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryExist)
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OperationQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OperationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OperationQuery) Clone() *OperationQuery {
	if oq == nil {
		return nil
	}
	return &OperationQuery{
		config:             oq.config,
		ctx:                oq.ctx.Clone(),
		order:              append([]operation.OrderOption{}, oq.order...),
		inters:             append([]Interceptor{}, oq.inters...),
		predicates:         append([]predicate.Operation{}, oq.predicates...),
		withModel:          oq.withModel.Clone(),
		withRoleOperations: oq.withRoleOperations.Clone(),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// WithModel tells the query-builder to eager-load the nodes that are connected to
// the "model" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OperationQuery) WithModel(opts ...func(*ModelQuery)) *OperationQuery {
	query := (&ModelClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withModel = query
	return oq
}

// WithRoleOperations tells the query-builder to eager-load the nodes that are connected to
// the "role_operations" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OperationQuery) WithRoleOperations(opts ...func(*RoleOperationQuery)) *OperationQuery {
	query := (&RoleOperationClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withRoleOperations = query
	return oq
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
//	client.Operation.Query().
//		GroupBy(operation.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oq *OperationQuery) GroupBy(field string, fields ...string) *OperationGroupBy {
	oq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OperationGroupBy{build: oq}
	grbuild.flds = &oq.ctx.Fields
	grbuild.label = operation.Label
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
//	client.Operation.Query().
//		Select(operation.FieldName).
//		Scan(ctx, &v)
func (oq *OperationQuery) Select(fields ...string) *OperationSelect {
	oq.ctx.Fields = append(oq.ctx.Fields, fields...)
	sbuild := &OperationSelect{OperationQuery: oq}
	sbuild.label = operation.Label
	sbuild.flds, sbuild.scan = &oq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OperationSelect configured with the given aggregations.
func (oq *OperationQuery) Aggregate(fns ...AggregateFunc) *OperationSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *OperationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.ctx.Fields {
		if !operation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OperationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Operation, error) {
	var (
		nodes       = []*Operation{}
		_spec       = oq.querySpec()
		loadedTypes = [2]bool{
			oq.withModel != nil,
			oq.withRoleOperations != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Operation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Operation{config: oq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oq.withModel; query != nil {
		if err := oq.loadModel(ctx, query, nodes, nil,
			func(n *Operation, e *Model) { n.Edges.Model = e }); err != nil {
			return nil, err
		}
	}
	if query := oq.withRoleOperations; query != nil {
		if err := oq.loadRoleOperations(ctx, query, nodes,
			func(n *Operation) { n.Edges.RoleOperations = []*RoleOperation{} },
			func(n *Operation, e *RoleOperation) { n.Edges.RoleOperations = append(n.Edges.RoleOperations, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *OperationQuery) loadModel(ctx context.Context, query *ModelQuery, nodes []*Operation, init func(*Operation), assign func(*Operation, *Model)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Operation)
	for i := range nodes {
		fk := nodes[i].ModelID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(model.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "model_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (oq *OperationQuery) loadRoleOperations(ctx context.Context, query *RoleOperationQuery, nodes []*Operation, init func(*Operation), assign func(*Operation, *RoleOperation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Operation)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(roleoperation.FieldOperationID)
	}
	query.Where(predicate.RoleOperation(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(operation.RoleOperationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OperationID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "operation_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (oq *OperationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	_spec.Node.Columns = oq.ctx.Fields
	if len(oq.ctx.Fields) > 0 {
		_spec.Unique = oq.ctx.Unique != nil && *oq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OperationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(operation.Table, operation.Columns, sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt))
	_spec.From = oq.sql
	if unique := oq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oq.path != nil {
		_spec.Unique = true
	}
	if fields := oq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, operation.FieldID)
		for i := range fields {
			if fields[i] != operation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if oq.withModel != nil {
			_spec.Node.AddColumnOnce(operation.FieldModelID)
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OperationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(operation.Table)
	columns := oq.ctx.Fields
	if len(columns) == 0 {
		columns = operation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.ctx.Unique != nil && *oq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OperationGroupBy is the group-by builder for Operation entities.
type OperationGroupBy struct {
	selector
	build *OperationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OperationGroupBy) Aggregate(fns ...AggregateFunc) *OperationGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *OperationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ogb.build.ctx, ent.OpQueryGroupBy)
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OperationQuery, *OperationGroupBy](ctx, ogb.build, ogb, ogb.build.inters, v)
}

func (ogb *OperationGroupBy) sqlScan(ctx context.Context, root *OperationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OperationSelect is the builder for selecting fields of Operation entities.
type OperationSelect struct {
	*OperationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *OperationSelect) Aggregate(fns ...AggregateFunc) *OperationSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *OperationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, os.ctx, ent.OpQuerySelect)
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OperationQuery, *OperationSelect](ctx, os.OperationQuery, os, os.inters, v)
}

func (os *OperationSelect) sqlScan(ctx context.Context, root *OperationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
