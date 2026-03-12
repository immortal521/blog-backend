package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PostCategory holds the schema definition for the PostCategory entity.
type PostCategory struct {
	ent.Schema
}

func (PostCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the PostCategory.
func (PostCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(100).
			Unique(),

		field.String("slug").
			MaxLen(100).
			Unique(),
	}
}

// Edges of the PostCategory.
func (PostCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posts", Post.Type).
			Ref("categories").
			Through("post_category_relations", PostCategoryRelation.Type),
	}
}
