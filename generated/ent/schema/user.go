package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("login").
			NotEmpty().Unique(),
		field.String("password").
			NotEmpty(),
		field.Float("balance"),
		field.Float("withdraw"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type),
		edge.To("withdrawals", Withdrawals.Type),
	}
}
