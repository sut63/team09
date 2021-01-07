// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/doctor"
	"github.com/team09/app/ent/office"
	"github.com/team09/app/ent/special_doctor"
	"github.com/team09/app/ent/specialist"
)

// SpecialDoctorCreate is the builder for creating a Special_Doctor entity.
type SpecialDoctorCreate struct {
	config
	mutation *SpecialDoctorMutation
	hooks    []Hook
}

// SetOther sets the Other field.
func (sdc *SpecialDoctorCreate) SetOther(s string) *SpecialDoctorCreate {
	sdc.mutation.SetOther(s)
	return sdc
}

// SetDoctorID sets the doctor edge to Doctor by id.
func (sdc *SpecialDoctorCreate) SetDoctorID(id int) *SpecialDoctorCreate {
	sdc.mutation.SetDoctorID(id)
	return sdc
}

// SetNillableDoctorID sets the doctor edge to Doctor by id if the given value is not nil.
func (sdc *SpecialDoctorCreate) SetNillableDoctorID(id *int) *SpecialDoctorCreate {
	if id != nil {
		sdc = sdc.SetDoctorID(*id)
	}
	return sdc
}

// SetDoctor sets the doctor edge to Doctor.
func (sdc *SpecialDoctorCreate) SetDoctor(d *Doctor) *SpecialDoctorCreate {
	return sdc.SetDoctorID(d.ID)
}

// SetDepartmentID sets the department edge to Department by id.
func (sdc *SpecialDoctorCreate) SetDepartmentID(id int) *SpecialDoctorCreate {
	sdc.mutation.SetDepartmentID(id)
	return sdc
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (sdc *SpecialDoctorCreate) SetNillableDepartmentID(id *int) *SpecialDoctorCreate {
	if id != nil {
		sdc = sdc.SetDepartmentID(*id)
	}
	return sdc
}

// SetDepartment sets the department edge to Department.
func (sdc *SpecialDoctorCreate) SetDepartment(d *Department) *SpecialDoctorCreate {
	return sdc.SetDepartmentID(d.ID)
}

// SetSpecialistID sets the specialist edge to Specialist by id.
func (sdc *SpecialDoctorCreate) SetSpecialistID(id int) *SpecialDoctorCreate {
	sdc.mutation.SetSpecialistID(id)
	return sdc
}

// SetNillableSpecialistID sets the specialist edge to Specialist by id if the given value is not nil.
func (sdc *SpecialDoctorCreate) SetNillableSpecialistID(id *int) *SpecialDoctorCreate {
	if id != nil {
		sdc = sdc.SetSpecialistID(*id)
	}
	return sdc
}

// SetSpecialist sets the specialist edge to Specialist.
func (sdc *SpecialDoctorCreate) SetSpecialist(s *Specialist) *SpecialDoctorCreate {
	return sdc.SetSpecialistID(s.ID)
}

// AddOfficeIDs adds the offices edge to Office by ids.
func (sdc *SpecialDoctorCreate) AddOfficeIDs(ids ...int) *SpecialDoctorCreate {
	sdc.mutation.AddOfficeIDs(ids...)
	return sdc
}

// AddOffices adds the offices edges to Office.
func (sdc *SpecialDoctorCreate) AddOffices(o ...*Office) *SpecialDoctorCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return sdc.AddOfficeIDs(ids...)
}

// Mutation returns the SpecialDoctorMutation object of the builder.
func (sdc *SpecialDoctorCreate) Mutation() *SpecialDoctorMutation {
	return sdc.mutation
}

// Save creates the Special_Doctor in the database.
func (sdc *SpecialDoctorCreate) Save(ctx context.Context) (*Special_Doctor, error) {
	if _, ok := sdc.mutation.Other(); !ok {
		return nil, &ValidationError{Name: "Other", err: errors.New("ent: missing required field \"Other\"")}
	}
	if v, ok := sdc.mutation.Other(); ok {
		if err := special_doctor.OtherValidator(v); err != nil {
			return nil, &ValidationError{Name: "Other", err: fmt.Errorf("ent: validator failed for field \"Other\": %w", err)}
		}
	}
	var (
		err  error
		node *Special_Doctor
	)
	if len(sdc.hooks) == 0 {
		node, err = sdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecialDoctorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sdc.mutation = mutation
			node, err = sdc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sdc.hooks) - 1; i >= 0; i-- {
			mut = sdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sdc *SpecialDoctorCreate) SaveX(ctx context.Context) *Special_Doctor {
	v, err := sdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sdc *SpecialDoctorCreate) sqlSave(ctx context.Context) (*Special_Doctor, error) {
	sd, _spec := sdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	sd.ID = int(id)
	return sd, nil
}

func (sdc *SpecialDoctorCreate) createSpec() (*Special_Doctor, *sqlgraph.CreateSpec) {
	var (
		sd    = &Special_Doctor{config: sdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: special_doctor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: special_doctor.FieldID,
			},
		}
	)
	if value, ok := sdc.mutation.Other(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: special_doctor.FieldOther,
		})
		sd.Other = value
	}
	if nodes := sdc.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   special_doctor.DoctorTable,
			Columns: []string{special_doctor.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sdc.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   special_doctor.DepartmentTable,
			Columns: []string{special_doctor.DepartmentColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sdc.mutation.SpecialistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   special_doctor.SpecialistTable,
			Columns: []string{special_doctor.SpecialistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: specialist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sdc.mutation.OfficesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   special_doctor.OfficesTable,
			Columns: []string{special_doctor.OfficesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return sd, _spec
}
