package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PostTagRelation struct {
	ent.Schema
}

func (PostTagRelation) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("post_id"),
		field.Uint("post_tag_id"),
	}
}

func (PostTagRelation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("post", Post.Type).Field("post_id").Unique().Required(),
		edge.To("tag", PostTag.Type).Field("post_tag_id").Unique().Required(),
	}
}

func (PostTagRelation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("post_id", "post_tag_id"),
		entsql.Annotation{Table: "post_tag_relations"},
	}
}
