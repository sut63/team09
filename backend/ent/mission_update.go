// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/detail"
	"github.com/team09/app/ent/mission"
	"github.com/team09/app/ent/predicate"
)

// MissionUpdate is the builder for updating Mission entities.
type MissionUpdate struct {
	config
	hooks      []Hook
	mutation   *MissionMutation
	predicates []predicate.Mission
}

// Where adds a new predicate for the builder.
func (mu *MissionUpdate) Where(ps ...predicate.Mission) *MissionUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetMission sets the mission field.
func (mu *MissionUpdate) SetMission(s string) *MissionUpdate {
	mu.mutation.SetMission(s)
	return mu
}

// AddDepartmentIDs adds the departments edge to Department by ids.
func (mu *MissionUpdate) AddDepartmentIDs(ids ...int) *MissionUpdate {
	mu.mutation.AddDepartmentIDs(ids...)
	return mu
}

// AddDepartments adds the departments edges to Department.
func (mu *MissionUpdate) AddDepartments(d ...*Department) *MissionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mu.AddDepartmentIDs(ids...)
}

// AddDetailIDs adds the details edge to Detail by ids.
func (mu *MissionUpdate) AddDetailIDs(ids ...int) *MissionUpdate {
	mu.mutation.AddDetailIDs(ids...)
	return mu
}

// AddDetails adds the details edges to Detail.
func (mu *MissionUpdate) AddDetails(d ...*Detail) *MissionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mu.AddDetailIDs(ids...)
}

// Mutation returns the MissionMutation object of the builder.
func (mu *MissionUpdate) Mutation() *MissionMutation {
	return mu.mutation
}

// RemoveDepartmentIDs removes the departments edge to Department by ids.
func (mu *MissionUpdate) RemoveDepartmentIDs(ids ...int) *MissionUpdate {
	mu.mutation.RemoveDepartmentIDs(ids...)
	return mu
}

// RemoveDepartments removes departments edges to Department.
func (mu *MissionUpdate) RemoveDepartments(d ...*Department) *MissionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mu.RemoveDepartmentIDs(ids...)
}

// RemoveDetailIDs removes the details edge to Detail by ids.
func (mu *MissionUpdate) RemoveDetailIDs(ids ...int) *MissionUpdate {
	mu.mutation.RemoveDetailIDs(ids...)
	return mu
}

// RemoveDetails removes details edges to Detail.
func (mu *MissionUpdate) RemoveDetails(d ...*Detail) *MissionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mu.RemoveDetailIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *MissionUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := mu.mutation.Mission(); ok {
		if err := mission.MissionValidator(v); err != nil {
			return 0, &ValidationError{Name: "mission", err: fmt.Errorf("ent: validator failed for field \"mission\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MissionUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MissionUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MissionUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mission.Table,
			Columns: mission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mission.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Mission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mission.FieldMission,
		})
	}
	if nodes := mu.mutation.RemovedDepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.DepartmentsTable,
			Columns: []string{mission.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.DepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.DepartmentsTable,
			Columns: []string{mission.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := mu.mutation.RemovedDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.DetailsTable,
			Columns: []string{mission.DetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: detail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.DetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.DetailsTable,
			Columns: []string{mission.DetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: detail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mission.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MissionUpdateOne is the builder for updating a single Mission entity.
type MissionUpdateOne struct {
	config
	hooks    []Hook
	mutation *MissionMutation
}

// SetMission sets the mission field.
func (muo *MissionUpdateOne) SetMission(s string) *MissionUpdateOne {
	muo.mutation.SetMission(s)
	return muo
}

// AddDepartmentIDs adds the departments edge to Department by ids.
func (muo *MissionUpdateOne) AddDepartmentIDs(ids ...int) *MissionUpdateOne {
	muo.mutation.AddDepartmentIDs(ids...)
	return muo
}

// AddDepartments adds the departments edges to Department.
func (muo *MissionUpdateOne) AddDepartments(d ...*Department) *MissionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return muo.AddDepartmentIDs(ids...)
}

// AddDetailIDs adds the details edge to Detail by ids.
func (muo *MissionUpdateOne) AddDetailIDs(ids ...int) *MissionUpdateOne {
	muo.mutation.AddDetailIDs(ids...)
	return muo
}

// AddDetails adds the details edges to Detail.
func (muo *MissionUpdateOne) AddDetails(d ...*Detail) *MissionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return muo.AddDetailIDs(ids...)
}

// Mutation returns the MissionMutation object of the builder.
func (muo *MissionUpdateOne) Mutation() *MissionMutation {
	return muo.mutation
}

// RemoveDepartmentIDs removes the departments edge to Department by ids.
func (muo *MissionUpdateOne) RemoveDepartmentIDs(ids ...int) *MissionUpdateOne {
	muo.mutation.RemoveDepartmentIDs(ids...)
	return muo
}

// RemoveDepartments removes departments edges to Department.
func (muo *MissionUpdateOne) RemoveDepartments(d ...*Department) *MissionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return muo.RemoveDepartmentIDs(ids...)
}

// RemoveDetailIDs removes the details edge to Detail by ids.
func (muo *MissionUpdateOne) RemoveDetailIDs(ids ...int) *MissionUpdateOne {
	muo.mutation.RemoveDetailIDs(ids...)
	return muo
}

// RemoveDetails removes details edges to Detail.
func (muo *MissionUpdateOne) RemoveDetails(d ...*Detail) *MissionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return muo.RemoveDetailIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *MissionUpdateOne) Save(ctx context.Context) (*Mission, error) {
	if v, ok := muo.mutation.Mission(); ok {
		if err := mission.MissionValidator(v); err != nil {
			return nil, &ValidationError{Name: "mission", err: fmt.Errorf("ent: validator failed for field \"mission\": %w", err)}
		}
	}

	var (
		err  error
		node *Mission
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MissionUpdateOne) SaveX(ctx context.Context) *Mission {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *MissionUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MissionUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MissionUpdateOne) sqlSave(ctx context.Context) (m *Mission, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mission.Table,
			Columns: mission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mission.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Mission.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.Mission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mission.FieldMission,
		})
	}
	if nodes := muo.mutation.RemovedDepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.DepartmentsTable,
			Columns: []string{mission.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.DepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.DepartmentsTable,
			Columns: []string{mission.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := muo.mutation.RemovedDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.DetailsTable,
			Columns: []string{mission.DetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: detail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.DetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.DetailsTable,
			Columns: []string{mission.DetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: detail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	m = &Mission{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mission.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}
