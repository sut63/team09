// Code generated by entc, DO NOT EDIT.

package office

const (
	// Label holds the string label denoting the office type in the database.
	Label = "office"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOfficename holds the string denoting the officename field in the database.
	FieldOfficename = "officename"
	// FieldRoomnumber holds the string denoting the roomnumber field in the database.
	FieldRoomnumber = "roomnumber"
	// FieldDoctoridcard holds the string denoting the doctoridcard field in the database.
	FieldDoctoridcard = "doctoridcard"
	// FieldAddedTime1 holds the string denoting the added_time1 field in the database.
	FieldAddedTime1 = "added_time1"
	// FieldAddedTime2 holds the string denoting the added_time2 field in the database.
	FieldAddedTime2 = "added_time2"

	// EdgeDoctor holds the string denoting the doctor edge name in mutations.
	EdgeDoctor = "doctor"
	// EdgeDepartment holds the string denoting the department edge name in mutations.
	EdgeDepartment = "department"
	// EdgeExtradoctor holds the string denoting the extradoctor edge name in mutations.
	EdgeExtradoctor = "extradoctor"
	// EdgeSchedules holds the string denoting the schedules edge name in mutations.
	EdgeSchedules = "schedules"

	// Table holds the table name of the office in the database.
	Table = "offices"
	// DoctorTable is the table the holds the doctor relation/edge.
	DoctorTable = "offices"
	// DoctorInverseTable is the table name for the Doctor entity.
	// It exists in this package in order to avoid circular dependency with the "doctor" package.
	DoctorInverseTable = "doctors"
	// DoctorColumn is the table column denoting the doctor relation/edge.
	DoctorColumn = "doctor_id"
	// DepartmentTable is the table the holds the department relation/edge.
	DepartmentTable = "offices"
	// DepartmentInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentInverseTable = "departments"
	// DepartmentColumn is the table column denoting the department relation/edge.
	DepartmentColumn = "department_id"
	// ExtradoctorTable is the table the holds the extradoctor relation/edge.
	ExtradoctorTable = "offices"
	// ExtradoctorInverseTable is the table name for the Extradoctor entity.
	// It exists in this package in order to avoid circular dependency with the "extradoctor" package.
	ExtradoctorInverseTable = "extradoctors"
	// ExtradoctorColumn is the table column denoting the extradoctor relation/edge.
	ExtradoctorColumn = "extradoctor_id"
	// SchedulesTable is the table the holds the schedules relation/edge.
	SchedulesTable = "schedules"
	// SchedulesInverseTable is the table name for the Schedule entity.
	// It exists in this package in order to avoid circular dependency with the "schedule" package.
	SchedulesInverseTable = "schedules"
	// SchedulesColumn is the table column denoting the schedules relation/edge.
	SchedulesColumn = "office_id"
)

// Columns holds all SQL columns for office fields.
var Columns = []string{
	FieldID,
	FieldOfficename,
	FieldRoomnumber,
	FieldDoctoridcard,
	FieldAddedTime1,
	FieldAddedTime2,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Office type.
var ForeignKeys = []string{
	"department_id",
	"doctor_id",
	"extradoctor_id",
}

var (
	// OfficenameValidator is a validator for the "officename" field. It is called by the builders before save.
	OfficenameValidator func(string) error
	// RoomnumberValidator is a validator for the "roomnumber" field. It is called by the builders before save.
	RoomnumberValidator func(string) error
	// DoctoridcardValidator is a validator for the "doctoridcard" field. It is called by the builders before save.
	DoctoridcardValidator func(string) error
)
