package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
)

// Gender holds the schema definition for the Gender entity.
type Gender struct {
	ent.Schema
}

// Fields of the Gender.
func (Gender) Fields() []ent.Field {
	return []ent.Field{
		field.String("gender").NotEmpty(),
	}
 }

// Edges of the Gender.
func (Gender) Edges() []ent.Edge {
	return nil
}
