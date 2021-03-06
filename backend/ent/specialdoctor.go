// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/doctor"
	"github.com/team09/app/ent/extradoctor"
	"github.com/team09/app/ent/specialdoctor"
)

// Specialdoctor is the model entity for the Specialdoctor schema.
type Specialdoctor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Roomnumber holds the value of the "Roomnumber" field.
	Roomnumber string `json:"Roomnumber,omitempty"`
	// Doctorid holds the value of the "Doctorid" field.
	Doctorid string `json:"Doctorid,omitempty"`
	// Other holds the value of the "Other" field.
	Other string `json:"Other,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SpecialdoctorQuery when eager-loading is set.
	Edges          SpecialdoctorEdges `json:"edges"`
	department_id  *int
	doctor_id      *int
	extradoctor_id *int
}

// SpecialdoctorEdges holds the relations/edges for other nodes in the graph.
type SpecialdoctorEdges struct {
	// Doctor holds the value of the doctor edge.
	Doctor *Doctor
	// Department holds the value of the department edge.
	Department *Department
	// Extradoctor holds the value of the extradoctor edge.
	Extradoctor *Extradoctor
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DoctorOrErr returns the Doctor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SpecialdoctorEdges) DoctorOrErr() (*Doctor, error) {
	if e.loadedTypes[0] {
		if e.Doctor == nil {
			// The edge doctor was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: doctor.Label}
		}
		return e.Doctor, nil
	}
	return nil, &NotLoadedError{edge: "doctor"}
}

// DepartmentOrErr returns the Department value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SpecialdoctorEdges) DepartmentOrErr() (*Department, error) {
	if e.loadedTypes[1] {
		if e.Department == nil {
			// The edge department was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: department.Label}
		}
		return e.Department, nil
	}
	return nil, &NotLoadedError{edge: "department"}
}

// ExtradoctorOrErr returns the Extradoctor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SpecialdoctorEdges) ExtradoctorOrErr() (*Extradoctor, error) {
	if e.loadedTypes[2] {
		if e.Extradoctor == nil {
			// The edge extradoctor was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: extradoctor.Label}
		}
		return e.Extradoctor, nil
	}
	return nil, &NotLoadedError{edge: "extradoctor"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Specialdoctor) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Roomnumber
		&sql.NullString{}, // Doctorid
		&sql.NullString{}, // Other
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Specialdoctor) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // department_id
		&sql.NullInt64{}, // doctor_id
		&sql.NullInt64{}, // extradoctor_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Specialdoctor fields.
func (s *Specialdoctor) assignValues(values ...interface{}) error {
	if m, n := len(values), len(specialdoctor.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Roomnumber", values[0])
	} else if value.Valid {
		s.Roomnumber = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Doctorid", values[1])
	} else if value.Valid {
		s.Doctorid = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Other", values[2])
	} else if value.Valid {
		s.Other = value.String
	}
	values = values[3:]
	if len(values) == len(specialdoctor.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field department_id", value)
		} else if value.Valid {
			s.department_id = new(int)
			*s.department_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field doctor_id", value)
		} else if value.Valid {
			s.doctor_id = new(int)
			*s.doctor_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field extradoctor_id", value)
		} else if value.Valid {
			s.extradoctor_id = new(int)
			*s.extradoctor_id = int(value.Int64)
		}
	}
	return nil
}

// QueryDoctor queries the doctor edge of the Specialdoctor.
func (s *Specialdoctor) QueryDoctor() *DoctorQuery {
	return (&SpecialdoctorClient{config: s.config}).QueryDoctor(s)
}

// QueryDepartment queries the department edge of the Specialdoctor.
func (s *Specialdoctor) QueryDepartment() *DepartmentQuery {
	return (&SpecialdoctorClient{config: s.config}).QueryDepartment(s)
}

// QueryExtradoctor queries the extradoctor edge of the Specialdoctor.
func (s *Specialdoctor) QueryExtradoctor() *ExtradoctorQuery {
	return (&SpecialdoctorClient{config: s.config}).QueryExtradoctor(s)
}

// Update returns a builder for updating this Specialdoctor.
// Note that, you need to call Specialdoctor.Unwrap() before calling this method, if this Specialdoctor
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Specialdoctor) Update() *SpecialdoctorUpdateOne {
	return (&SpecialdoctorClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Specialdoctor) Unwrap() *Specialdoctor {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Specialdoctor is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Specialdoctor) String() string {
	var builder strings.Builder
	builder.WriteString("Specialdoctor(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", Roomnumber=")
	builder.WriteString(s.Roomnumber)
	builder.WriteString(", Doctorid=")
	builder.WriteString(s.Doctorid)
	builder.WriteString(", Other=")
	builder.WriteString(s.Other)
	builder.WriteByte(')')
	return builder.String()
}

// Specialdoctors is a parsable slice of Specialdoctor.
type Specialdoctors []*Specialdoctor

func (s Specialdoctors) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
