// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/doctor"
	"github.com/team09/app/ent/office"
	"github.com/team09/app/ent/specialist"
	"github.com/team09/app/ent/workingtime"
)

// Office is the model entity for the Office schema.
type Office struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Officename holds the value of the "officename" field.
	Officename string `json:"officename,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OfficeQuery when eager-loading is set.
	Edges          OfficeEdges `json:"edges"`
	department_id  *int
	doctor_id      *int
	specialist_id  *int
	workingtime_id *int
}

// OfficeEdges holds the relations/edges for other nodes in the graph.
type OfficeEdges struct {
	// Doctor holds the value of the doctor edge.
	Doctor *Doctor
	// Workingtime holds the value of the workingtime edge.
	Workingtime *Workingtime
	// Department holds the value of the department edge.
	Department *Department
	// Specialist holds the value of the specialist edge.
	Specialist *Specialist
	// Schedules holds the value of the schedules edge.
	Schedules []*Schedule
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// DoctorOrErr returns the Doctor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OfficeEdges) DoctorOrErr() (*Doctor, error) {
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

// WorkingtimeOrErr returns the Workingtime value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OfficeEdges) WorkingtimeOrErr() (*Workingtime, error) {
	if e.loadedTypes[1] {
		if e.Workingtime == nil {
			// The edge workingtime was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workingtime.Label}
		}
		return e.Workingtime, nil
	}
	return nil, &NotLoadedError{edge: "workingtime"}
}

// DepartmentOrErr returns the Department value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OfficeEdges) DepartmentOrErr() (*Department, error) {
	if e.loadedTypes[2] {
		if e.Department == nil {
			// The edge department was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: department.Label}
		}
		return e.Department, nil
	}
	return nil, &NotLoadedError{edge: "department"}
}

// SpecialistOrErr returns the Specialist value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OfficeEdges) SpecialistOrErr() (*Specialist, error) {
	if e.loadedTypes[3] {
		if e.Specialist == nil {
			// The edge specialist was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: specialist.Label}
		}
		return e.Specialist, nil
	}
	return nil, &NotLoadedError{edge: "specialist"}
}

// SchedulesOrErr returns the Schedules value or an error if the edge
// was not loaded in eager-loading.
func (e OfficeEdges) SchedulesOrErr() ([]*Schedule, error) {
	if e.loadedTypes[4] {
		return e.Schedules, nil
	}
	return nil, &NotLoadedError{edge: "schedules"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Office) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // officename
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Office) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // department_id
		&sql.NullInt64{}, // doctor_id
		&sql.NullInt64{}, // specialist_id
		&sql.NullInt64{}, // workingtime_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Office fields.
func (o *Office) assignValues(values ...interface{}) error {
	if m, n := len(values), len(office.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	o.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field officename", values[0])
	} else if value.Valid {
		o.Officename = value.String
	}
	values = values[1:]
	if len(values) == len(office.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field department_id", value)
		} else if value.Valid {
			o.department_id = new(int)
			*o.department_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field doctor_id", value)
		} else if value.Valid {
			o.doctor_id = new(int)
			*o.doctor_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field specialist_id", value)
		} else if value.Valid {
			o.specialist_id = new(int)
			*o.specialist_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field workingtime_id", value)
		} else if value.Valid {
			o.workingtime_id = new(int)
			*o.workingtime_id = int(value.Int64)
		}
	}
	return nil
}

// QueryDoctor queries the doctor edge of the Office.
func (o *Office) QueryDoctor() *DoctorQuery {
	return (&OfficeClient{config: o.config}).QueryDoctor(o)
}

// QueryWorkingtime queries the workingtime edge of the Office.
func (o *Office) QueryWorkingtime() *WorkingtimeQuery {
	return (&OfficeClient{config: o.config}).QueryWorkingtime(o)
}

// QueryDepartment queries the department edge of the Office.
func (o *Office) QueryDepartment() *DepartmentQuery {
	return (&OfficeClient{config: o.config}).QueryDepartment(o)
}

// QuerySpecialist queries the specialist edge of the Office.
func (o *Office) QuerySpecialist() *SpecialistQuery {
	return (&OfficeClient{config: o.config}).QuerySpecialist(o)
}

// QuerySchedules queries the schedules edge of the Office.
func (o *Office) QuerySchedules() *ScheduleQuery {
	return (&OfficeClient{config: o.config}).QuerySchedules(o)
}

// Update returns a builder for updating this Office.
// Note that, you need to call Office.Unwrap() before calling this method, if this Office
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Office) Update() *OfficeUpdateOne {
	return (&OfficeClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (o *Office) Unwrap() *Office {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Office is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Office) String() string {
	var builder strings.Builder
	builder.WriteString("Office(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", officename=")
	builder.WriteString(o.Officename)
	builder.WriteByte(')')
	return builder.String()
}

// Offices is a parsable slice of Office.
type Offices []*Office

func (o Offices) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
