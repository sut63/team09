package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Detail holds the schema definition for the Detail entity.
type Detail struct {
	ent.Schema
}

// Fields of the Detail.
func (Detail) Fields() []ent.Field {
	return []ent.Field{
		field.String("explain").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๏\\s]+$", s)
			if !match {
				return errors.New("กรอกรายละเอียดเป็นภาษาไทยเท่านั้น")
			}
			return nil
		}),
		field.String("phone").MaxLen(10).MinLen(10),
		field.String("email").Match(regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")),
		field.String("departmentid").MaxLen(3).MinLen(3),
	}
}

// Edges of the Detail.
func (Detail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).Ref("details").Unique(),
		edge.From("department", Department.Type).Ref("details").Unique(),
		edge.From("doctor", Doctor.Type).
            Ref("details").
            Unique().
            Required(),
	}
}
