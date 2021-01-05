// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team09/app/ent/doctor"
	"github.com/team09/app/ent/predicate"
)

// DoctorUpdate is the builder for updating Doctor entities.
type DoctorUpdate struct {
	config
	hooks      []Hook
	mutation   *DoctorMutation
	predicates []predicate.Doctor
}

// Where adds a new predicate for the builder.
func (du *DoctorUpdate) Where(ps ...predicate.Doctor) *DoctorUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetName sets the name field.
func (du *DoctorUpdate) SetName(i int) *DoctorUpdate {
	du.mutation.ResetName()
	du.mutation.SetName(i)
	return du
}

// AddName adds i to name.
func (du *DoctorUpdate) AddName(i int) *DoctorUpdate {
	du.mutation.AddName(i)
	return du
}

// SetAge sets the age field.
func (du *DoctorUpdate) SetAge(i int) *DoctorUpdate {
	du.mutation.ResetAge()
	du.mutation.SetAge(i)
	return du
}

// AddAge adds i to age.
func (du *DoctorUpdate) AddAge(i int) *DoctorUpdate {
	du.mutation.AddAge(i)
	return du
}

// SetEmail sets the email field.
func (du *DoctorUpdate) SetEmail(i int) *DoctorUpdate {
	du.mutation.ResetEmail()
	du.mutation.SetEmail(i)
	return du
}

// AddEmail adds i to email.
func (du *DoctorUpdate) AddEmail(i int) *DoctorUpdate {
	du.mutation.AddEmail(i)
	return du
}

// SetPnumber sets the pnumber field.
func (du *DoctorUpdate) SetPnumber(i int) *DoctorUpdate {
	du.mutation.ResetPnumber()
	du.mutation.SetPnumber(i)
	return du
}

// AddPnumber adds i to pnumber.
func (du *DoctorUpdate) AddPnumber(i int) *DoctorUpdate {
	du.mutation.AddPnumber(i)
	return du
}

// SetAddress sets the address field.
func (du *DoctorUpdate) SetAddress(i int) *DoctorUpdate {
	du.mutation.ResetAddress()
	du.mutation.SetAddress(i)
	return du
}

// AddAddress adds i to address.
func (du *DoctorUpdate) AddAddress(i int) *DoctorUpdate {
	du.mutation.AddAddress(i)
	return du
}

// SetEducational sets the educational field.
func (du *DoctorUpdate) SetEducational(i int) *DoctorUpdate {
	du.mutation.ResetEducational()
	du.mutation.SetEducational(i)
	return du
}

// AddEducational adds i to educational.
func (du *DoctorUpdate) AddEducational(i int) *DoctorUpdate {
	du.mutation.AddEducational(i)
	return du
}

