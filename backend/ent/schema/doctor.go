package schema

import (
    "regexp"
    "errors"
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// Doctor holds the schema definition for the Doctor entity.
type Doctor struct {
    ent.Schema
}

// Fields of the Doctor.
func (Doctor) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").Validate(func(s string) error {
            match, _ := regexp.MatchString("^[ก-๏]+$",s)
                if !match {
                    return errors.New("กรุณากรอกชื่อเป็นภาษาไทยเท่านั้น")
                }
                return nil
            }),
        // field.String("name").NotEmpty(),
        field.Int("age").Min(0),
        field.String("email").Match(regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")),
        field.String("password").MaxLen(8).MinLen(8),
        field.String("address").NotEmpty(),
        field.String("educational").NotEmpty(),
        field.String("phone").MaxLen(10).MinLen(10),

    }
}

// Edges of the Doctor.
func (Doctor) Edges() []ent.Edge {
    return []ent.Edge{

        edge.From("title", Title.Type).
            Ref("doctors").
            Unique(),

        edge.From("gender", Gender.Type).
            Ref("doctors").
            Unique(),

        edge.From("position", Position.Type).
            Ref("doctors").
            Unique(),

        edge.From("disease", Disease.Type).
            Ref("doctors").
            Unique(),

        edge.To("offices", Office.Type).
            StorageKey(edge.Column("doctor_id")),

        edge.To("departments", Department.Type).
            StorageKey(edge.Column("doctor_id")),

        edge.To("schedules", Schedule.Type).
            StorageKey(edge.Column("schedule_id")),

        edge.To("trainings",Training.Type).
            StorageKey(edge.Column("doctor_id")),

        edge.To("specialdoctors",Specialdoctor.Type).
            StorageKey(edge.Column("doctor_id")),

    }
}
