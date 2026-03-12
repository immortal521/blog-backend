// Package user usecase
package user

import (
	"context"

	"blog-server/entity"

	"github.com/google/uuid"
)

type Reader interface {
	Get(ctx context.Context, uuid uuid.UUID) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type Writer interface {
	Create(ctx context.Context, e *entity.User) (*entity.User, error)
	Update(ctx context.Context, e *entity.User) (*entity.User, error)
	Delete(ctx context.Context, uuid uuid.UUID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {

}
