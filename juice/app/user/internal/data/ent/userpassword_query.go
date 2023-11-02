// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"juice/app/user/internal/data/ent/predicate"
	"juice/app/user/internal/data/ent/userpassword"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPasswordQuery is the builder for querying UserPassword entities.
type UserPasswordQuery struct {
	config
	ctx        *QueryContext
	order      []userpassword.OrderOption
	inters     []Interceptor
	predicates []predicate.UserPassword
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserPasswordQuery builder.
func (upq *UserPasswordQuery) Where(ps ...predicate.UserPassword) *UserPasswordQuery {
	upq.predicates = append(upq.predicates, ps...)
	return upq
}

// Limit the number of records to be returned by this query.
func (upq *UserPasswordQuery) Limit(limit int) *UserPasswordQuery {
	upq.ctx.Limit = &limit
	return upq
}

// Offset to start from.
func (upq *UserPasswordQuery) Offset(offset int) *UserPasswordQuery {
	upq.ctx.Offset = &offset
	return upq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (upq *UserPasswordQuery) Unique(unique bool) *UserPasswordQuery {
	upq.ctx.Unique = &unique
	return upq
}

// Order specifies how the records should be ordered.
func (upq *UserPasswordQuery) Order(o ...userpassword.OrderOption) *UserPasswordQuery {
	upq.order = append(upq.order, o...)
	return upq
}

// First returns the first UserPassword entity from the query.
// Returns a *NotFoundError when no UserPassword was found.
func (upq *UserPasswordQuery) First(ctx context.Context) (*UserPassword, error) {
	nodes, err := upq.Limit(1).All(setContextOp(ctx, upq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userpassword.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (upq *UserPasswordQuery) FirstX(ctx context.Context) *UserPassword {
	node, err := upq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserPassword ID from the query.
// Returns a *NotFoundError when no UserPassword ID was found.
func (upq *UserPasswordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = upq.Limit(1).IDs(setContextOp(ctx, upq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userpassword.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (upq *UserPasswordQuery) FirstIDX(ctx context.Context) int {
	id, err := upq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserPassword entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserPassword entity is found.
// Returns a *NotFoundError when no UserPassword entities are found.
func (upq *UserPasswordQuery) Only(ctx context.Context) (*UserPassword, error) {
	nodes, err := upq.Limit(2).All(setContextOp(ctx, upq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userpassword.Label}
	default:
		return nil, &NotSingularError{userpassword.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (upq *UserPasswordQuery) OnlyX(ctx context.Context) *UserPassword {
	node, err := upq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserPassword ID in the query.
// Returns a *NotSingularError when more than one UserPassword ID is found.
// Returns a *NotFoundError when no entities are found.
func (upq *UserPasswordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = upq.Limit(2).IDs(setContextOp(ctx, upq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userpassword.Label}
	default:
		err = &NotSingularError{userpassword.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (upq *UserPasswordQuery) OnlyIDX(ctx context.Context) int {
	id, err := upq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserPasswords.
func (upq *UserPasswordQuery) All(ctx context.Context) ([]*UserPassword, error) {
	ctx = setContextOp(ctx, upq.ctx, "All")
	if err := upq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserPassword, *UserPasswordQuery]()
	return withInterceptors[[]*UserPassword](ctx, upq, qr, upq.inters)
}

// AllX is like All, but panics if an error occurs.
func (upq *UserPasswordQuery) AllX(ctx context.Context) []*UserPassword {
	nodes, err := upq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserPassword IDs.
func (upq *UserPasswordQuery) IDs(ctx context.Context) (ids []int, err error) {
	if upq.ctx.Unique == nil && upq.path != nil {
		upq.Unique(true)
	}
	ctx = setContextOp(ctx, upq.ctx, "IDs")
	if err = upq.Select(userpassword.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (upq *UserPasswordQuery) IDsX(ctx context.Context) []int {
	ids, err := upq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (upq *UserPasswordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, upq.ctx, "Count")
	if err := upq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, upq, querierCount[*UserPasswordQuery](), upq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (upq *UserPasswordQuery) CountX(ctx context.Context) int {
	count, err := upq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (upq *UserPasswordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, upq.ctx, "Exist")
	switch _, err := upq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (upq *UserPasswordQuery) ExistX(ctx context.Context) bool {
	exist, err := upq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserPasswordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (upq *UserPasswordQuery) Clone() *UserPasswordQuery {
	if upq == nil {
		return nil
	}
	return &UserPasswordQuery{
		config:     upq.config,
		ctx:        upq.ctx.Clone(),
		order:      append([]userpassword.OrderOption{}, upq.order...),
		inters:     append([]Interceptor{}, upq.inters...),
		predicates: append([]predicate.UserPassword{}, upq.predicates...),
		// clone intermediate query.
		sql:  upq.sql.Clone(),
		path: upq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uint64 `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserPassword.Query().
//		GroupBy(userpassword.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (upq *UserPasswordQuery) GroupBy(field string, fields ...string) *UserPasswordGroupBy {
	upq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserPasswordGroupBy{build: upq}
	grbuild.flds = &upq.ctx.Fields
	grbuild.label = userpassword.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uint64 `json:"user_id,omitempty"`
//	}
//
//	client.UserPassword.Query().
//		Select(userpassword.FieldUserID).
//		Scan(ctx, &v)
func (upq *UserPasswordQuery) Select(fields ...string) *UserPasswordSelect {
	upq.ctx.Fields = append(upq.ctx.Fields, fields...)
	sbuild := &UserPasswordSelect{UserPasswordQuery: upq}
	sbuild.label = userpassword.Label
	sbuild.flds, sbuild.scan = &upq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserPasswordSelect configured with the given aggregations.
func (upq *UserPasswordQuery) Aggregate(fns ...AggregateFunc) *UserPasswordSelect {
	return upq.Select().Aggregate(fns...)
}

func (upq *UserPasswordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range upq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, upq); err != nil {
				return err
			}
		}
	}
	for _, f := range upq.ctx.Fields {
		if !userpassword.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if upq.path != nil {
		prev, err := upq.path(ctx)
		if err != nil {
			return err
		}
		upq.sql = prev
	}
	return nil
}

func (upq *UserPasswordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserPassword, error) {
	var (
		nodes = []*UserPassword{}
		_spec = upq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserPassword).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserPassword{config: upq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, upq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (upq *UserPasswordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := upq.querySpec()
	_spec.Node.Columns = upq.ctx.Fields
	if len(upq.ctx.Fields) > 0 {
		_spec.Unique = upq.ctx.Unique != nil && *upq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, upq.driver, _spec)
}

func (upq *UserPasswordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userpassword.Table, userpassword.Columns, sqlgraph.NewFieldSpec(userpassword.FieldID, field.TypeInt))
	_spec.From = upq.sql
	if unique := upq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if upq.path != nil {
		_spec.Unique = true
	}
	if fields := upq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userpassword.FieldID)
		for i := range fields {
			if fields[i] != userpassword.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := upq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := upq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := upq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := upq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (upq *UserPasswordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(upq.driver.Dialect())
	t1 := builder.Table(userpassword.Table)
	columns := upq.ctx.Fields
	if len(columns) == 0 {
		columns = userpassword.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if upq.sql != nil {
		selector = upq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if upq.ctx.Unique != nil && *upq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range upq.predicates {
		p(selector)
	}
	for _, p := range upq.order {
		p(selector)
	}
	if offset := upq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := upq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserPasswordGroupBy is the group-by builder for UserPassword entities.
type UserPasswordGroupBy struct {
	selector
	build *UserPasswordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (upgb *UserPasswordGroupBy) Aggregate(fns ...AggregateFunc) *UserPasswordGroupBy {
	upgb.fns = append(upgb.fns, fns...)
	return upgb
}

// Scan applies the selector query and scans the result into the given value.
func (upgb *UserPasswordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, upgb.build.ctx, "GroupBy")
	if err := upgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserPasswordQuery, *UserPasswordGroupBy](ctx, upgb.build, upgb, upgb.build.inters, v)
}

func (upgb *UserPasswordGroupBy) sqlScan(ctx context.Context, root *UserPasswordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(upgb.fns))
	for _, fn := range upgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*upgb.flds)+len(upgb.fns))
		for _, f := range *upgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*upgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := upgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserPasswordSelect is the builder for selecting fields of UserPassword entities.
type UserPasswordSelect struct {
	*UserPasswordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ups *UserPasswordSelect) Aggregate(fns ...AggregateFunc) *UserPasswordSelect {
	ups.fns = append(ups.fns, fns...)
	return ups
}

// Scan applies the selector query and scans the result into the given value.
func (ups *UserPasswordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ups.ctx, "Select")
	if err := ups.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserPasswordQuery, *UserPasswordSelect](ctx, ups.UserPasswordQuery, ups, ups.inters, v)
}

func (ups *UserPasswordSelect) sqlScan(ctx context.Context, root *UserPasswordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ups.fns))
	for _, fn := range ups.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ups.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ups.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}