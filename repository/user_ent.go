package repository

import (
	"context"
	"fmt"
	"strings"

	"blog-server/ent"
	"blog-server/ent/user"
	"blog-server/entity"
	usecase "blog-server/usecase/user"

	"github.com/google/uuid"
)

type userRepoEnt struct {
	client *ent.Client
}

func (r *userRepoEnt) Create(ctx context.Context, e *entity.User) (*entity.User, error) {
	u, err := r.client.User.Create().
		SetEmail(e.Email).
		SetPassword(e.Password).
		SetUsername(e.Username).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create user failed: %v", err)
	}

	user := &entity.User{
		User: *u,
	}

	return user, nil
}

// Delete implements [user.Repository].
func (r *userRepoEnt) Delete(ctx context.Context, uuid uuid.UUID) error {
	panic("unimplemented")
}

// Get implements [user.Repository].
func (r *userRepoEnt) Get(ctx context.Context, uuid uuid.UUID) (*entity.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.UUIDEQ(uuid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user failed: %v", err)
	}

	return &entity.User{User: *u}, nil
}

// GetByEmail implements [user.Repository].
func (r *userRepoEnt) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user failed: %v", err)
	}

	return &entity.User{User: *u}, nil
}

// Update implements [user.Repository].
func (r *userRepoEnt) Update(ctx context.Context, e *entity.User) (*entity.User, error) {
	query := r.client.User.UpdateOneID(e.ID)

	if strings.TrimSpace(e.Password) != "" {
		query.SetPassword(e.Password)
	}

	u, err := query.
		SetEmail(e.Email).
		SetNillableAvatar(e.Avatar).
		SetUsername(e.Username).
		SetRole(e.Role).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("update user failed: %v", err)
	}

	return &entity.User{User: *u}, nil
}

func NewUserRepoEnt(client *ent.Client) usecase.Repository {
	return &userRepoEnt{
		client: client,
	}
}
