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
		field.String("branch").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๏]+$", s)
			if !match {
				return errors.New("กรอกสาขาให้เป็นภาษาไทยเท่านั้น")
			}
			return nil
		}),
		field.Time("dateone"),
		field.Time("datetwo"),
		field.String("doctoridcard").MaxLen(10).MinLen(10),
		field.String("hour").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[0-9]+$", s)
			if !match {
				return errors.New("กรอกชั่วโมงเป็นตัวเลขเท่านั้น")
			}
			return nil
		}),
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
