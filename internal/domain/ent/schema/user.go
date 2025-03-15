package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("username").NotEmpty(),
		field.Int("age").Positive().Max(100),
		field.String("description").NotEmpty(),
		field.String("interests").NotEmpty(),
		field.Int("rating").Positive().Max(5),
		field.String("image_url").Nillable().Optional(),
		field.String("tags").Nillable().Optional(),
		field.String("password").MinLen(8).MaxLen(100),
		field.String("type").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
