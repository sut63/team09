// Code generated by entc, DO NOT EDIT.

package workingtime

const (
	// Label holds the string label denoting the workingtime type in the database.
	Label = "workingtime"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAddedTime holds the string denoting the added_time field in the database.
	FieldAddedTime = "added_time"

	// EdgeOffices holds the string denoting the offices edge name in mutations.
	EdgeOffices = "offices"

	// Table holds the table name of the workingtime in the database.
	Table = "workingtimes"
	// OfficesTable is the table the holds the offices relation/edge.
	OfficesTable = "offices"
	// OfficesInverseTable is the table name for the Office entity.
	// It exists in this package in order to avoid circular dependency with the "office" package.
	OfficesInverseTable = "offices"
	// OfficesColumn is the table column denoting the offices relation/edge.
	OfficesColumn = "workingtime_id"
)

// Columns holds all SQL columns for workingtime fields.
var Columns = []string{
	FieldID,
	FieldAddedTime,
}
