package schema
 
import (
	"github.com/facebookincubator/ent/schema/edge"
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
)

// Detail holds the schema definition for the Detail entity.
type Detail struct {
	ent.Schema
}

// Fields of the Detail.
func (Detail) Fields() []ent.Field {
	return []ent.Field{
		field.String("explain").NotEmpty(),
	}
 }

// Edges of the Detail.
func (Detail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course",Course.Type).Ref("details").Unique(),
		edge.From("mission", Mission.Type).Ref("details").Unique(),
		edge.From("department", Department.Type).Ref("details").Unique(),

	}
}