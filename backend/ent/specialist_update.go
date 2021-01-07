// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team09/app/ent/predicate"
	"github.com/team09/app/ent/special_doctor"
	"github.com/team09/app/ent/specialist"
)

// SpecialistUpdate is the builder for updating Specialist entities.
type SpecialistUpdate struct {
	config
	hooks      []Hook
	mutation   *SpecialistMutation
	predicates []predicate.Specialist
}

// Where adds a new predicate for the builder.
func (su *SpecialistUpdate) Where(ps ...predicate.Specialist) *SpecialistUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetSpecialist sets the specialist field.
func (su *SpecialistUpdate) SetSpecialist(s string) *SpecialistUpdate {
	su.mutation.SetSpecialist(s)
	return su
}

// AddSpecialDoctorIDs adds the special_doctors edge to Special_Doctor by ids.
func (su *SpecialistUpdate) AddSpecialDoctorIDs(ids ...int) *SpecialistUpdate {
	su.mutation.AddSpecialDoctorIDs(ids...)
	return su
}

// AddSpecialDoctors adds the special_doctors edges to Special_Doctor.
func (su *SpecialistUpdate) AddSpecialDoctors(s ...*Special_Doctor) *SpecialistUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddSpecialDoctorIDs(ids...)
}

// Mutation returns the SpecialistMutation object of the builder.
func (su *SpecialistUpdate) Mutation() *SpecialistMutation {
	return su.mutation
}

// RemoveSpecialDoctorIDs removes the special_doctors edge to Special_Doctor by ids.
func (su *SpecialistUpdate) RemoveSpecialDoctorIDs(ids ...int) *SpecialistUpdate {
	su.mutation.RemoveSpecialDoctorIDs(ids...)
	return su
}

// RemoveSpecialDoctors removes special_doctors edges to Special_Doctor.
func (su *SpecialistUpdate) RemoveSpecialDoctors(s ...*Special_Doctor) *SpecialistUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveSpecialDoctorIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SpecialistUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := su.mutation.Specialist(); ok {
		if err := specialist.SpecialistValidator(v); err != nil {
			return 0, &ValidationError{Name: "specialist", err: fmt.Errorf("ent: validator failed for field \"specialist\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecialistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SpecialistUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SpecialistUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SpecialistUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SpecialistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   specialist.Table,
			Columns: specialist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: specialist.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Specialist(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: specialist.FieldSpecialist,
		})
	}
	if nodes := su.mutation.RemovedSpecialDoctorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   specialist.SpecialDoctorsTable,
			Columns: []string{specialist.SpecialDoctorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: special_doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SpecialDoctorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   specialist.SpecialDoctorsTable,
			Columns: []string{specialist.SpecialDoctorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: special_doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{specialist.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SpecialistUpdateOne is the builder for updating a single Specialist entity.
type SpecialistUpdateOne struct {
	config
	hooks    []Hook
	mutation *SpecialistMutation
}

// SetSpecialist sets the specialist field.
func (suo *SpecialistUpdateOne) SetSpecialist(s string) *SpecialistUpdateOne {
	suo.mutation.SetSpecialist(s)
	return suo
}

// AddSpecialDoctorIDs adds the special_doctors edge to Special_Doctor by ids.
func (suo *SpecialistUpdateOne) AddSpecialDoctorIDs(ids ...int) *SpecialistUpdateOne {
	suo.mutation.AddSpecialDoctorIDs(ids...)
	return suo
}

// AddSpecialDoctors adds the special_doctors edges to Special_Doctor.
func (suo *SpecialistUpdateOne) AddSpecialDoctors(s ...*Special_Doctor) *SpecialistUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddSpecialDoctorIDs(ids...)
}

// Mutation returns the SpecialistMutation object of the builder.
func (suo *SpecialistUpdateOne) Mutation() *SpecialistMutation {
	return suo.mutation
}

// RemoveSpecialDoctorIDs removes the special_doctors edge to Special_Doctor by ids.
func (suo *SpecialistUpdateOne) RemoveSpecialDoctorIDs(ids ...int) *SpecialistUpdateOne {
	suo.mutation.RemoveSpecialDoctorIDs(ids...)
	return suo
}

// RemoveSpecialDoctors removes special_doctors edges to Special_Doctor.
func (suo *SpecialistUpdateOne) RemoveSpecialDoctors(s ...*Special_Doctor) *SpecialistUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveSpecialDoctorIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *SpecialistUpdateOne) Save(ctx context.Context) (*Specialist, error) {
	if v, ok := suo.mutation.Specialist(); ok {
		if err := specialist.SpecialistValidator(v); err != nil {
			return nil, &ValidationError{Name: "specialist", err: fmt.Errorf("ent: validator failed for field \"specialist\": %w", err)}
		}
	}

	var (
		err  error
		node *Specialist
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecialistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SpecialistUpdateOne) SaveX(ctx context.Context) *Specialist {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SpecialistUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SpecialistUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SpecialistUpdateOne) sqlSave(ctx context.Context) (s *Specialist, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   specialist.Table,
			Columns: specialist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: specialist.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Specialist.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Specialist(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: specialist.FieldSpecialist,
		})
	}
	if nodes := suo.mutation.RemovedSpecialDoctorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   specialist.SpecialDoctorsTable,
			Columns: []string{specialist.SpecialDoctorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: special_doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SpecialDoctorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   specialist.SpecialDoctorsTable,
			Columns: []string{specialist.SpecialDoctorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: special_doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Specialist{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{specialist.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
