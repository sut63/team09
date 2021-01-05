// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team09/app/ent/disease"
	"github.com/team09/app/ent/predicate"
)

// DiseaseUpdate is the builder for updating Disease entities.
type DiseaseUpdate struct {
	config
	hooks      []Hook
	mutation   *DiseaseMutation
	predicates []predicate.Disease
}

// Where adds a new predicate for the builder.
func (du *DiseaseUpdate) Where(ps ...predicate.Disease) *DiseaseUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDisease sets the disease field.
func (du *DiseaseUpdate) SetDisease(s string) *DiseaseUpdate {
	du.mutation.SetDisease(s)
	return du
}

// Mutation returns the DiseaseMutation object of the builder.
func (du *DiseaseUpdate) Mutation() *DiseaseMutation {
	return du.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DiseaseUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.Disease(); ok {
		if err := disease.DiseaseValidator(v); err != nil {
			return 0, &ValidationError{Name: "disease", err: fmt.Errorf("ent: validator failed for field \"disease\": %w", err)}
		}
	}
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiseaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DiseaseUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DiseaseUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DiseaseUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DiseaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   disease.Table,
			Columns: disease.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: disease.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Disease(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: disease.FieldDisease,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{disease.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DiseaseUpdateOne is the builder for updating a single Disease entity.
type DiseaseUpdateOne struct {
	config
	hooks    []Hook
	mutation *DiseaseMutation
}

// SetDisease sets the disease field.
func (duo *DiseaseUpdateOne) SetDisease(s string) *DiseaseUpdateOne {
	duo.mutation.SetDisease(s)
	return duo
}

// Mutation returns the DiseaseMutation object of the builder.
func (duo *DiseaseUpdateOne) Mutation() *DiseaseMutation {
	return duo.mutation
}

// Save executes the query and returns the updated entity.
func (duo *DiseaseUpdateOne) Save(ctx context.Context) (*Disease, error) {
	if v, ok := duo.mutation.Disease(); ok {
		if err := disease.DiseaseValidator(v); err != nil {
			return nil, &ValidationError{Name: "disease", err: fmt.Errorf("ent: validator failed for field \"disease\": %w", err)}
		}
	}
	var (
		err  error
		node *Disease
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiseaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DiseaseUpdateOne) SaveX(ctx context.Context) *Disease {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DiseaseUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DiseaseUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DiseaseUpdateOne) sqlSave(ctx context.Context) (d *Disease, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   disease.Table,
			Columns: disease.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: disease.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Disease.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Disease(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: disease.FieldDisease,
		})
	}
	d = &Disease{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{disease.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
