package repository

import (
	"context"

	"github.com/tunangoo/full-time-go-dev/internal/model"

	"github.com/uptrace/bun"
)

type UserRepository interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	ListUsers(ctx context.Context) ([]*model.User, error)
}

type userRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUser(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	return &user, err
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r *userRepository) ListUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := r.db.NewSelect().Model(&users).Scan(ctx)
	return users, err
}
