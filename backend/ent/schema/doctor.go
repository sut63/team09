package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	 "github.com/facebookincubator/ent/schema/edge"
)

// Doctor holds the schema definition for the Doctor entity.
type Doctor struct {
	ent.Schema
}

// Fields of the Doctor.
func (Doctor) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("age").Positive(),
		field.String("email").NotEmpty(),
		field.Int("pnumber").Positive(),
		field.String("address").NotEmpty(),
		field.String("educational").NotEmpty(),
	}
}

// Edges of the Doctor.
func (Doctor) Edges() []ent.Edge {
	return []ent.Edge{

		 edge.From("title", Title.Type).
		 Ref("doctors").
		 Unique(),

		 edge.From("gender", Gender.Type).
		 Ref("doctors").
		 Unique(),

		 edge.From("position", Position.Type).
		 Ref("doctors").
		 Unique(),

		 edge.From("disease", Disease.Type).
		 Ref("doctors").
		 Unique(),

		 edge.To("departments",Department.Type).StorageKey(edge.Column("doctor_id")),
		 edge.To("offices",Office.Type).
		 StorageKey(edge.Column("doctor_id")),

		 edge.To("departments",Department.Type).
		 StorageKey(edge.Column("doctor_id")),

	}
}
