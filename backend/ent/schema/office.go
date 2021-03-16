package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Office holds the schema definition for the Office entity.
type Office struct {
	ent.Schema
}

// Fields of the Office.
func (Office) Fields() []ent.Field {
	return []ent.Field{
		field.String("officename").Validate(func(s string) error {
		match, _ := regexp.MatchString("^[ก-๏]+$",s)
			if !match {
				return errors.New("กรุณากรอกสถานที่ทำงาน")
			}
			return nil
		}),
		field.String("roomnumber").Validate(func(s string) error {
			match, _ := regexp.MatchString("[ABC]\\d{4}",s)
				if !match {
					return errors.New("กรุณากรอกหมายเลขห้องทำงาน")
				}
				return nil
			}),
		field.String("doctoridcard").MaxLen(10).MinLen(10),
		field.Time("firsttime"),
		field.Time("finallytime"),
	}
}

// Edges of the Office.
func (Office) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("doctor", Doctor.Type).
			Ref("offices").Unique(),

		// edge.From("workingtime", Workingtime.Type).
		// 	Ref("offices").Unique(),

		edge.From("department", Department.Type).
			Ref("offices").Unique(),

		edge.From("extradoctor", Extradoctor.Type).
			Ref("offices").Unique(),

		edge.To("schedules", Schedule.Type).
			StorageKey(edge.Column("office_id")),
	}
}