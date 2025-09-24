package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ClientSession holds the schema definition for the ClientSession entity.
type ClientSession struct {
	ent.Schema
}

// Fields of the ClientSession.
func (ClientSession) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Unique(),
	}
}

// Edges of the ClientSession.
func (ClientSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("client_sessions").Required().Unique(),
	}
}
