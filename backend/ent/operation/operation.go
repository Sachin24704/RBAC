// Code generated by ent, DO NOT EDIT.

package operation

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the operation type in the database.
	Label = "operation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldModelID holds the string denoting the model_id field in the database.
	FieldModelID = "model_id"
	// EdgeModel holds the string denoting the model edge name in mutations.
	EdgeModel = "model"
	// EdgeRoleOperations holds the string denoting the role_operations edge name in mutations.
	EdgeRoleOperations = "role_operations"
	// Table holds the table name of the operation in the database.
	Table = "operations"
	// ModelTable is the table that holds the model relation/edge.
	ModelTable = "operations"
	// ModelInverseTable is the table name for the Model entity.
	// It exists in this package in order to avoid circular dependency with the "model" package.
	ModelInverseTable = "models"
	// ModelColumn is the table column denoting the model relation/edge.
	ModelColumn = "model_id"
	// RoleOperationsTable is the table that holds the role_operations relation/edge.
	RoleOperationsTable = "role_operations"
	// RoleOperationsInverseTable is the table name for the RoleOperation entity.
	// It exists in this package in order to avoid circular dependency with the "roleoperation" package.
	RoleOperationsInverseTable = "role_operations"
	// RoleOperationsColumn is the table column denoting the role_operations relation/edge.
	RoleOperationsColumn = "operation_id"
)

// Columns holds all SQL columns for operation fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldModelID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the Operation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByModelID orders the results by the model_id field.
func ByModelID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelID, opts...).ToFunc()
}

// ByModelField orders the results by model field.
func ByModelField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newModelStep(), sql.OrderByField(field, opts...))
	}
}

// ByRoleOperationsCount orders the results by role_operations count.
func ByRoleOperationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoleOperationsStep(), opts...)
	}
}

// ByRoleOperations orders the results by role_operations terms.
func ByRoleOperations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoleOperationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newModelStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ModelInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ModelTable, ModelColumn),
	)
}
func newRoleOperationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoleOperationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RoleOperationsTable, RoleOperationsColumn),
	)
}
