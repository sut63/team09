package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("Detail").NotEmpty(),
		field.String("Name").NotEmpty(),

	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mission", Mission.Type).Ref("departments").Unique(),
		edge.From("doctor", Doctor.Type).Ref("departments").Unique(),

		edge.To("offices", Office.Type).
			StorageKey(edge.Column("department_id")),

		edge.To("schedules", Schedule.Type).
			StorageKey(edge.Column("department_id")),

		edge.To("trainings",Training.Type).
			StorageKey(edge.Column("department_id")),
		
		edge.From("specialist", Specialist.Type).
			Ref("departments").Unique(),
	}
}
