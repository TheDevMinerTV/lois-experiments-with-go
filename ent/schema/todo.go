package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("item"),
		field.Bool("done").Default(false),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
