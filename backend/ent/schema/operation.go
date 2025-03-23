package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Operation holds the schema definition for the Operation entity.
type Operation struct {
	ent.Schema
}

// Fields of the Operation.
func (Operation) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("model_id"),
	}
}

// Edges of the Operation.
func (Operation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("model", Model.Type).Ref("operations").Field("model_id").Unique().Required(),
		edge.To("role_operations", RoleOperation.Type).Comment("operations which are assigned to roles"),
	}
}
