// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/disease"
	"github.com/team09/app/ent/doctor"
	"github.com/team09/app/ent/gender"
	"github.com/team09/app/ent/office"
	"github.com/team09/app/ent/position"
	"github.com/team09/app/ent/predicate"
	"github.com/team09/app/ent/title"
)

// DoctorQuery is the builder for querying Doctor entities.
type DoctorQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Doctor
	// eager-loading edges.
	withTitle       *TitleQuery
	withGender      *GenderQuery
	withPosition    *PositionQuery
	withDisease     *DiseaseQuery
	withOffices     *OfficeQuery
	withDepartments *DepartmentQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (dq *DoctorQuery) Where(ps ...predicate.Doctor) *DoctorQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DoctorQuery) Limit(limit int) *DoctorQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DoctorQuery) Offset(offset int) *DoctorQuery {
	dq.offset = &offset
	return dq
}

// Order adds an order step to the query.
func (dq *DoctorQuery) Order(o ...OrderFunc) *DoctorQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryTitle chains the current query on the title edge.
func (dq *DoctorQuery) QueryTitle() *TitleQuery {
	query := &TitleQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctor.Table, doctor.FieldID, dq.sqlQuery()),
			sqlgraph.To(title.Table, title.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, doctor.TitleTable, doctor.TitleColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGender chains the current query on the gender edge.
func (dq *DoctorQuery) QueryGender() *GenderQuery {
	query := &GenderQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctor.Table, doctor.FieldID, dq.sqlQuery()),
			sqlgraph.To(gender.Table, gender.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, doctor.GenderTable, doctor.GenderColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPosition chains the current query on the position edge.
func (dq *DoctorQuery) QueryPosition() *PositionQuery {
	query := &PositionQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctor.Table, doctor.FieldID, dq.sqlQuery()),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, doctor.PositionTable, doctor.PositionColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDisease chains the current query on the disease edge.
func (dq *DoctorQuery) QueryDisease() *DiseaseQuery {
	query := &DiseaseQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctor.Table, doctor.FieldID, dq.sqlQuery()),
			sqlgraph.To(disease.Table, disease.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, doctor.DiseaseTable, doctor.DiseaseColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOffices chains the current query on the offices edge.
func (dq *DoctorQuery) QueryOffices() *OfficeQuery {
	query := &OfficeQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctor.Table, doctor.FieldID, dq.sqlQuery()),
			sqlgraph.To(office.Table, office.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, doctor.OfficesTable, doctor.OfficesColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDepartments chains the current query on the departments edge.
func (dq *DoctorQuery) QueryDepartments() *DepartmentQuery {
	query := &DepartmentQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctor.Table, doctor.FieldID, dq.sqlQuery()),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, doctor.DepartmentsTable, doctor.DepartmentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Doctor entity in the query. Returns *NotFoundError when no doctor was found.
func (dq *DoctorQuery) First(ctx context.Context) (*Doctor, error) {
	ds, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ds) == 0 {
		return nil, &NotFoundError{doctor.Label}
	}
	return ds[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DoctorQuery) FirstX(ctx context.Context) *Doctor {
	d, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return d
}

// FirstID returns the first Doctor id in the query. Returns *NotFoundError when no id was found.
func (dq *DoctorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{doctor.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (dq *DoctorQuery) FirstXID(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Doctor entity in the query, returns an error if not exactly one entity was returned.
func (dq *DoctorQuery) Only(ctx context.Context) (*Doctor, error) {
	ds, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ds) {
	case 1:
		return ds[0], nil
	case 0:
		return nil, &NotFoundError{doctor.Label}
	default:
		return nil, &NotSingularError{doctor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DoctorQuery) OnlyX(ctx context.Context) *Doctor {
	d, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// OnlyID returns the only Doctor id in the query, returns an error if not exactly one id was returned.
func (dq *DoctorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = &NotSingularError{doctor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DoctorQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Doctors.
func (dq *DoctorQuery) All(ctx context.Context) ([]*Doctor, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DoctorQuery) AllX(ctx context.Context) []*Doctor {
	ds, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ds
}

// IDs executes the query and returns a list of Doctor ids.
func (dq *DoctorQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dq.Select(doctor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DoctorQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DoctorQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DoctorQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DoctorQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DoctorQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DoctorQuery) Clone() *DoctorQuery {
	return &DoctorQuery{
		config:     dq.config,
		limit:      dq.limit,
		offset:     dq.offset,
		order:      append([]OrderFunc{}, dq.order...),
		unique:     append([]string{}, dq.unique...),
		predicates: append([]predicate.Doctor{}, dq.predicates...),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

//  WithTitle tells the query-builder to eager-loads the nodes that are connected to
// the "title" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorQuery) WithTitle(opts ...func(*TitleQuery)) *DoctorQuery {
	query := &TitleQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withTitle = query
	return dq
}

//  WithGender tells the query-builder to eager-loads the nodes that are connected to
// the "gender" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorQuery) WithGender(opts ...func(*GenderQuery)) *DoctorQuery {
	query := &GenderQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withGender = query
	return dq
}

//  WithPosition tells the query-builder to eager-loads the nodes that are connected to
// the "position" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorQuery) WithPosition(opts ...func(*PositionQuery)) *DoctorQuery {
	query := &PositionQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withPosition = query
	return dq
}

//  WithDisease tells the query-builder to eager-loads the nodes that are connected to
// the "disease" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorQuery) WithDisease(opts ...func(*DiseaseQuery)) *DoctorQuery {
	query := &DiseaseQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withDisease = query
	return dq
}

//  WithOffices tells the query-builder to eager-loads the nodes that are connected to
// the "offices" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorQuery) WithOffices(opts ...func(*OfficeQuery)) *DoctorQuery {
	query := &OfficeQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withOffices = query
	return dq
}

//  WithDepartments tells the query-builder to eager-loads the nodes that are connected to
// the "departments" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorQuery) WithDepartments(opts ...func(*DepartmentQuery)) *DoctorQuery {
	query := &DepartmentQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withDepartments = query
	return dq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Doctor.Query().
//		GroupBy(doctor.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dq *DoctorQuery) GroupBy(field string, fields ...string) *DoctorGroupBy {
	group := &DoctorGroupBy{config: dq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Doctor.Query().
//		Select(doctor.FieldName).
//		Scan(ctx, &v)
//
func (dq *DoctorQuery) Select(field string, fields ...string) *DoctorSelect {
	selector := &DoctorSelect{config: dq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(), nil
	}
	return selector
}

func (dq *DoctorQuery) prepareQuery(ctx context.Context) error {
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DoctorQuery) sqlAll(ctx context.Context) ([]*Doctor, error) {
	var (
		nodes       = []*Doctor{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [6]bool{
			dq.withTitle != nil,
			dq.withGender != nil,
			dq.withPosition != nil,
			dq.withDisease != nil,
			dq.withOffices != nil,
			dq.withDepartments != nil,
		}
	)
	if dq.withTitle != nil || dq.withGender != nil || dq.withPosition != nil || dq.withDisease != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, doctor.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Doctor{config: dq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dq.withTitle; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Doctor)
		for i := range nodes {
			if fk := nodes[i].title_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(title.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "title_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Title = n
			}
		}
	}

	if query := dq.withGender; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Doctor)
		for i := range nodes {
			if fk := nodes[i].gender_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(gender.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "gender_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Gender = n
			}
		}
	}

	if query := dq.withPosition; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Doctor)
		for i := range nodes {
			if fk := nodes[i].position_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(position.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "position_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Position = n
			}
		}
	}

	if query := dq.withDisease; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Doctor)
		for i := range nodes {
			if fk := nodes[i].disease_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(disease.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "disease_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Disease = n
			}
		}
	}

	if query := dq.withOffices; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Doctor)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Office(func(s *sql.Selector) {
			s.Where(sql.InValues(doctor.OfficesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.doctor_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "doctor_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "doctor_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Offices = append(node.Edges.Offices, n)
		}
	}

	if query := dq.withDepartments; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Doctor)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Department(func(s *sql.Selector) {
			s.Where(sql.InValues(doctor.DepartmentsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.doctor_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "doctor_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "doctor_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Departments = append(node.Edges.Departments, n)
		}
	}

	return nodes, nil
}

func (dq *DoctorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DoctorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (dq *DoctorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctor.Table,
			Columns: doctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DoctorQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(doctor.Table)
	selector := builder.Select(t1.Columns(doctor.Columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(doctor.Columns...)...)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DoctorGroupBy is the builder for group-by Doctor entities.
type DoctorGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DoctorGroupBy) Aggregate(fns ...AggregateFunc) *DoctorGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scan the result into the given value.
func (dgb *DoctorGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dgb *DoctorGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dgb *DoctorGroupBy) StringsX(ctx context.Context) []string {
	v, err := dgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dgb *DoctorGroupBy) StringX(ctx context.Context) string {
	v, err := dgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dgb *DoctorGroupBy) IntsX(ctx context.Context) []int {
	v, err := dgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dgb *DoctorGroupBy) IntX(ctx context.Context) int {
	v, err := dgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dgb *DoctorGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dgb *DoctorGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dgb *DoctorGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dgb *DoctorGroupBy) BoolX(ctx context.Context) bool {
	v, err := dgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dgb *DoctorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dgb.sqlQuery().Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DoctorGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql
	columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
	columns = append(columns, dgb.fields...)
	for _, fn := range dgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(dgb.fields...)
}

// DoctorSelect is the builder for select fields of Doctor entities.
type DoctorSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ds *DoctorSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ds.path(ctx)
	if err != nil {
		return err
	}
	ds.sql = query
	return ds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ds *DoctorSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ds *DoctorSelect) StringsX(ctx context.Context) []string {
	v, err := ds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ds *DoctorSelect) StringX(ctx context.Context) string {
	v, err := ds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ds *DoctorSelect) IntsX(ctx context.Context) []int {
	v, err := ds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ds *DoctorSelect) IntX(ctx context.Context) int {
	v, err := ds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ds *DoctorSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ds *DoctorSelect) Float64X(ctx context.Context) float64 {
	v, err := ds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ds *DoctorSelect) BoolsX(ctx context.Context) []bool {
	v, err := ds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = fmt.Errorf("ent: DoctorSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ds *DoctorSelect) BoolX(ctx context.Context) bool {
	v, err := ds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ds *DoctorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sqlQuery().Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ds *DoctorSelect) sqlQuery() sql.Querier {
	selector := ds.sql
	selector.Select(selector.Columns(ds.fields...)...)
	return selector
}
