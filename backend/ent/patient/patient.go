// Code generated by entc, DO NOT EDIT.

package patient

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"

	// EdgeMedicalfiles holds the string denoting the medicalfiles edge name in mutations.
	EdgeMedicalfiles = "medicalfiles"

	// Table holds the table name of the patient in the database.
	Table = "patients"
	// MedicalfilesTable is the table the holds the medicalfiles relation/edge.
	MedicalfilesTable = "medicalfiles"
	// MedicalfilesInverseTable is the table name for the Medicalfile entity.
	// It exists in this package in order to avoid circular dependency with the "medicalfile" package.
	MedicalfilesInverseTable = "medicalfiles"
	// MedicalfilesColumn is the table column denoting the medicalfiles relation/edge.
	MedicalfilesColumn = "patient_id"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAge,
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
)
