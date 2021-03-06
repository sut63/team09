// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team09/app/ent/extradoctor"
)

// Extradoctor is the model entity for the Extradoctor schema.
type Extradoctor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Specialname holds the value of the "specialname" field.
	Specialname string `json:"specialname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExtradoctorQuery when eager-loading is set.
	Edges ExtradoctorEdges `json:"edges"`
}

// ExtradoctorEdges holds the relations/edges for other nodes in the graph.
type ExtradoctorEdges struct {
	// Specialdoctors holds the value of the specialdoctors edge.
	Specialdoctors []*Specialdoctor
	// Offices holds the value of the offices edge.
	Offices []*Office
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SpecialdoctorsOrErr returns the Specialdoctors value or an error if the edge
// was not loaded in eager-loading.
func (e ExtradoctorEdges) SpecialdoctorsOrErr() ([]*Specialdoctor, error) {
	if e.loadedTypes[0] {
		return e.Specialdoctors, nil
	}
	return nil, &NotLoadedError{edge: "specialdoctors"}
}

// OfficesOrErr returns the Offices value or an error if the edge
// was not loaded in eager-loading.
func (e ExtradoctorEdges) OfficesOrErr() ([]*Office, error) {
	if e.loadedTypes[1] {
		return e.Offices, nil
	}
	return nil, &NotLoadedError{edge: "offices"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Extradoctor) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // specialname
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Extradoctor fields.
func (e *Extradoctor) assignValues(values ...interface{}) error {
	if m, n := len(values), len(extradoctor.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	e.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field specialname", values[0])
	} else if value.Valid {
		e.Specialname = value.String
	}
	return nil
}

// QuerySpecialdoctors queries the specialdoctors edge of the Extradoctor.
func (e *Extradoctor) QuerySpecialdoctors() *SpecialdoctorQuery {
	return (&ExtradoctorClient{config: e.config}).QuerySpecialdoctors(e)
}

// QueryOffices queries the offices edge of the Extradoctor.
func (e *Extradoctor) QueryOffices() *OfficeQuery {
	return (&ExtradoctorClient{config: e.config}).QueryOffices(e)
}

// Update returns a builder for updating this Extradoctor.
// Note that, you need to call Extradoctor.Unwrap() before calling this method, if this Extradoctor
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Extradoctor) Update() *ExtradoctorUpdateOne {
	return (&ExtradoctorClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (e *Extradoctor) Unwrap() *Extradoctor {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Extradoctor is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Extradoctor) String() string {
	var builder strings.Builder
	builder.WriteString("Extradoctor(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", specialname=")
	builder.WriteString(e.Specialname)
	builder.WriteByte(')')
	return builder.String()
}

// Extradoctors is a parsable slice of Extradoctor.
type Extradoctors []*Extradoctor

func (e Extradoctors) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
