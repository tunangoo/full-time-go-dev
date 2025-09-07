package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tunangoo/full-time-go-dev/internal/model"
	"github.com/tunangoo/full-time-go-dev/internal/service"
)

type AuthHandler struct {
	userService service.UserService
	authService service.AuthService
}

func NewAuthHandler(
	userService service.UserService,
	authService service.AuthService,
) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		authService: authService,
	}
}

func (h *AuthHandler) RegisterRoutes(router gin.IRouter) {
	g := router.Group("/auth")
	g.POST("/login", h.Login)
	g.POST("/register", h.Register)
}

// Register godoc
// @Summary Register user
// @Description Register user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.CreateUserRequest true "Register user request"
// @Success 201 {object} gin.H{message=string,user=model.User}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}

	user, err := h.userService.CreateUser(c.Request.Context(), &req)
	if err != nil {
		log.Println("Error creating user:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

// Login godoc
// @Summary Login user
// @Description Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "Login user request"
// @Success 200 {object} gin.H{message=string,user=model.User}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}

	token, err := h.authService.Login(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "User logged in successfully",
		"access_token": token,
	})
}
