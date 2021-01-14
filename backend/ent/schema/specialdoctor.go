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
		field.String("other").NotEmpty(),
	}
}

// Edges of the Specialdoctor.
func (Specialdoctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("offices", Office.Type).
			StorageKey(edge.Column("Specialdoctor_id")),

		edge.From("doctor", Doctor.Type).
			Ref("specialdoctors").Unique(),
		
		edge.From("department", Department.Type).
			Ref("specialdoctors").Unique(),

		edge.From("extradoctor", Extradoctor.Type).
			Ref("specialdoctors").Unique(),
		
	}
}
