// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kaews/app/ent/patient"
)

// Patient is the model entity for the Patient schema.
type Patient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatientQuery when eager-loading is set.
	Edges PatientEdges `json:"edges"`
}

// PatientEdges holds the relations/edges for other nodes in the graph.
type PatientEdges struct {
	// Medicalfiles holds the value of the medicalfiles edge.
	Medicalfiles []*Medicalfile
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MedicalfilesOrErr returns the Medicalfiles value or an error if the edge
// was not loaded in eager-loading.
func (e PatientEdges) MedicalfilesOrErr() ([]*Medicalfile, error) {
	if e.loadedTypes[0] {
		return e.Medicalfiles, nil
	}
	return nil, &NotLoadedError{edge: "medicalfiles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Patient) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullInt64{},  // age
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Patient fields.
func (pa *Patient) assignValues(values ...interface{}) error {
	if m, n := len(values), len(patient.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		pa.Name = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[1])
	} else if value.Valid {
		pa.Age = int(value.Int64)
	}
	return nil
}

// QueryMedicalfiles queries the medicalfiles edge of the Patient.
func (pa *Patient) QueryMedicalfiles() *MedicalfileQuery {
	return (&PatientClient{config: pa.config}).QueryMedicalfiles(pa)
}

// Update returns a builder for updating this Patient.
// Note that, you need to call Patient.Unwrap() before calling this method, if this Patient
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Patient) Update() *PatientUpdateOne {
	return (&PatientClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Patient) Unwrap() *Patient {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Patient is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Patient) String() string {
	var builder strings.Builder
	builder.WriteString("Patient(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", name=")
	builder.WriteString(pa.Name)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", pa.Age))
	builder.WriteByte(')')
	return builder.String()
}

// Patients is a parsable slice of Patient.
type Patients []*Patient

func (pa Patients) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
