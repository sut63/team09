// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team09/app/ent/course"
	"github.com/team09/app/ent/training"
)

// CourseCreate is the builder for creating a Course entity.
type CourseCreate struct {
	config
	mutation *CourseMutation
	hooks    []Hook
}

// SetNamecourse sets the namecourse field.
func (cc *CourseCreate) SetNamecourse(s string) *CourseCreate {
	cc.mutation.SetNamecourse(s)
	return cc
}

// AddTrainingIDs adds the trainings edge to Training by ids.
func (cc *CourseCreate) AddTrainingIDs(ids ...int) *CourseCreate {
	cc.mutation.AddTrainingIDs(ids...)
	return cc
}

// AddTrainings adds the trainings edges to Training.
func (cc *CourseCreate) AddTrainings(t ...*Training) *CourseCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cc.AddTrainingIDs(ids...)
}

// Mutation returns the CourseMutation object of the builder.
func (cc *CourseCreate) Mutation() *CourseMutation {
	return cc.mutation
}

// Save creates the Course in the database.
func (cc *CourseCreate) Save(ctx context.Context) (*Course, error) {
	if _, ok := cc.mutation.Namecourse(); !ok {
		return nil, &ValidationError{Name: "namecourse", err: errors.New("ent: missing required field \"namecourse\"")}
	}
	if v, ok := cc.mutation.Namecourse(); ok {
		if err := course.NamecourseValidator(v); err != nil {
			return nil, &ValidationError{Name: "namecourse", err: fmt.Errorf("ent: validator failed for field \"namecourse\": %w", err)}
		}
	}
	var (
		err  error
		node *Course
	)
	if len(cc.hooks) == 0 {
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CourseCreate) SaveX(ctx context.Context) *Course {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CourseCreate) sqlSave(ctx context.Context) (*Course, error) {
	c, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}

func (cc *CourseCreate) createSpec() (*Course, *sqlgraph.CreateSpec) {
	var (
		c     = &Course{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: course.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: course.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Namecourse(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: course.FieldNamecourse,
		})
		c.Namecourse = value
	}
	if nodes := cc.mutation.TrainingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.TrainingsTable,
			Columns: []string{course.TrainingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: training.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return c, _spec
}