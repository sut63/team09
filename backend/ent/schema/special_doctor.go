package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Special_Doctor holds the schema definition for the Special_Doctor entity.
type Special_Doctor struct {
	ent.Schema
}

// Fields of the Special_Doctor.
func (Special_Doctor) Fields() []ent.Field {
	return []ent.Field{
		field.String("Other").NotEmpty(),
	}
}

// Edges of the Special_Doctor.
func (Special_Doctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("doctor", Doctor.Type).
		Unique(),
			
		edge.From("department", Department.Type).
		Ref("special_doctors").Unique(),

		edge.From("specialist", Specialist.Type).
		Ref("special_doctors").
		Unique(),

		edge.To("offices",Office.Type).
		StorageKey(edge.Column("special_doctor_id")),


	}

}
