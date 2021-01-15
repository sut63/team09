// Code generated by entc, DO NOT EDIT.

package specialdoctor

const (
	// Label holds the string label denoting the specialdoctor type in the database.
	Label = "specialdoctor"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOther holds the string denoting the other field in the database.
	FieldOther = "other"

	// EdgeOffices holds the string denoting the offices edge name in mutations.
	EdgeOffices = "offices"
	// EdgeDoctor holds the string denoting the doctor edge name in mutations.
	EdgeDoctor = "doctor"
	// EdgeDepartment holds the string denoting the department edge name in mutations.
	EdgeDepartment = "department"
	// EdgeExtradoctor holds the string denoting the extradoctor edge name in mutations.
	EdgeExtradoctor = "extradoctor"

	// Table holds the table name of the specialdoctor in the database.
	Table = "specialdoctors"
	// OfficesTable is the table the holds the offices relation/edge.
	OfficesTable = "offices"
	// OfficesInverseTable is the table name for the Office entity.
	// It exists in this package in order to avoid circular dependency with the "office" package.
	OfficesInverseTable = "offices"
	// OfficesColumn is the table column denoting the offices relation/edge.
	OfficesColumn = "Specialdoctor_id"
	// DoctorTable is the table the holds the doctor relation/edge.
	DoctorTable = "specialdoctors"
	// DoctorInverseTable is the table name for the Doctor entity.
	// It exists in this package in order to avoid circular dependency with the "doctor" package.
	DoctorInverseTable = "doctors"
	// DoctorColumn is the table column denoting the doctor relation/edge.
	DoctorColumn = "doctor_id"
	// DepartmentTable is the table the holds the department relation/edge.
	DepartmentTable = "specialdoctors"
	// DepartmentInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentInverseTable = "departments"
	// DepartmentColumn is the table column denoting the department relation/edge.
	DepartmentColumn = "department_id"
	// ExtradoctorTable is the table the holds the extradoctor relation/edge.
	ExtradoctorTable = "specialdoctors"
	// ExtradoctorInverseTable is the table name for the Extradoctor entity.
	// It exists in this package in order to avoid circular dependency with the "extradoctor" package.
	ExtradoctorInverseTable = "extradoctors"
	// ExtradoctorColumn is the table column denoting the extradoctor relation/edge.
	ExtradoctorColumn = "extradoctor_id"
)

// Columns holds all SQL columns for specialdoctor fields.
var Columns = []string{
	FieldID,
	FieldOther,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Specialdoctor type.
var ForeignKeys = []string{
	"department_id",
	"doctor_id",
	"extradoctor_id",
}

var (
	// OtherValidator is a validator for the "other" field. It is called by the builders before save.
	OtherValidator func(string) error
)