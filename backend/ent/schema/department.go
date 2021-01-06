package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field" 
	"github.com/facebookincubator/ent/schema/edge"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field {
		field.String("DepartmentType").NotEmpty(),
		field.String("Name").NotEmpty(),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("mission", Mission.Type).Ref("departments").Unique(),
		edge.From("doctor", Doctor.Type).Ref("departments").Unique(),
	}
}


