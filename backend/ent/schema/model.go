package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Model holds the schema definition for the Model entity.
type Model struct {
	ent.Schema
}

// Fields of the Model.
func (Model) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the Model.
func (Model) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("operations", Operation.Type).Comment("model has many operations"),
	}
}
