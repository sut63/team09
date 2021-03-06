// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team09/app/ent/course"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/doctor"
	"github.com/team09/app/ent/training"
)

// TrainingCreate is the builder for creating a Training entity.
type TrainingCreate struct {
	config
	mutation *TrainingMutation
	hooks    []Hook
}

// SetTrainingplace sets the trainingplace field.
func (tc *TrainingCreate) SetTrainingplace(s string) *TrainingCreate {
	tc.mutation.SetTrainingplace(s)
	return tc
}

// SetFirstday sets the firstday field.
func (tc *TrainingCreate) SetFirstday(t time.Time) *TrainingCreate {
	tc.mutation.SetFirstday(t)
	return tc
}

// SetLastday sets the lastday field.
func (tc *TrainingCreate) SetLastday(t time.Time) *TrainingCreate {
	tc.mutation.SetLastday(t)
	return tc
}

// SetDoctoridcard sets the doctoridcard field.
func (tc *TrainingCreate) SetDoctoridcard(s string) *TrainingCreate {
	tc.mutation.SetDoctoridcard(s)
	return tc
}

// SetHour sets the hour field.
func (tc *TrainingCreate) SetHour(i int) *TrainingCreate {
	tc.mutation.SetHour(i)
	return tc
}

// SetCourseID sets the course edge to Course by id.
func (tc *TrainingCreate) SetCourseID(id int) *TrainingCreate {
	tc.mutation.SetCourseID(id)
	return tc
}

// SetNillableCourseID sets the course edge to Course by id if the given value is not nil.
func (tc *TrainingCreate) SetNillableCourseID(id *int) *TrainingCreate {
	if id != nil {
		tc = tc.SetCourseID(*id)
	}
	return tc
}

// SetCourse sets the course edge to Course.
func (tc *TrainingCreate) SetCourse(c *Course) *TrainingCreate {
	return tc.SetCourseID(c.ID)
}

// SetDoctorID sets the doctor edge to Doctor by id.
func (tc *TrainingCreate) SetDoctorID(id int) *TrainingCreate {
	tc.mutation.SetDoctorID(id)
	return tc
}

// SetNillableDoctorID sets the doctor edge to Doctor by id if the given value is not nil.
func (tc *TrainingCreate) SetNillableDoctorID(id *int) *TrainingCreate {
	if id != nil {
		tc = tc.SetDoctorID(*id)
	}
	return tc
}

// SetDoctor sets the doctor edge to Doctor.
func (tc *TrainingCreate) SetDoctor(d *Doctor) *TrainingCreate {
	return tc.SetDoctorID(d.ID)
}

// SetDepartmentID sets the department edge to Department by id.
func (tc *TrainingCreate) SetDepartmentID(id int) *TrainingCreate {
	tc.mutation.SetDepartmentID(id)
	return tc
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (tc *TrainingCreate) SetNillableDepartmentID(id *int) *TrainingCreate {
	if id != nil {
		tc = tc.SetDepartmentID(*id)
	}
	return tc
}

// SetDepartment sets the department edge to Department.
func (tc *TrainingCreate) SetDepartment(d *Department) *TrainingCreate {
	return tc.SetDepartmentID(d.ID)
}

// Mutation returns the TrainingMutation object of the builder.
func (tc *TrainingCreate) Mutation() *TrainingMutation {
	return tc.mutation
}

// Save creates the Training in the database.
func (tc *TrainingCreate) Save(ctx context.Context) (*Training, error) {
	if _, ok := tc.mutation.Trainingplace(); !ok {
		return nil, &ValidationError{Name: "trainingplace", err: errors.New("ent: missing required field \"trainingplace\"")}
	}
	if v, ok := tc.mutation.Trainingplace(); ok {
		if err := training.TrainingplaceValidator(v); err != nil {
			return nil, &ValidationError{Name: "trainingplace", err: fmt.Errorf("ent: validator failed for field \"trainingplace\": %w", err)}
		}
	}
	if _, ok := tc.mutation.Firstday(); !ok {
		return nil, &ValidationError{Name: "firstday", err: errors.New("ent: missing required field \"firstday\"")}
	}
	if _, ok := tc.mutation.Lastday(); !ok {
		return nil, &ValidationError{Name: "lastday", err: errors.New("ent: missing required field \"lastday\"")}
	}
	if _, ok := tc.mutation.Doctoridcard(); !ok {
		return nil, &ValidationError{Name: "doctoridcard", err: errors.New("ent: missing required field \"doctoridcard\"")}
	}
	if v, ok := tc.mutation.Doctoridcard(); ok {
		if err := training.DoctoridcardValidator(v); err != nil {
			return nil, &ValidationError{Name: "doctoridcard", err: fmt.Errorf("ent: validator failed for field \"doctoridcard\": %w", err)}
		}
	}
	if _, ok := tc.mutation.Hour(); !ok {
		return nil, &ValidationError{Name: "hour", err: errors.New("ent: missing required field \"hour\"")}
	}
	if v, ok := tc.mutation.Hour(); ok {
		if err := training.HourValidator(v); err != nil {
			return nil, &ValidationError{Name: "hour", err: fmt.Errorf("ent: validator failed for field \"hour\": %w", err)}
		}
	}
	var (
		err  error
		node *Training
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrainingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TrainingCreate) SaveX(ctx context.Context) *Training {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TrainingCreate) sqlSave(ctx context.Context) (*Training, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *TrainingCreate) createSpec() (*Training, *sqlgraph.CreateSpec) {
	var (
		t     = &Training{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: training.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: training.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Trainingplace(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: training.FieldTrainingplace,
		})
		t.Trainingplace = value
	}
	if value, ok := tc.mutation.Firstday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: training.FieldFirstday,
		})
		t.Firstday = value
	}
	if value, ok := tc.mutation.Lastday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: training.FieldLastday,
		})
		t.Lastday = value
	}
	if value, ok := tc.mutation.Doctoridcard(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: training.FieldDoctoridcard,
		})
		t.Doctoridcard = value
	}
	if value, ok := tc.mutation.Hour(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: training.FieldHour,
		})
		t.Hour = value
	}
	if nodes := tc.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   training.CourseTable,
			Columns: []string{training.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   training.DoctorTable,
			Columns: []string{training.DoctorColumn},
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
	if nodes := tc.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   training.DepartmentTable,
			Columns: []string{training.DepartmentColumn},
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
	return t, _spec
}
