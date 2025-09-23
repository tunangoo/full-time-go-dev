package service

import (
	"context"
	"log"

	"github.com/tunangoo/full-time-go-dev/internal/config"
	"github.com/tunangoo/full-time-go-dev/internal/model"
	"github.com/tunangoo/full-time-go-dev/internal/repository"
	"github.com/tunangoo/full-time-go-dev/internal/util"
)

type AuthService interface {
	Login(ctx context.Context, req *model.LoginRequest) (string, error)
}

type authService struct {
	userRepository repository.UserRepository
	jwtProvider    config.TokenProvider
}

func NewAuthService(
	userRepository repository.UserRepository,
	jwtProvider config.TokenProvider,
) AuthService {
	return &authService{
		userRepository: userRepository,
		jwtProvider:    jwtProvider,
	}
}

func (s *authService) Login(ctx context.Context, req *model.LoginRequest) (string, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Println("Error getting user by email:", err)
		return "", err
	}

	if err := util.VerifyPassword(req.Password, user.Password); err != nil {
		log.Println("Error verifying password:", err)
		return "", err
	}

	// generate token
	token, err := s.jwtProvider.Generate(config.TokenPayload{UserID: user.ID}, 3600) // 1 hour
	if err != nil {
		log.Println("Error generating token:", err)
		return "", err
	}
	return token, nil
}
