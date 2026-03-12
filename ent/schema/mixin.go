// Package schema
package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

type SoftDeleteMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").
			Unique().
			Immutable(),
		field.
			Time("created_at").
			Default(time.Now),
		field.
			Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.
			Time("deleted_at").
			Optional().
			Nillable(),
	}
}

func (SoftDeleteMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.TraverseFunc(func(ctx context.Context, q ent.Query) error {
			if skip, _ := ctx.Value(SoftDeleteKey{}).(bool); skip {
				return nil
			}
			type SoftDeleter interface {
				WhereP(...func(*sql.Selector))
			}
			if d, ok := q.(SoftDeleter); ok {
				d.WhereP(func(s *sql.Selector) {
					s.Where(sql.IsNull(s.C("deleted_at")))
				})
			}
			return nil
		}),
	}
}

type SoftDeleteKey struct{}

func (SoftDeleteMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		// 将 Delete 操作转换为更新 deleted_at
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				if m.Op() != ent.OpDelete && m.Op() != ent.OpDeleteOne {
					return next.Mutate(ctx, m)
				}

				// 跳过软删除
				if skip, _ := ctx.Value(SoftDeleteKey{}).(bool); skip {
					return next.Mutate(ctx, m)
				}

				// 将 Delete 转为 Update deleted_at
				type SoftDeleteMutator interface {
					SetOp(ent.Op)
					SetDeletedAt(time.Time)
					WhereP(...func(*sql.Selector))
				}
				mx, ok := m.(SoftDeleteMutator)
				if !ok {
					return next.Mutate(ctx, m)
				}
				mx.SetOp(ent.OpUpdate)
				mx.SetDeletedAt(time.Now())
				return next.Mutate(ctx, m)
			})
		},
	}
}
