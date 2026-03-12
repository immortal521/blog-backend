package post

import (
	"context"

	"blog-server/entity"
)

type Reader interface {
	Get(ctx context.Context, id uint) (*entity.Post, error)
	ListNoContent(ctx context.Context) ([]*entity.Post, error)
	ListMeta(ctx context.Context) ([]*entity.Post, error)
	Count(ctx context.Context) (int64, error)
}

type Writer interface {
	UpdateViewCounts(ctx context.Context, updates map[uint]int64) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	ListPosts(ctx context.Context) ([]*entity.Post, error)
	ListPostsNoContent(ctx context.Context) ([]*entity.Post, error)
	GetPost(ctx context.Context, id uint) (*entity.Post, error)
	ListPostsMeta(ctx context.Context) ([]*entity.Post, error)
}
