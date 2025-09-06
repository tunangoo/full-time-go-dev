package service

import (
	"context"
	"log"

	"github.com/tunangoo/full-time-go-dev/internal/model"
	"github.com/tunangoo/full-time-go-dev/internal/repository"
	"github.com/tunangoo/full-time-go-dev/internal/util"
)

type UserService interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, user *model.CreateUserRequest) (*model.User, error)
	ListAllUser(ctx context.Context) ([]*model.User, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, id int, req *model.UpdateUserRequest) error
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

func (s *userService) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		log.Println("Error getting user by ID:", err)
		return nil, err
	}
	return user, nil
}

func (s *userService) CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
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
		log.Println("Error creating user:", err)
		return nil, err
	}

	return user, nil
}

func (s *userService) ListAllUser(ctx context.Context) ([]*model.User, error) {
	users, err := s.userRepository.ListUsers(ctx)
	if err != nil {
		log.Println("Error listing all users:", err)
		return nil, err
	}
	return users, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	err := s.userRepository.DeleteUser(ctx, id)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}
	return nil
}

func (s *userService) UpdateUser(ctx context.Context, id int, req *model.UpdateUserRequest) error {
	user, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		log.Println("Error getting user by ID:", err)
		return err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email

	err = s.userRepository.UpdateUser(ctx, user)
	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}
	return nil
}
