package schema

import (
	"unicode/utf8"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/sijiaoh/go-godot-template/api_server/validators"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(12 * utf8.UTFMax).
			Validate(validators.ValidateStringLength(0, 12)),
		field.String("token").
			NotEmpty().
			Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
