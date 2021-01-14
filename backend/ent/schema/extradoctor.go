package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// extradoctor holds the schema definition for the Extradoctor entity.
type Extradoctor struct {
	ent.Schema
}

// Fields of the Extradoctor.
func (Extradoctor) Fields() []ent.Field {
	return []ent.Field{
		field.String("specialname").NotEmpty(),
	}
}

// Edges of the Extradoctor.
func (Extradoctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("specialdoctors",Specialdoctor.Type).
			StorageKey(edge.Column("extradoctor_id")),
	}

}
