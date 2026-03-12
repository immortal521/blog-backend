package repository

import (
	"context"

	"blog-server/ent"
	"blog-server/ent/post"
	"blog-server/ent/user"
	"blog-server/entity"
	"blog-server/errs"
	usecase "blog-server/usecase/post"

	"entgo.io/ent/dialect/sql"
)

type postRepoEnt struct {
	client *ent.Client
}

func (r *postRepoEnt) Count(ctx context.Context) (int64, error) {
	count, err := r.client.Post.
		Query().
		Count(ctx)
	if err != nil {
		return 0, errs.New(errs.CodeDatabaseError, "database error", err)
	}
	return int64(count), nil
}

func (r *postRepoEnt) Get(ctx context.Context, id uint) (*entity.Post, error) {
	entPost, err := r.client.Post.
		Query().
		Select(post.FieldID,
			post.FieldTitle,
			post.FieldSummary,
			post.FieldCover,
			post.FieldReadTimeMinutes,
			post.FieldViewCount,
			post.FieldStatus,
			post.FieldPublishedAt).
		WithAuthor(func(q *ent.UserQuery) {
			q.Select(user.FieldUsername)
		}).
		Where(post.StatusEQ("published")).
		First(ctx)
	if err == nil {
		return &entity.Post{Post: *entPost}, nil
	}
	if ent.IsNotFound(err) {
		return nil, errs.New(errs.CodeResourceNotFound, "post not found", err)
	}
	return nil, errs.New(errs.CodeDatabaseError, "database error", err)
}

func (r *postRepoEnt) ListMeta(ctx context.Context) ([]*entity.Post, error) {
	entPosts, err := r.client.Post.
		Query().
		Select(
			post.FieldID,
			post.FieldUpdatedAt,
		).
		Where(post.StatusEQ("published")).
		All(ctx)
	if err != nil {
		return nil, errs.New(errs.CodeDatabaseError, "database error", err)
	}
	var posts []*entity.Post
	for i, p := range entPosts {
		posts[i] = &entity.Post{Post: *p}
	}
	return posts, nil
}

func (r *postRepoEnt) ListNoContent(ctx context.Context) ([]*entity.Post, error) {
	entPosts, err := r.client.Post.
		Query().
		Select(post.FieldID,
			post.FieldTitle,
			post.FieldSummary,
			post.FieldCover,
			post.FieldReadTimeMinutes,
			post.FieldViewCount,
			post.FieldStatus,
			post.FieldPublishedAt).
		WithAuthor(func(q *ent.UserQuery) {
			q.Select(user.FieldUsername)
		}).
		Where(post.StatusEQ("published")).
		Order(
			post.ByPublishedAt(
				sql.OrderDesc(),
			),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	posts := make([]*entity.Post, len(entPosts))
	for i, p := range entPosts {
		posts[i] = &entity.Post{Post: *p}
	}

	return posts, nil
}

func (r *postRepoEnt) UpdateViewCounts(ctx context.Context, updates map[uint]int64) error {
	panic("unimplemented")
}

func NewPostRepoEnt(client *ent.Client) usecase.Repository {
	return &postRepoEnt{
		client: client,
	}
}
