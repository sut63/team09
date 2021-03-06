// Code generated by entc, DO NOT EDIT.

package detail

const (
	// Label holds the string label denoting the detail type in the database.
	Label = "detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldExplain holds the string denoting the explain field in the database.
	FieldExplain = "explain"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"

	// EdgeCourse holds the string denoting the course edge name in mutations.
	EdgeCourse = "course"
	// EdgeMission holds the string denoting the mission edge name in mutations.
	EdgeMission = "mission"
	// EdgeDepartment holds the string denoting the department edge name in mutations.
	EdgeDepartment = "department"

	// Table holds the table name of the detail in the database.
	Table = "details"
	// CourseTable is the table the holds the course relation/edge.
	CourseTable = "details"
	// CourseInverseTable is the table name for the Course entity.
	// It exists in this package in order to avoid circular dependency with the "course" package.
	CourseInverseTable = "courses"
	// CourseColumn is the table column denoting the course relation/edge.
	CourseColumn = "course_id"
	// MissionTable is the table the holds the mission relation/edge.
	MissionTable = "details"
	// MissionInverseTable is the table name for the Mission entity.
	// It exists in this package in order to avoid circular dependency with the "mission" package.
	MissionInverseTable = "missions"
	// MissionColumn is the table column denoting the mission relation/edge.
	MissionColumn = "mission_id"
	// DepartmentTable is the table the holds the department relation/edge.
	DepartmentTable = "details"
	// DepartmentInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentInverseTable = "departments"
	// DepartmentColumn is the table column denoting the department relation/edge.
	DepartmentColumn = "department_id"
)

// Columns holds all SQL columns for detail fields.
var Columns = []string{
	FieldID,
	FieldExplain,
	FieldPhone,
	FieldEmail,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Detail type.
var ForeignKeys = []string{
	"course_id",
	"department_id",
	"mission_id",
}

var (
	// ExplainValidator is a validator for the "explain" field. It is called by the builders before save.
	ExplainValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
)
