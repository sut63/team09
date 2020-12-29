package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	// "github.com/facebookincubator/ent/schema/edge"
)

// Doctor holds the schema definition for the Doctor entity.
type Doctor struct {
	ent.Schema
}

// Fields of the Doctor.
func (Doctor) Fields() []ent.Field {
	return []ent.Field{
		field.Int("name").Positive(),
	}
}

// Edges of the Doctor.
func (Doctor) Edges() []ent.Edge {
	return []ent.Edge{
            		
			// edge.From("title", Title.Type).
            // Ref("titles").
			// Unique(),
			
			// edge.From("gender", Gender.Type).
            // Ref("genders").
			// Unique(),
			
			// edge.From("position", Position.Type).
            // Ref("positions").
			// Unique(),
			
			// edge.From("disease", Disease.Type).
            // Ref("disease").
            // Unique(),
    }
}