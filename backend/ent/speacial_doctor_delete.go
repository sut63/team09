// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team09/app/ent/predicate"
	"github.com/team09/app/ent/speacial_doctor"
)

// SpeacialDoctorDelete is the builder for deleting a SpeacialDoctor entity.
type SpeacialDoctorDelete struct {
	config
	hooks      []Hook
	mutation   *SpeacialDoctorMutation
	predicates []predicate.Speacial_doctor
}

// Where adds a new predicate to the delete builder.
func (sdd *SpeacialDoctorDelete) Where(ps ...predicate.Speacial_doctor) *SpeacialDoctorDelete {
	sdd.predicates = append(sdd.predicates, ps...)
	return sdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sdd *SpeacialDoctorDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sdd.hooks) == 0 {
		affected, err = sdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpeacialDoctorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sdd.mutation = mutation
			affected, err = sdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sdd.hooks) - 1; i >= 0; i-- {
			mut = sdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdd *SpeacialDoctorDelete) ExecX(ctx context.Context) int {
	n, err := sdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sdd *SpeacialDoctorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: speacial_doctor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: speacial_doctor.FieldID,
			},
		},
	}
	if ps := sdd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sdd.driver, _spec)
}

// SpeacialDoctorDeleteOne is the builder for deleting a single Speacial_doctor entity.
type SpeacialDoctorDeleteOne struct {
	sdd *SpeacialDoctorDelete
}

// Exec executes the deletion query.
func (sddo *SpeacialDoctorDeleteOne) Exec(ctx context.Context) error {
	n, err := sddo.sdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{speacial_doctor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sddo *SpeacialDoctorDeleteOne) ExecX(ctx context.Context) {
	sddo.sdd.ExecX(ctx)
}
