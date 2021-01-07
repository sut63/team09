package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Office holds the schema definition for the Office entity.
type Office struct {
	ent.Schema
}

// Fields of the Office.
func (Office) Fields() []ent.Field {
	return []ent.Field{
		field.String("officename").NotEmpty(),
	}
}

// Edges of the Office.
func (Office) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("doctor", Doctor.Type).
			Ref("offices").Unique(),

		edge.From("workingtime", Workingtime.Type).
			Ref("offices").Unique(),

		edge.From("department", Department.Type).
			Ref("offices").Unique(),

		edge.From("speacial_doctor", Special_Doctor.Type).
			Ref("offices").Unique(),

		edge.To("schedules", Schedule.Type).
			StorageKey(edge.Column("office_id")),
	}
}
