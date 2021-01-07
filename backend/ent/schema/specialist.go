package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Specialist holds the schema definition for the Specialist entity.
type Specialist struct {
	ent.Schema
}

// Fields of the Specialist.
func (Specialist) Fields() []ent.Field {
	return []ent.Field{
		field.String("specialist").NotEmpty(),
	}
}

// Edges of the Specialist.
func (Specialist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("special_doctors",Special_Doctor.Type).
		StorageKey(edge.Column("specialist_id")),
	}
}
