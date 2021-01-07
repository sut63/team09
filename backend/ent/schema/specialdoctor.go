package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Specialdoctor holds the schema definition for the Specialdoctor entity.
type Specialdoctor struct {
	ent.Schema
}

// Fields of the Specialdoctor.
func (Specialdoctor) Fields() []ent.Field {
	return []ent.Field{
		field.String("Other").NotEmpty(),
	}
}
// Edges of the Specialdoctor.
func (Specialdoctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("doctor", Doctor.Type).
		StorageKey(edge.Column("specialdoctor_id")),
			
		edge.From("department", Department.Type).
		Ref("specialdoctors").Unique(),

		edge.From("specialist", Specialist.Type).
		Ref("specialdoctors").
		Unique(),

		edge.To("offices",Office.Type).
		StorageKey(edge.Column("specialdoctor_id")),
	}
}
