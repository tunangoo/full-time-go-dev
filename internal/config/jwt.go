package config

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtConfig struct {
	JwtSecret string `mapstructure:"jwt_secret"`
}

type TokenPayload struct {
	UserID int64 `json:"user_id"`
}

type TokenProvider interface {
	Generate(payload TokenPayload, expiry int) (string, error)
	Validate(token string) (*TokenPayload, error)
}

type jwtProvider struct {
	secret string
}

func NewJwtProvider(secret string) TokenProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

func (j *jwtProvider) Generate(payload TokenPayload, expiry int) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
		payload.UserID,
	})

	token, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *jwtProvider) Validate(token string) (*TokenPayload, error) {
	t, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := t.Claims.(*myClaims)
	if !ok {
		return nil, err
	}

	return &TokenPayload{UserID: claims.UserID}, nil
}
