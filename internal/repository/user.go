package repository

import (
	"context"
	"time"

	"github.com/tunangoo/full-time-go-dev/internal/model"

	"github.com/uptrace/bun"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	ListUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
	DeleteUser(ctx context.Context, id int64) error
	UpdateUser(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	return &user, err
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.NewSelect().Model(&user).Where("email = ?", email).Scan(ctx)
	return &user, err
}

func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	_, err := r.db.NewDelete().Model(&model.User{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r *userRepository) ListUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := r.db.NewSelect().Model(&users).Scan(ctx)
	return users, err
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	user.UpdatedAt = time.Now()
	_, err := r.db.NewUpdate().Model(user).Where("id = ?", user.ID).Exec(ctx)
	return err
}
