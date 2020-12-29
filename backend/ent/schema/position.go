package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
)

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("position").NotEmpty(),
	}
 }

// Edges of the Position.
func (Position) Edges() []ent.Edge {
	return nil
}
