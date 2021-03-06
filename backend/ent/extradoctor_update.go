// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team09/app/ent/extradoctor"
	"github.com/team09/app/ent/office"
	"github.com/team09/app/ent/predicate"
	"github.com/team09/app/ent/specialdoctor"
)

// ExtradoctorUpdate is the builder for updating Extradoctor entities.
type ExtradoctorUpdate struct {
	config
	hooks      []Hook
	mutation   *ExtradoctorMutation
	predicates []predicate.Extradoctor
}

// Where adds a new predicate for the builder.
func (eu *ExtradoctorUpdate) Where(ps ...predicate.Extradoctor) *ExtradoctorUpdate {
	eu.predicates = append(eu.predicates, ps...)
	return eu
}

// SetSpecialname sets the specialname field.
func (eu *ExtradoctorUpdate) SetSpecialname(s string) *ExtradoctorUpdate {
	eu.mutation.SetSpecialname(s)
	return eu
}

// AddSpecialdoctorIDs adds the specialdoctors edge to Specialdoctor by ids.
func (eu *ExtradoctorUpdate) AddSpecialdoctorIDs(ids ...int) *ExtradoctorUpdate {
	eu.mutation.AddSpecialdoctorIDs(ids...)
	return eu
}

// AddSpecialdoctors adds the specialdoctors edges to Specialdoctor.
func (eu *ExtradoctorUpdate) AddSpecialdoctors(s ...*Specialdoctor) *ExtradoctorUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return eu.AddSpecialdoctorIDs(ids...)
}

// AddOfficeIDs adds the offices edge to Office by ids.
func (eu *ExtradoctorUpdate) AddOfficeIDs(ids ...int) *ExtradoctorUpdate {
	eu.mutation.AddOfficeIDs(ids...)
	return eu
}

// AddOffices adds the offices edges to Office.
func (eu *ExtradoctorUpdate) AddOffices(o ...*Office) *ExtradoctorUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return eu.AddOfficeIDs(ids...)
}

// Mutation returns the ExtradoctorMutation object of the builder.
func (eu *ExtradoctorUpdate) Mutation() *ExtradoctorMutation {
	return eu.mutation
}

// RemoveSpecialdoctorIDs removes the specialdoctors edge to Specialdoctor by ids.
func (eu *ExtradoctorUpdate) RemoveSpecialdoctorIDs(ids ...int) *ExtradoctorUpdate {
	eu.mutation.RemoveSpecialdoctorIDs(ids...)
	return eu
}

// RemoveSpecialdoctors removes specialdoctors edges to Specialdoctor.
func (eu *ExtradoctorUpdate) RemoveSpecialdoctors(s ...*Specialdoctor) *ExtradoctorUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return eu.RemoveSpecialdoctorIDs(ids...)
}

// RemoveOfficeIDs removes the offices edge to Office by ids.
func (eu *ExtradoctorUpdate) RemoveOfficeIDs(ids ...int) *ExtradoctorUpdate {
	eu.mutation.RemoveOfficeIDs(ids...)
	return eu
}

