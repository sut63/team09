package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Workingtime holds the schema definition for the Workingtime entity.
type Workingtime struct {
	ent.Schema
}

// Fields of the Workingtime.
func (Workingtime) Fields() []ent.Field {
	return []ent.Field {
		field.Time("added_time1"),
		field.Time("added_time2"),
	}
}

// Edges of the Workingtime.
func (Workingtime) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("offices", Office.Type).
		StorageKey(edge.Column("workingtime_id")),
	}
}
 