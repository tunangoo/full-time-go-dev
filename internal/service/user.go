package service

import (
	"context"

	"github.com/tunangoo/full-time-go-dev/internal/model"
	"github.com/tunangoo/full-time-go-dev/internal/repository"
	"github.com/tunangoo/full-time-go-dev/internal/util"
)

type UserService interface {
	GetUser(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, user *model.CreateUserRequest) (*model.User, error)
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

func (s *userService) GetUser(ctx context.Context, id int) (*model.User, error) {
	return s.userRepository.GetUser(ctx, id)
}

func (s *userService) CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
	}

	err = s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	err = s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) ListAllUser(ctx context.Context) ([]*model.User, error) {
	return s.userRepository.ListUsers(ctx)
}
