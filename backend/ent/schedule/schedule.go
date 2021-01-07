// Code generated by entc, DO NOT EDIT.

package schedule

const (
	// Label holds the string label denoting the schedule type in the database.
	Label = "schedule"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldActivity holds the string denoting the activity field in the database.
	FieldActivity = "activity"
	// FieldAddedTime holds the string denoting the added_time field in the database.
	FieldAddedTime = "added_time"

	// EdgeDocter holds the string denoting the docter edge name in mutations.
	EdgeDocter = "docter"
	// EdgeDepartment holds the string denoting the department edge name in mutations.
	EdgeDepartment = "department"
	// EdgeOffice holds the string denoting the office edge name in mutations.
	EdgeOffice = "office"

	// Table holds the table name of the schedule in the database.
	Table = "schedules"
	// DocterTable is the table the holds the docter relation/edge.
	DocterTable = "schedules"
	// DocterInverseTable is the table name for the Doctor entity.
	// It exists in this package in order to avoid circular dependency with the "doctor" package.
	DocterInverseTable = "doctors"
	// DocterColumn is the table column denoting the docter relation/edge.
	DocterColumn = "schedule_id"
	// DepartmentTable is the table the holds the department relation/edge.
	DepartmentTable = "schedules"
	// DepartmentInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentInverseTable = "departments"
	// DepartmentColumn is the table column denoting the department relation/edge.
	DepartmentColumn = "department_id"
	// OfficeTable is the table the holds the office relation/edge.
	OfficeTable = "schedules"
	// OfficeInverseTable is the table name for the Office entity.
	// It exists in this package in order to avoid circular dependency with the "office" package.
	OfficeInverseTable = "offices"
	// OfficeColumn is the table column denoting the office relation/edge.
	OfficeColumn = "office_id"
)

// Columns holds all SQL columns for schedule fields.
var Columns = []string{
	FieldID,
	FieldActivity,
	FieldAddedTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Schedule type.
var ForeignKeys = []string{
	"department_id",
	"schedule_id",
	"office_id",
}

var (
	// ActivityValidator is a validator for the "activity" field. It is called by the builders before save.
	ActivityValidator func(string) error
)
