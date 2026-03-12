package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LinkCategory holds the schema definition for the LinkCategory entity.
type LinkCategory struct {
	ent.Schema
}

func (LinkCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the LinkCategory.
func (LinkCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(20).
			Unique(),

		field.Int("sort_order").
			Default(0),
	}
}

// Edges of the LinkCategory.
func (LinkCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("links", Link.Type),
	}
}
