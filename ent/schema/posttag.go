package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PostTag holds the schema definition for the PostTag entity.
type PostTag struct {
	ent.Schema
}

func (PostTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the PostTag.
func (PostTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(100).
			Unique(),

		field.String("slug").
			MaxLen(100).
			Unique(),
	}
}

// Edges of the PostTag.
func (PostTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posts", Post.Type).
			Ref("tags").
			Through("post_tag_relations", PostTagRelation.Type),
	}
}
