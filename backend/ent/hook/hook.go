// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/team09/app/ent"
)

// The CourseFunc type is an adapter to allow the use of ordinary
// function as Course mutator.
type CourseFunc func(context.Context, *ent.CourseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CourseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CourseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CourseMutation", m)
	}
	return f(ctx, mv)
}

// The DepartmentFunc type is an adapter to allow the use of ordinary
// function as Department mutator.
type DepartmentFunc func(context.Context, *ent.DepartmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DepartmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DepartmentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DepartmentMutation", m)
	}
	return f(ctx, mv)
}

// The DiseaseFunc type is an adapter to allow the use of ordinary
// function as Disease mutator.
type DiseaseFunc func(context.Context, *ent.DiseaseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DiseaseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DiseaseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DiseaseMutation", m)
	}
	return f(ctx, mv)
}

// The DoctorFunc type is an adapter to allow the use of ordinary
// function as Doctor mutator.
type DoctorFunc func(context.Context, *ent.DoctorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DoctorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DoctorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DoctorMutation", m)
	}
	return f(ctx, mv)
}

// The GenderFunc type is an adapter to allow the use of ordinary
// function as Gender mutator.
type GenderFunc func(context.Context, *ent.GenderMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GenderFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GenderMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GenderMutation", m)
	}
	return f(ctx, mv)
}

// The MissionFunc type is an adapter to allow the use of ordinary
// function as Mission mutator.
type MissionFunc func(context.Context, *ent.MissionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MissionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MissionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MissionMutation", m)
	}
	return f(ctx, mv)
}

// The OfficeFunc type is an adapter to allow the use of ordinary
// function as Office mutator.
type OfficeFunc func(context.Context, *ent.OfficeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OfficeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OfficeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OfficeMutation", m)
	}
	return f(ctx, mv)
}

// The PositionFunc type is an adapter to allow the use of ordinary
// function as Position mutator.
type PositionFunc func(context.Context, *ent.PositionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PositionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PositionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PositionMutation", m)
	}
	return f(ctx, mv)
}

// The ScheduleFunc type is an adapter to allow the use of ordinary
// function as Schedule mutator.
type ScheduleFunc func(context.Context, *ent.ScheduleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScheduleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ScheduleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScheduleMutation", m)
	}
	return f(ctx, mv)
}

// The SpecialdoctorFunc type is an adapter to allow the use of ordinary
// function as Specialdoctor mutator.
type SpecialdoctorFunc func(context.Context, *ent.SpecialdoctorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SpecialdoctorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SpecialdoctorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SpecialdoctorMutation", m)
	}
	return f(ctx, mv)
}

// The SpecialistFunc type is an adapter to allow the use of ordinary
// function as Specialist mutator.
type SpecialistFunc func(context.Context, *ent.SpecialistMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SpecialistFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SpecialistMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SpecialistMutation", m)
	}
	return f(ctx, mv)
}

// The TitleFunc type is an adapter to allow the use of ordinary
// function as Title mutator.
type TitleFunc func(context.Context, *ent.TitleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TitleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TitleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TitleMutation", m)
	}
	return f(ctx, mv)
}

// The TrainingFunc type is an adapter to allow the use of ordinary
// function as Training mutator.
type TrainingFunc func(context.Context, *ent.TrainingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TrainingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TrainingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TrainingMutation", m)
	}
	return f(ctx, mv)
}

// The WorkingtimeFunc type is an adapter to allow the use of ordinary
// function as Workingtime mutator.
type WorkingtimeFunc func(context.Context, *ent.WorkingtimeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkingtimeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.WorkingtimeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkingtimeMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	Hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(_ context.Context, m ent.Mutation) (ent.Value, error) {
			return nil, fmt.Errorf("%s operation is not allowed", m.Op())
		})
	}
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
