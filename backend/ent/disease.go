// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team09/app/ent/disease"
)

// Disease is the model entity for the Disease schema.
type Disease struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Disease holds the value of the "disease" field.
	Disease string `json:"disease,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DiseaseQuery when eager-loading is set.
	Edges DiseaseEdges `json:"edges"`
}

// DiseaseEdges holds the relations/edges for other nodes in the graph.
type DiseaseEdges struct {
	// Doctors holds the value of the doctors edge.
	Doctors []*Doctor
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DoctorsOrErr returns the Doctors value or an error if the edge
// was not loaded in eager-loading.
func (e DiseaseEdges) DoctorsOrErr() ([]*Doctor, error) {
	if e.loadedTypes[0] {
		return e.Doctors, nil
	}
	return nil, &NotLoadedError{edge: "doctors"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Disease) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // disease
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Disease fields.
func (d *Disease) assignValues(values ...interface{}) error {
	if m, n := len(values), len(disease.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field disease", values[0])
	} else if value.Valid {
		d.Disease = value.String
	}
	return nil
}

// QueryDoctors queries the doctors edge of the Disease.
func (d *Disease) QueryDoctors() *DoctorQuery {
	return (&DiseaseClient{config: d.config}).QueryDoctors(d)
}

// Update returns a builder for updating this Disease.
// Note that, you need to call Disease.Unwrap() before calling this method, if this Disease
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Disease) Update() *DiseaseUpdateOne {
	return (&DiseaseClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Disease) Unwrap() *Disease {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Disease is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Disease) String() string {
	var builder strings.Builder
	builder.WriteString("Disease(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", disease=")
	builder.WriteString(d.Disease)
	builder.WriteByte(')')
	return builder.String()
}

// Diseases is a parsable slice of Disease.
type Diseases []*Disease

func (d Diseases) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
