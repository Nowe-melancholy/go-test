package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Middleware holds the schema definition for the Middleware entity.
type Middleware struct {
	ent.Schema
}

// Fields of the Middleware.
func (Middleware) Fields() []ent.Field {
	return []ent.Field{
		field.String("l_id").
			SchemaType(map[string]string{
				dialect.Postgres: "int",
			}),
		field.String("d_id").
			SchemaType(map[string]string{
				dialect.Postgres: "int",
			}),
		field.String("sys_id").
			SchemaType(map[string]string{
				dialect.Postgres: "int",
			}),
	}
}

// Edges of the Middleware.
func (Middleware) Edges() []ent.Edge {
	return nil
}
