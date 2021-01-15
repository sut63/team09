package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("namecourse").NotEmpty(),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("trainings",Training.Type).StorageKey(edge.Column("course_id")),
		edge.To("details",Detail.Type).StorageKey(edge.Column("course_id")),
	}
}