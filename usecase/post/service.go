package post

import (
	"context"

	"blog-server/entity"
)

type Service struct {
	repo Repository
}

func (s *Service) GetPost(ctx context.Context, id uint) (*entity.Post, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) ListPosts(ctx context.Context) ([]*entity.Post, error) {
	panic("unimplemented")
}

func (s *Service) ListPostsMeta(ctx context.Context) ([]*entity.Post, error) {
	return s.repo.ListMeta(ctx)
}

func (s *Service) ListPostsNoContent(ctx context.Context) ([]*entity.Post, error) {
	return s.repo.ListNoContent(ctx)
}

func NewService(repo Repository) UseCase {
	return &Service{
		repo: repo,
	}
}
