// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/operation"
	"backend/ent/role"
	"backend/ent/roleoperation"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// RoleOperation is the model entity for the RoleOperation schema.
type RoleOperation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RoleID holds the value of the "role_id" field.
	RoleID int `json:"role_id,omitempty"`
	// OperationID holds the value of the "operation_id" field.
	OperationID int `json:"operation_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleOperationQuery when eager-loading is set.
	Edges        RoleOperationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoleOperationEdges holds the relations/edges for other nodes in the graph.
type RoleOperationEdges struct {
	// Role holds the value of the role edge.
	Role *Role `json:"role,omitempty"`
	// Operations holds the value of the operations edge.
	Operations *Operation `json:"operations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoleOperationEdges) RoleOrErr() (*Role, error) {
	if e.Role != nil {
		return e.Role, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: role.Label}
	}
	return nil, &NotLoadedError{edge: "role"}
}

// OperationsOrErr returns the Operations value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoleOperationEdges) OperationsOrErr() (*Operation, error) {
	if e.Operations != nil {
		return e.Operations, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: operation.Label}
	}
	return nil, &NotLoadedError{edge: "operations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoleOperation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case roleoperation.FieldID, roleoperation.FieldRoleID, roleoperation.FieldOperationID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoleOperation fields.
func (ro *RoleOperation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case roleoperation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ro.ID = int(value.Int64)
		case roleoperation.FieldRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				ro.RoleID = int(value.Int64)
			}
		case roleoperation.FieldOperationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field operation_id", values[i])
			} else if value.Valid {
				ro.OperationID = int(value.Int64)
			}
		default:
			ro.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RoleOperation.
// This includes values selected through modifiers, order, etc.
func (ro *RoleOperation) Value(name string) (ent.Value, error) {
	return ro.selectValues.Get(name)
}

// QueryRole queries the "role" edge of the RoleOperation entity.
func (ro *RoleOperation) QueryRole() *RoleQuery {
	return NewRoleOperationClient(ro.config).QueryRole(ro)
}

// QueryOperations queries the "operations" edge of the RoleOperation entity.
func (ro *RoleOperation) QueryOperations() *OperationQuery {
	return NewRoleOperationClient(ro.config).QueryOperations(ro)
}

// Update returns a builder for updating this RoleOperation.
// Note that you need to call RoleOperation.Unwrap() before calling this method if this RoleOperation
// was returned from a transaction, and the transaction was committed or rolled back.
func (ro *RoleOperation) Update() *RoleOperationUpdateOne {
	return NewRoleOperationClient(ro.config).UpdateOne(ro)
}

// Unwrap unwraps the RoleOperation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ro *RoleOperation) Unwrap() *RoleOperation {
	_tx, ok := ro.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoleOperation is not a transactional entity")
	}
	ro.config.driver = _tx.drv
	return ro
}

// String implements the fmt.Stringer.
func (ro *RoleOperation) String() string {
	var builder strings.Builder
	builder.WriteString("RoleOperation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ro.ID))
	builder.WriteString("role_id=")
	builder.WriteString(fmt.Sprintf("%v", ro.RoleID))
	builder.WriteString(", ")
	builder.WriteString("operation_id=")
	builder.WriteString(fmt.Sprintf("%v", ro.OperationID))
	builder.WriteByte(')')
	return builder.String()
}

// RoleOperations is a parsable slice of RoleOperation.
type RoleOperations []*RoleOperation
