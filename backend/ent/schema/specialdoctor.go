package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Specialdoctor holds the schema definition for the Specialdoctor entity.
type Specialdoctor struct {
	ent.Schema
}

// Fields of the Specialdoctor.
func (Specialdoctor) Fields() []ent.Field {
	return []ent.Field{
		field.String("Roomnumber").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ABC]\\d{4}", s)
			if !match {
				return errors.New("กรุณากรอกหมายเลขห้องทำงาน")
			}
			return nil
		}),
		field.String("Doctorid").MaxLen(10).MinLen(10),
		field.String("Other").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๏]+$", s)
			if !match {
				return errors.New("กรุณากรอกเฉพาะภาษาไทย")
			}
			return nil
		}),
	}
}

// Edges of the Specialdoctor.
func (Specialdoctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("doctor", Doctor.Type).
			Ref("specialdoctors").Unique(),

		edge.From("department", Department.Type).
			Ref("specialdoctors").Unique(),

		edge.From("extradoctor", Extradoctor.Type).
			Ref("specialdoctors").Unique(),
	}
}
