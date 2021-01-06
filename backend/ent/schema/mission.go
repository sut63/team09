package schema
 
import (
	"github.com/facebookincubator/ent/schema/edge"
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
)

// Mission holds the schema definition for the Mission entity.
type Mission struct {
	ent.Schema
}

// Fields of the Mission.
func (Mission) Fields() []ent.Field {
	return []ent.Field{
		field.String("MissionType").NotEmpty(),
	}
 }

// Edges of the Mission.
func (Mission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("departments",Department.Type).StorageKey(edge.Column("mission_id")),

		// edge.To("drugs",Drug.Type).StorageKey(edge.Column("unit_id")),
	}
}
