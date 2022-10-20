package query

// sql类型
const (
	// UnknownType is the zero value for a Type
	UnknownType Type = iota
	// Select represents a SELECT query
	Select
	// Update represents an UPDATE query
	Update
	// Insert represents an INSERT query
	Insert
	// Delete represents a DELETE query
	Delete
)

// TypeString is a string slice with the names of all types in order
var TypeString = []string{
	"UnknownType",
	"Select",
	"Update",
	"Insert",
	"Delete",
}

const (
	// UnknownOperator is the zero value for an Operator
	UnknownOperator Operator = iota
	// Eq -> "="
	Eq
	// Ne -> "!="
	Ne
	// Gt -> ">"
	Gt
	// Lt -> "<"
	Lt
	// Gte -> ">="
	Gte
	// Lte -> "<="
	Lte
)

// OperatorString is a string slice with the names of all operators in order
var OperatorString = []string{
	"UnknownOperator",
	"Eq",
	"Ne",
	"Gt",
	"Lt",
	"Gte",
	"Lte",
}

// Type is the type of SQL query, e.g. SELECT/UPDATE
type Type int

// Operator is between operands in a condition
type Operator int

// Condition is a single boolean condition in a WHERE clause
type Condition struct {
	// Operand1 is the left hand side operand
	Operand1 string
	// Operand1IsField determines if Operand1 is a literal or a field name
	Operand1IsField bool
	// Operator is e.g. "=", ">"
	Operator Operator
	// Operand1 is the right hand side operand
	Operand2 string
	// Operand2IsField determines if Operand2 is a literal or a field name
	Operand2IsField bool
}

type Query struct {
	Type       Type
	TableName  string            // 表名
	Conditions []Condition       // 条件集合
	Updates    map[string]string // 更新字段的 key与value
	Inserts    [][]string
	Fields     []string          // Used for SELECT (i.e. SELECTed field names) and INSERT (INSERTEDed field names)
	Aliases    map[string]string // 别名
}
