package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Withdrawals holds the schema definition for the Withdrawals entity.
type Withdrawals struct {
	ent.Schema
}

// Fields of the Withdrawals.
func (Withdrawals) Fields() []ent.Field {
	return []ent.Field{
		field.String("order").
			Unique(),
		field.Float("sum"),
		field.Time("timestamp").SchemaType(map[string]string{
			dialect.Postgres: "timestamptz",
		}),
	}
}

// Edges of the Withdrawals.
func (Withdrawals) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("withdrawals").
			Unique(),
	}
}
