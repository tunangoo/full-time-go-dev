package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tunangoo/full-time-go-dev/hotel-reservation/internal/config"
	"github.com/tunangoo/full-time-go-dev/hotel-reservation/internal/model"
	"github.com/tunangoo/full-time-go-dev/hotel-reservation/internal/repository"
)

type JWTMiddleware struct {
	userRepo    repository.UserRepository
	jwtProvider config.TokenProvider
}

func NewJWTMiddleware(
	jwtProvider config.TokenProvider,
	userRepo repository.UserRepository,
) *JWTMiddleware {
	return &JWTMiddleware{
		userRepo:    userRepo,
		jwtProvider: jwtProvider,
	}
}

func (m *JWTMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract bearer token from Authorization header
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, model.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "missing Authorization header",
			})
			c.Abort()
			return
		}

		// Split the token to get the actual token
		token = token[7:]

		claims, err := m.jwtProvider.Validate(token)
		if err != nil {
			c.JSON(401, model.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "invalid token",
				Detail:  err.Error(),
			})
			c.Abort()
			return
		}

		user, err := m.userRepo.GetUserByID(c.Request.Context(), claims.UserID)
		if err != nil {
			c.JSON(500, model.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "internal server error",
				Detail:  err.Error(),
			})
			c.Abort()
			return
		}
		log.Printf("Middleware User: %+v", user)
		c.Set("user", user)
		c.Next()
	}
}
