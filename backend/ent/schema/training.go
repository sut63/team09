package schema

import (
	"time"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Training holds the schema definition for the Training entity.
type Training struct {
	ent.Schema
}

// Fields of the Training.
func (Training) Fields() []ent.Field {
	return []ent.Field{
		field.String("branch").NotEmpty(),
		field.Time("dateone").Default(time.Now).Immutable(),
		field.Time("datetwo").Default(time.Now).Immutable(),
		 
    }

}

// Edges of the Training.
func (Training) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("course",Course.Type).Ref("trainings").Unique(),
		edge.From("doctor", Doctor.Type).Ref("trainings").Unique(),
		edge.From("department", Department.Type).Ref("trainings").Unique(),
		      
    }

}
