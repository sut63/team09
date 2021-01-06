package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)
// Speacial_doctor holds the schema definition for the Speacial_doctor entity.
type Speacial_doctor struct {
	ent.Schema
}

// Fields of the Speacial_doctor.
func (Speacial_doctor) Fields() []ent.Field {
	return []ent.Field {
		field.String("name").NotEmpty(),
	}
}

// Edges of the Speacial_doctor.
func (Speacial_doctor) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("offices", Office.Type).
		StorageKey(edge.Column("speacial_doctor_id")),
	}
}
