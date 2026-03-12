package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PostCategoryRelation 是 post_category_relations 中间表
type PostCategoryRelation struct {
	ent.Schema
}

func (PostCategoryRelation) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("post_id"),
		field.Uint("post_category_id"),
	}
}

func (PostCategoryRelation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("post", Post.Type).Field("post_id").Unique().Required(),
		edge.To("category", PostCategory.Type).Field("post_category_id").Unique().Required(),
	}
}

func (PostCategoryRelation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("post_id", "post_category_id"), // 复合主键
		entsql.Annotation{Table: "post_category_relations"},
	}
}