// Mutation returns the DoctorMutation object of the builder.
func (du *DoctorUpdate) Mutation() *DoctorMutation {
	return du.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DoctorUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.Name(); ok {
		if err := doctor.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := du.mutation.Age(); ok {
		if err := doctor.AgeValidator(v); err != nil {
			return 0, &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if v, ok := du.mutation.Email(); ok {
		if err := doctor.EmailValidator(v); err != nil {
			return 0, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := du.mutation.Pnumber(); ok {
		if err := doctor.PnumberValidator(v); err != nil {
			return 0, &ValidationError{Name: "pnumber", err: fmt.Errorf("ent: validator failed for field \"pnumber\": %w", err)}
		}
	}
	if v, ok := du.mutation.Address(); ok {
		if err := doctor.AddressValidator(v); err != nil {
			return 0, &ValidationError{Name: "address", err: fmt.Errorf("ent: validator failed for field \"address\": %w", err)}
		}
	}
	if v, ok := du.mutation.Educational(); ok {
		if err := doctor.EducationalValidator(v); err != nil {
			return 0, &ValidationError{Name: "educational", err: fmt.Errorf("ent: validator failed for field \"educational\": %w", err)}
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
			mutation, ok := m.(*DoctorMutation)
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
func (du *DoctorUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DoctorUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DoctorUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DoctorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctor.Table,
			Columns: doctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
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
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldName,
		})
	}
	if value, ok := du.mutation.AddedName(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldName,
		})
	}
	if value, ok := du.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldAge,
		})
	}
	if value, ok := du.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldAge,
		})
	}
	if value, ok := du.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldEmail,
		})
	}
	if value, ok := du.mutation.AddedEmail(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldEmail,
		})
	}
	if value, ok := du.mutation.Pnumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldPnumber,
		})
	}
	if value, ok := du.mutation.AddedPnumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldPnumber,
		})
	}
	if value, ok := du.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldAddress,
		})
	}
	if value, ok := du.mutation.AddedAddress(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldAddress,
		})
	}
	if value, ok := du.mutation.Educational(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldEducational,
		})
	}
	if value, ok := du.mutation.AddedEducational(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldEducational,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doctor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DoctorUpdateOne is the builder for updating a single Doctor entity.
type DoctorUpdateOne struct {
	config
	hooks    []Hook
	mutation *DoctorMutation
}

// SetName sets the name field.
func (duo *DoctorUpdateOne) SetName(i int) *DoctorUpdateOne {
	duo.mutation.ResetName()
	duo.mutation.SetName(i)
	return duo
}

// AddName adds i to name.
func (duo *DoctorUpdateOne) AddName(i int) *DoctorUpdateOne {
	duo.mutation.AddName(i)
	return duo
}

// SetAge sets the age field.
func (duo *DoctorUpdateOne) SetAge(i int) *DoctorUpdateOne {
	duo.mutation.ResetAge()
	duo.mutation.SetAge(i)
	return duo
}

// AddAge adds i to age.
func (duo *DoctorUpdateOne) AddAge(i int) *DoctorUpdateOne {
	duo.mutation.AddAge(i)
	return duo
}

// SetEmail sets the email field.
func (duo *DoctorUpdateOne) SetEmail(i int) *DoctorUpdateOne {
	duo.mutation.ResetEmail()
	duo.mutation.SetEmail(i)
	return duo
}

// AddEmail adds i to email.
func (duo *DoctorUpdateOne) AddEmail(i int) *DoctorUpdateOne {
	duo.mutation.AddEmail(i)
	return duo
}

// SetPnumber sets the pnumber field.
func (duo *DoctorUpdateOne) SetPnumber(i int) *DoctorUpdateOne {
	duo.mutation.ResetPnumber()
	duo.mutation.SetPnumber(i)
	return duo
}

// AddPnumber adds i to pnumber.
func (duo *DoctorUpdateOne) AddPnumber(i int) *DoctorUpdateOne {
	duo.mutation.AddPnumber(i)
	return duo
}

// SetAddress sets the address field.
func (duo *DoctorUpdateOne) SetAddress(i int) *DoctorUpdateOne {
	duo.mutation.ResetAddress()
	duo.mutation.SetAddress(i)
	return duo
}

// AddAddress adds i to address.
func (duo *DoctorUpdateOne) AddAddress(i int) *DoctorUpdateOne {
	duo.mutation.AddAddress(i)
	return duo
}

// SetEducational sets the educational field.
func (duo *DoctorUpdateOne) SetEducational(i int) *DoctorUpdateOne {
	duo.mutation.ResetEducational()
	duo.mutation.SetEducational(i)
	return duo
}

// AddEducational adds i to educational.
func (duo *DoctorUpdateOne) AddEducational(i int) *DoctorUpdateOne {
	duo.mutation.AddEducational(i)
	return duo
}

// Mutation returns the DoctorMutation object of the builder.
func (duo *DoctorUpdateOne) Mutation() *DoctorMutation {
	return duo.mutation
}

// Save executes the query and returns the updated entity.
func (duo *DoctorUpdateOne) Save(ctx context.Context) (*Doctor, error) {
	if v, ok := duo.mutation.Name(); ok {
		if err := doctor.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Age(); ok {
		if err := doctor.AgeValidator(v); err != nil {
			return nil, &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Email(); ok {
		if err := doctor.EmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Pnumber(); ok {
		if err := doctor.PnumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "pnumber", err: fmt.Errorf("ent: validator failed for field \"pnumber\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Address(); ok {
		if err := doctor.AddressValidator(v); err != nil {
			return nil, &ValidationError{Name: "address", err: fmt.Errorf("ent: validator failed for field \"address\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Educational(); ok {
		if err := doctor.EducationalValidator(v); err != nil {
			return nil, &ValidationError{Name: "educational", err: fmt.Errorf("ent: validator failed for field \"educational\": %w", err)}
		}
	}
	var (
		err  error
		node *Doctor
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorMutation)
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
func (duo *DoctorUpdateOne) SaveX(ctx context.Context) *Doctor {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DoctorUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DoctorUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DoctorUpdateOne) sqlSave(ctx context.Context) (d *Doctor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctor.Table,
			Columns: doctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Doctor.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldName,
		})
	}
	if value, ok := duo.mutation.AddedName(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldName,
		})
	}
	if value, ok := duo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldAge,
		})
	}
	if value, ok := duo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldAge,
		})
	}
	if value, ok := duo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldEmail,
		})
	}
	if value, ok := duo.mutation.AddedEmail(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldEmail,
		})
	}
	if value, ok := duo.mutation.Pnumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldPnumber,
		})
	}
	if value, ok := duo.mutation.AddedPnumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldPnumber,
		})
	}
	if value, ok := duo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldAddress,
		})
	}
	if value, ok := duo.mutation.AddedAddress(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldAddress,
		})
	}
	if value, ok := duo.mutation.Educational(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldEducational,
		})
	}
	if value, ok := duo.mutation.AddedEducational(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: doctor.FieldEducational,
		})
	}
	d = &Doctor{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doctor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}