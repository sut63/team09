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
	"github.com/team09/app/ent/schedule"
	"github.com/team09/app/ent/specialist"
	"github.com/team09/app/ent/workingtime"
)

// OfficeCreate is the builder for creating a Office entity.
type OfficeCreate struct {
	config
	mutation *OfficeMutation
	hooks    []Hook
}

// SetOfficename sets the officename field.
func (oc *OfficeCreate) SetOfficename(s string) *OfficeCreate {
	oc.mutation.SetOfficename(s)
	return oc
}

// SetDoctorID sets the doctor edge to Doctor by id.
func (oc *OfficeCreate) SetDoctorID(id int) *OfficeCreate {
	oc.mutation.SetDoctorID(id)
	return oc
}

// SetNillableDoctorID sets the doctor edge to Doctor by id if the given value is not nil.
func (oc *OfficeCreate) SetNillableDoctorID(id *int) *OfficeCreate {
	if id != nil {
		oc = oc.SetDoctorID(*id)
	}
	return oc
}

// SetDoctor sets the doctor edge to Doctor.
func (oc *OfficeCreate) SetDoctor(d *Doctor) *OfficeCreate {
	return oc.SetDoctorID(d.ID)
}

// SetWorkingtimeID sets the workingtime edge to Workingtime by id.
func (oc *OfficeCreate) SetWorkingtimeID(id int) *OfficeCreate {
	oc.mutation.SetWorkingtimeID(id)
	return oc
}

// SetNillableWorkingtimeID sets the workingtime edge to Workingtime by id if the given value is not nil.
func (oc *OfficeCreate) SetNillableWorkingtimeID(id *int) *OfficeCreate {
	if id != nil {
		oc = oc.SetWorkingtimeID(*id)
	}
	return oc
}

// SetWorkingtime sets the workingtime edge to Workingtime.
func (oc *OfficeCreate) SetWorkingtime(w *Workingtime) *OfficeCreate {
	return oc.SetWorkingtimeID(w.ID)
}

// SetDepartmentID sets the department edge to Department by id.
func (oc *OfficeCreate) SetDepartmentID(id int) *OfficeCreate {
	oc.mutation.SetDepartmentID(id)
	return oc
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (oc *OfficeCreate) SetNillableDepartmentID(id *int) *OfficeCreate {
	if id != nil {
		oc = oc.SetDepartmentID(*id)
	}
	return oc
}

// SetDepartment sets the department edge to Department.
func (oc *OfficeCreate) SetDepartment(d *Department) *OfficeCreate {
	return oc.SetDepartmentID(d.ID)
}

// SetSpecialistID sets the specialist edge to Specialist by id.
func (oc *OfficeCreate) SetSpecialistID(id int) *OfficeCreate {
	oc.mutation.SetSpecialistID(id)
	return oc
}

// SetNillableSpecialistID sets the specialist edge to Specialist by id if the given value is not nil.
func (oc *OfficeCreate) SetNillableSpecialistID(id *int) *OfficeCreate {
	if id != nil {
		oc = oc.SetSpecialistID(*id)
	}
	return oc
}

// SetSpecialist sets the specialist edge to Specialist.
func (oc *OfficeCreate) SetSpecialist(s *Specialist) *OfficeCreate {
	return oc.SetSpecialistID(s.ID)
}

// AddScheduleIDs adds the schedules edge to Schedule by ids.
func (oc *OfficeCreate) AddScheduleIDs(ids ...int) *OfficeCreate {
	oc.mutation.AddScheduleIDs(ids...)
	return oc
}

// AddSchedules adds the schedules edges to Schedule.
func (oc *OfficeCreate) AddSchedules(s ...*Schedule) *OfficeCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return oc.AddScheduleIDs(ids...)
}

// Mutation returns the OfficeMutation object of the builder.
func (oc *OfficeCreate) Mutation() *OfficeMutation {
	return oc.mutation
}

// Save creates the Office in the database.
func (oc *OfficeCreate) Save(ctx context.Context) (*Office, error) {
	if _, ok := oc.mutation.Officename(); !ok {
		return nil, &ValidationError{Name: "officename", err: errors.New("ent: missing required field \"officename\"")}
	}
	if v, ok := oc.mutation.Officename(); ok {
		if err := office.OfficenameValidator(v); err != nil {
			return nil, &ValidationError{Name: "officename", err: fmt.Errorf("ent: validator failed for field \"officename\": %w", err)}
		}
	}
	var (
		err  error
		node *Office
	)
	if len(oc.hooks) == 0 {
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OfficeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oc.mutation = mutation
			node, err = oc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OfficeCreate) SaveX(ctx context.Context) *Office {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oc *OfficeCreate) sqlSave(ctx context.Context) (*Office, error) {
	o, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	o.ID = int(id)
	return o, nil
}

func (oc *OfficeCreate) createSpec() (*Office, *sqlgraph.CreateSpec) {
	var (
		o     = &Office{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: office.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: office.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.Officename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: office.FieldOfficename,
		})
		o.Officename = value
	}
	if nodes := oc.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   office.DoctorTable,
			Columns: []string{office.DoctorColumn},
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
	if nodes := oc.mutation.WorkingtimeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   office.WorkingtimeTable,
			Columns: []string{office.WorkingtimeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workingtime.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   office.DepartmentTable,
			Columns: []string{office.DepartmentColumn},
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
	if nodes := oc.mutation.SpecialistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   office.SpecialistTable,
			Columns: []string{office.SpecialistColumn},
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
	if nodes := oc.mutation.SchedulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   office.SchedulesTable,
			Columns: []string{office.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: schedule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return o, _spec
}
