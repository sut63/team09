package schema

import (
	"errors"
	"regexp"

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
		field.String("trainingplace").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๏]+$", s)
			if !match {
				return errors.New("กรอกสถานที่เข้าร่วมอบรมให้เป็นภาษาไทยเท่านั้น")
			}
			return nil
		}),
		field.Time("firstday"),
		field.Time("lastday"),
		field.String("doctoridcard").MaxLen(10).MinLen(10),
		field.Int("hour").Range(1,100),
	}
}

// Edges of the Training.
func (Training) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).Ref("trainings").Unique(),
		edge.From("doctor", Doctor.Type).Ref("trainings").Unique(),
		edge.From("department", Department.Type).Ref("trainings").Unique(),
	}

}
