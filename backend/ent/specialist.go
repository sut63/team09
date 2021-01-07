// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team09/app/ent/specialist"
)

// Specialist is the model entity for the Specialist schema.
type Specialist struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Specialist holds the value of the "specialist" field.
	Specialist string `json:"specialist,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SpecialistQuery when eager-loading is set.
	Edges SpecialistEdges `json:"edges"`
}

// SpecialistEdges holds the relations/edges for other nodes in the graph.
type SpecialistEdges struct {
	// Specialdoctors holds the value of the specialdoctors edge.
	Specialdoctors []*Specialdoctor
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SpecialdoctorsOrErr returns the Specialdoctors value or an error if the edge
// was not loaded in eager-loading.
func (e SpecialistEdges) SpecialdoctorsOrErr() ([]*Specialdoctor, error) {
	if e.loadedTypes[0] {
		return e.Specialdoctors, nil
	}
	return nil, &NotLoadedError{edge: "specialdoctors"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Specialist) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // specialist
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Specialist fields.
func (s *Specialist) assignValues(values ...interface{}) error {
	if m, n := len(values), len(specialist.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field specialist", values[0])
	} else if value.Valid {
		s.Specialist = value.String
	}
	return nil
}

// QuerySpecialdoctors queries the specialdoctors edge of the Specialist.
func (s *Specialist) QuerySpecialdoctors() *SpecialdoctorQuery {
	return (&SpecialistClient{config: s.config}).QuerySpecialdoctors(s)
}

// Update returns a builder for updating this Specialist.
// Note that, you need to call Specialist.Unwrap() before calling this method, if this Specialist
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Specialist) Update() *SpecialistUpdateOne {
	return (&SpecialistClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Specialist) Unwrap() *Specialist {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Specialist is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Specialist) String() string {
	var builder strings.Builder
	builder.WriteString("Specialist(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", specialist=")
	builder.WriteString(s.Specialist)
	builder.WriteByte(')')
	return builder.String()
}

// Specialists is a parsable slice of Specialist.
type Specialists []*Specialist

func (s Specialists) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
