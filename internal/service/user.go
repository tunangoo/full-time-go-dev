package service

import (
	"context"

	"github.com/tunangoo/full-time-go-dev/internal/model"
	"github.com/tunangoo/full-time-go-dev/internal/repository"
)

type UserService interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	ListAllUser(ctx context.Context) ([]*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(
	userRepository repository.UserRepository,
) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.userRepository.GetUser(ctx, id)
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	err := s.userRepository.CreateUser(ctx, user)
	return err
}

func (s *userService) ListAllUser(ctx context.Context) ([]*model.User, error) {
	return s.userRepository.ListUsers(ctx)
}
