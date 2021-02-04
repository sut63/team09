package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Schedule holds the schema definition for the Schedule entity.
type Schedule struct {
	ent.Schema
}

// Fields of the Schedule.
func (Schedule) Fields() []ent.Field {
	return []ent.Field{
		field.String("Activity").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๏\\s]+$", s)
			if !match {
				return errors.New("กรุณากรอกกิจกรรมของแพทย์")
			}
			return nil
		}),
		field.String("Roomnumber").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ABC]\\d{4}", s)
			if !match {
				return errors.New("กรุณากรอกหมายเลขห้องทำงาน")
			}
			return nil
		}),
		field.String("Docterid").MaxLen(10).MinLen(10),
		field.Time("added_time"),
	}
}

// Edges of the Schedule.
func (Schedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("doctor", Doctor.Type).Ref("schedules").Unique(),
		edge.From("department", Department.Type).Ref("schedules").Unique(),
		edge.From("office", Office.Type).Ref("schedules").Unique(),
	}
}