// RemoveOffices removes offices edges to Office.
func (eu *ExtradoctorUpdate) RemoveOffices(o ...*Office) *ExtradoctorUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return eu.RemoveOfficeIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (eu *ExtradoctorUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := eu.mutation.Specialname(); ok {
		if err := extradoctor.SpecialnameValidator(v); err != nil {
			return 0, &ValidationError{Name: "specialname", err: fmt.Errorf("ent: validator failed for field \"specialname\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExtradoctorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *ExtradoctorUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *ExtradoctorUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *ExtradoctorUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *ExtradoctorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   extradoctor.Table,
			Columns: extradoctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: extradoctor.FieldID,
			},
		},
	}
	if ps := eu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Specialname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: extradoctor.FieldSpecialname,
		})
	}
	if nodes := eu.mutation.RemovedSpecialdoctorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   extradoctor.SpecialdoctorsTable,
			Columns: []string{extradoctor.SpecialdoctorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: specialdoctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.SpecialdoctorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   extradoctor.SpecialdoctorsTable,
			Columns: []string{extradoctor.SpecialdoctorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: specialdoctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := eu.mutation.RemovedOfficesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   extradoctor.OfficesTable,
			Columns: []string{extradoctor.OfficesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: office.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.OfficesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   extradoctor.OfficesTable,
			Columns: []string{extradoctor.OfficesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: office.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{extradoctor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ExtradoctorUpdateOne is the builder for updating a single Extradoctor entity.
type ExtradoctorUpdateOne struct {
	config
	hooks    []Hook
	mutation *ExtradoctorMutation
}

// SetSpecialname sets the specialname field.
func (euo *ExtradoctorUpdateOne) SetSpecialname(s string) *ExtradoctorUpdateOne {
	euo.mutation.SetSpecialname(s)
	return euo
}

// AddSpecialdoctorIDs adds the specialdoctors edge to Specialdoctor by ids.
func (euo *ExtradoctorUpdateOne) AddSpecialdoctorIDs(ids ...int) *ExtradoctorUpdateOne {
	euo.mutation.AddSpecialdoctorIDs(ids...)
	return euo
}

// AddSpecialdoctors adds the specialdoctors edges to Specialdoctor.
func (euo *ExtradoctorUpdateOne) AddSpecialdoctors(s ...*Specialdoctor) *ExtradoctorUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return euo.AddSpecialdoctorIDs(ids...)
}

// AddOfficeIDs adds the offices edge to Office by ids.
func (euo *ExtradoctorUpdateOne) AddOfficeIDs(ids ...int) *ExtradoctorUpdateOne {
	euo.mutation.AddOfficeIDs(ids...)
	return euo
}

// AddOffices adds the offices edges to Office.
func (euo *ExtradoctorUpdateOne) AddOffices(o ...*Office) *ExtradoctorUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return euo.AddOfficeIDs(ids...)
}

// Mutation returns the ExtradoctorMutation object of the builder.
func (euo *ExtradoctorUpdateOne) Mutation() *ExtradoctorMutation {
	return euo.mutation
}

// RemoveSpecialdoctorIDs removes the specialdoctors edge to Specialdoctor by ids.
func (euo *ExtradoctorUpdateOne) RemoveSpecialdoctorIDs(ids ...int) *ExtradoctorUpdateOne {
	euo.mutation.RemoveSpecialdoctorIDs(ids...)
	return euo
}

// RemoveSpecialdoctors removes specialdoctors edges to Specialdoctor.
func (euo *ExtradoctorUpdateOne) RemoveSpecialdoctors(s ...*Specialdoctor) *ExtradoctorUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return euo.RemoveSpecialdoctorIDs(ids...)
}

// RemoveOfficeIDs removes the offices edge to Office by ids.
func (euo *ExtradoctorUpdateOne) RemoveOfficeIDs(ids ...int) *ExtradoctorUpdateOne {
	euo.mutation.RemoveOfficeIDs(ids...)
	return euo
}

// RemoveOffices removes offices edges to Office.
func (euo *ExtradoctorUpdateOne) RemoveOffices(o ...*Office) *ExtradoctorUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return euo.RemoveOfficeIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (euo *ExtradoctorUpdateOne) Save(ctx context.Context) (*Extradoctor, error) {
	if v, ok := euo.mutation.Specialname(); ok {
		if err := extradoctor.SpecialnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "specialname", err: fmt.Errorf("ent: validator failed for field \"specialname\": %w", err)}
		}
	}

	var (
		err  error
		node *Extradoctor
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExtradoctorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *ExtradoctorUpdateOne) SaveX(ctx context.Context) *Extradoctor {
	e, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return e
}

// Exec executes the query on the entity.
func (euo *ExtradoctorUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *ExtradoctorUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *ExtradoctorUpdateOne) sqlSave(ctx context.Context) (e *Extradoctor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   extradoctor.Table,
			Columns: extradoctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: extradoctor.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Extradoctor.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := euo.mutation.Specialname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: extradoctor.FieldSpecialname,
		})
	}
	if nodes := euo.mutation.RemovedSpecialdoctorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   extradoctor.SpecialdoctorsTable,
			Columns: []string{extradoctor.SpecialdoctorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: specialdoctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.SpecialdoctorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   extradoctor.SpecialdoctorsTable,
			Columns: []string{extradoctor.SpecialdoctorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: specialdoctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := euo.mutation.RemovedOfficesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   extradoctor.OfficesTable,
			Columns: []string{extradoctor.OfficesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: office.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.OfficesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   extradoctor.OfficesTable,
			Columns: []string{extradoctor.OfficesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: office.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	e = &Extradoctor{config: euo.config}
	_spec.Assign = e.assignValues
	_spec.ScanValues = e.scanValues()
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{extradoctor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return e, nil
}
