// Code generated by ent, DO NOT EDIT.

package roleoperation

import (
	"backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldLTE(FieldID, id))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldEQ(FieldRoleID, v))
}

// OperationID applies equality check predicate on the "operation_id" field. It's identical to OperationIDEQ.
func OperationID(v int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldEQ(FieldOperationID, v))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldNotIn(FieldRoleID, vs...))
}

// OperationIDEQ applies the EQ predicate on the "operation_id" field.
func OperationIDEQ(v int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldEQ(FieldOperationID, v))
}

// OperationIDNEQ applies the NEQ predicate on the "operation_id" field.
func OperationIDNEQ(v int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldNEQ(FieldOperationID, v))
}

// OperationIDIn applies the In predicate on the "operation_id" field.
func OperationIDIn(vs ...int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldIn(FieldOperationID, vs...))
}

// OperationIDNotIn applies the NotIn predicate on the "operation_id" field.
func OperationIDNotIn(vs ...int) predicate.RoleOperation {
	return predicate.RoleOperation(sql.FieldNotIn(FieldOperationID, vs...))
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.RoleOperation {
	return predicate.RoleOperation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.RoleOperation {
	return predicate.RoleOperation(func(s *sql.Selector) {
		step := newRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOperations applies the HasEdge predicate on the "operations" edge.
func HasOperations() predicate.RoleOperation {
	return predicate.RoleOperation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OperationsTable, OperationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOperationsWith applies the HasEdge predicate on the "operations" edge with a given conditions (other predicates).
func HasOperationsWith(preds ...predicate.Operation) predicate.RoleOperation {
	return predicate.RoleOperation(func(s *sql.Selector) {
		step := newOperationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RoleOperation) predicate.RoleOperation {
	return predicate.RoleOperation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RoleOperation) predicate.RoleOperation {
	return predicate.RoleOperation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RoleOperation) predicate.RoleOperation {
	return predicate.RoleOperation(sql.NotPredicates(p))
}
