package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Schedule holds the schema definition for the Schedule entity.
type Schedule struct {
	ent.Schema
}

// Fields of the Schedule.
func (Schedule) Fields() []ent.Field {
	return []ent.Field{
		field.String("activity").NotEmpty(),
		field.Time("added_time"),
	}
}

// Edges of the Schedule.
func (Schedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("docter", Doctor.Type).Ref("schedules").Unique(),
		edge.From("department", Department.Type).Ref("schedules").Unique(),
		edge.From("office", Office.Type).Ref("schedules").Unique(),
	}
}
