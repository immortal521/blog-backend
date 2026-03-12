package schema

import (
	"entgo.io/ent"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{}
}
