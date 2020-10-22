// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// DentistsColumns holds the columns for the "dentists" table.
	DentistsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// DentistsTable holds the schema information for the "dentists" table.
	DentistsTable = &schema.Table{
		Name:        "dentists",
		Columns:     DentistsColumns,
		PrimaryKey:  []*schema.Column{DentistsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// EmployeesColumns holds the columns for the "employees" table.
	EmployeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// EmployeesTable holds the schema information for the "employees" table.
	EmployeesTable = &schema.Table{
		Name:        "employees",
		Columns:     EmployeesColumns,
		PrimaryKey:  []*schema.Column{EmployeesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MedicalfilesColumns holds the columns for the "medicalfiles" table.
	MedicalfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "detail", Type: field.TypeString},
		{Name: "added_time", Type: field.TypeTime},
		{Name: "dentist_id", Type: field.TypeInt, Nullable: true},
		{Name: "employee_id", Type: field.TypeInt, Nullable: true},
		{Name: "patient_id", Type: field.TypeInt, Nullable: true},
	}
	// MedicalfilesTable holds the schema information for the "medicalfiles" table.
	MedicalfilesTable = &schema.Table{
		Name:       "medicalfiles",
		Columns:    MedicalfilesColumns,
		PrimaryKey: []*schema.Column{MedicalfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "medicalfiles_dentists_medicalfiles",
				Columns: []*schema.Column{MedicalfilesColumns[3]},

				RefColumns: []*schema.Column{DentistsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "medicalfiles_employees_medicalfiles",
				Columns: []*schema.Column{MedicalfilesColumns[4]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "medicalfiles_patients_medicalfiles",
				Columns: []*schema.Column{MedicalfilesColumns[5]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PatientsColumns holds the columns for the "patients" table.
	PatientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
	}
	// PatientsTable holds the schema information for the "patients" table.
	PatientsTable = &schema.Table{
		Name:        "patients",
		Columns:     PatientsColumns,
		PrimaryKey:  []*schema.Column{PatientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DentistsTable,
		EmployeesTable,
		MedicalfilesTable,
		PatientsTable,
	}
)

func init() {
	MedicalfilesTable.ForeignKeys[0].RefTable = DentistsTable
	MedicalfilesTable.ForeignKeys[1].RefTable = EmployeesTable
	MedicalfilesTable.ForeignKeys[2].RefTable = PatientsTable
}
