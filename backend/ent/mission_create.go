// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/detail"
	"github.com/team09/app/ent/mission"
)

// MissionCreate is the builder for creating a Mission entity.
type MissionCreate struct {
	config
	mutation *MissionMutation
	hooks    []Hook
}

// SetMission sets the mission field.
func (mc *MissionCreate) SetMission(s string) *MissionCreate {
	mc.mutation.SetMission(s)
	return mc
}

// AddDepartmentIDs adds the departments edge to Department by ids.
func (mc *MissionCreate) AddDepartmentIDs(ids ...int) *MissionCreate {
	mc.mutation.AddDepartmentIDs(ids...)
	return mc
}

// AddDepartments adds the departments edges to Department.
func (mc *MissionCreate) AddDepartments(d ...*Department) *MissionCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mc.AddDepartmentIDs(ids...)
}

// AddDetailIDs adds the details edge to Detail by ids.
func (mc *MissionCreate) AddDetailIDs(ids ...int) *MissionCreate {
	mc.mutation.AddDetailIDs(ids...)
	return mc
}

// AddDetails adds the details edges to Detail.
func (mc *MissionCreate) AddDetails(d ...*Detail) *MissionCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mc.AddDetailIDs(ids...)
}

// Mutation returns the MissionMutation object of the builder.
func (mc *MissionCreate) Mutation() *MissionMutation {
	return mc.mutation
}

// Save creates the Mission in the database.
func (mc *MissionCreate) Save(ctx context.Context) (*Mission, error) {
	if _, ok := mc.mutation.Mission(); !ok {
		return nil, &ValidationError{Name: "mission", err: errors.New("ent: missing required field \"mission\"")}
	}
	if v, ok := mc.mutation.Mission(); ok {
		if err := mission.MissionValidator(v); err != nil {
			return nil, &ValidationError{Name: "mission", err: fmt.Errorf("ent: validator failed for field \"mission\": %w", err)}
		}
	}
	var (
		err  error
		node *Mission
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MissionCreate) SaveX(ctx context.Context) *Mission {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MissionCreate) sqlSave(ctx context.Context) (*Mission, error) {
	m, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}

func (mc *MissionCreate) createSpec() (*Mission, *sqlgraph.CreateSpec) {
	var (
		m     = &Mission{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mission.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mission.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Mission(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mission.FieldMission,
		})
		m.Mission = value
	}
	if nodes := mc.mutation.DepartmentsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.DetailsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return m, _spec
}
