package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RoleOperation holds the schema definition for the RoleOperation entity.
type RoleOperation struct {
	ent.Schema
}

// Fields of the RoleOperation.
func (RoleOperation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("role_id"),
		field.Int("operation_id"),
	}
}

// Edges of the RoleOperation.
func (RoleOperation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).Ref("role_operations").Field("role_id").Unique().Required(),
		edge.From("operations", Operation.Type).Ref("role_operations").Field("operation_id").Unique().Required(),
	}
}
