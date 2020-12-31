package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Office holds the schema definition for the Office entity.
type Office struct {
	ent.Schema
}

// Fields of the Office.
func (Office) Fields() []ent.Field {
	return []ent.Field {
		field.String("officename").NotEmpty(),
	}
}

// Edges of the Office.
func (Office) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("doctors", Doctor.Type).
		StorageKey(edge.Column("office_id")),
	}
}
