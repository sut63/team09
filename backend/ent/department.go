// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team09/app/ent/department"
	"github.com/team09/app/ent/doctor"
	"github.com/team09/app/ent/mission"
)

// Department is the model entity for the Department schema.
type Department struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Detail holds the value of the "Detail" field.
	Detail string `json:"Detail,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DepartmentQuery when eager-loading is set.
	Edges      DepartmentEdges `json:"edges"`
	doctor_id  *int
	mission_id *int
}

// DepartmentEdges holds the relations/edges for other nodes in the graph.
type DepartmentEdges struct {
	// Mission holds the value of the mission edge.
	Mission *Mission
	// Doctor holds the value of the doctor edge.
	Doctor *Doctor
	// Offices holds the value of the offices edge.
	Offices []*Office
	// Schedules holds the value of the schedules edge.
	Schedules []*Schedule
	// Trainings holds the value of the trainings edge.
	Trainings []*Training
	// Specialdoctors holds the value of the specialdoctors edge.
	Specialdoctors []*Specialdoctor
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// MissionOrErr returns the Mission value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DepartmentEdges) MissionOrErr() (*Mission, error) {
	if e.loadedTypes[0] {
		if e.Mission == nil {
			// The edge mission was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: mission.Label}
		}
		return e.Mission, nil
	}
	return nil, &NotLoadedError{edge: "mission"}
}

// DoctorOrErr returns the Doctor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DepartmentEdges) DoctorOrErr() (*Doctor, error) {
	if e.loadedTypes[1] {
		if e.Doctor == nil {
			// The edge doctor was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: doctor.Label}
		}
		return e.Doctor, nil
	}
	return nil, &NotLoadedError{edge: "doctor"}
}

// OfficesOrErr returns the Offices value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) OfficesOrErr() ([]*Office, error) {
	if e.loadedTypes[2] {
		return e.Offices, nil
	}
	return nil, &NotLoadedError{edge: "offices"}
}

// SchedulesOrErr returns the Schedules value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) SchedulesOrErr() ([]*Schedule, error) {
	if e.loadedTypes[3] {
		return e.Schedules, nil
	}
	return nil, &NotLoadedError{edge: "schedules"}
}

// TrainingsOrErr returns the Trainings value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) TrainingsOrErr() ([]*Training, error) {
	if e.loadedTypes[4] {
		return e.Trainings, nil
	}
	return nil, &NotLoadedError{edge: "trainings"}
}

// SpecialdoctorsOrErr returns the Specialdoctors value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) SpecialdoctorsOrErr() ([]*Specialdoctor, error) {
	if e.loadedTypes[5] {
		return e.Specialdoctors, nil
	}
	return nil, &NotLoadedError{edge: "specialdoctors"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Department) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Detail
		&sql.NullString{}, // Name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Department) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // doctor_id
		&sql.NullInt64{}, // mission_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Department fields.
func (d *Department) assignValues(values ...interface{}) error {
	if m, n := len(values), len(department.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Detail", values[0])
	} else if value.Valid {
		d.Detail = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Name", values[1])
	} else if value.Valid {
		d.Name = value.String
	}
	values = values[2:]
	if len(values) == len(department.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field doctor_id", value)
		} else if value.Valid {
			d.doctor_id = new(int)
			*d.doctor_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field mission_id", value)
		} else if value.Valid {
			d.mission_id = new(int)
			*d.mission_id = int(value.Int64)
		}
	}
	return nil
}

// QueryMission queries the mission edge of the Department.
func (d *Department) QueryMission() *MissionQuery {
	return (&DepartmentClient{config: d.config}).QueryMission(d)
}

// QueryDoctor queries the doctor edge of the Department.
func (d *Department) QueryDoctor() *DoctorQuery {
	return (&DepartmentClient{config: d.config}).QueryDoctor(d)
}

// QueryOffices queries the offices edge of the Department.
func (d *Department) QueryOffices() *OfficeQuery {
	return (&DepartmentClient{config: d.config}).QueryOffices(d)
}

// QuerySchedules queries the schedules edge of the Department.
func (d *Department) QuerySchedules() *ScheduleQuery {
	return (&DepartmentClient{config: d.config}).QuerySchedules(d)
}

// QueryTrainings queries the trainings edge of the Department.
func (d *Department) QueryTrainings() *TrainingQuery {
	return (&DepartmentClient{config: d.config}).QueryTrainings(d)
}

// QuerySpecialdoctors queries the specialdoctors edge of the Department.
func (d *Department) QuerySpecialdoctors() *SpecialdoctorQuery {
	return (&DepartmentClient{config: d.config}).QuerySpecialdoctors(d)
}

// Update returns a builder for updating this Department.
// Note that, you need to call Department.Unwrap() before calling this method, if this Department
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Department) Update() *DepartmentUpdateOne {
	return (&DepartmentClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Department) Unwrap() *Department {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Department is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Department) String() string {
	var builder strings.Builder
	builder.WriteString("Department(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", Detail=")
	builder.WriteString(d.Detail)
	builder.WriteString(", Name=")
	builder.WriteString(d.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Departments is a parsable slice of Department.
type Departments []*Department

func (d Departments) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
