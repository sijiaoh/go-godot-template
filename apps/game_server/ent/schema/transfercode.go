package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TransferCode holds the schema definition for the TransferCode entity.
type TransferCode struct {
	ent.Schema
}

// Fields of the TransferCode.
func (TransferCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Unique(),
	}
}

// Edges of the TransferCode.
func (TransferCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("transfer_code").Required().Unique(),
	}
}
