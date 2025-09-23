package handler

import (
	"log"
	"net/http"

	"github.com/tunangoo/full-time-go-dev/internal/service"

	"github.com/tunangoo/full-time-go-dev/internal/model"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(
	userService service.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterRoutes(router gin.IRouter, authMiddleware gin.HandlerFunc) {
	g := router.Group("/user", authMiddleware)
	g.GET("", h.GetUser)
	g.DELETE("", h.DeleteUser)
	g.PUT("", h.UpdateUser)
}

// GetUser godoc
// @Summary Get user
// @Description Get user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} gin.H{user=model.User}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/user [get]
// @Security BearerAuth
func (h *UserHandler) GetUser(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	user, err := h.userService.GetUserByID(c.Request.Context(), user.ID)
	if err != nil {
		log.Println("Error getting user by ID:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/user [delete]
// @Security BearerAuth
func (h *UserHandler) DeleteUser(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	err := h.userService.DeleteUser(c.Request.Context(), user.ID)
	if err != nil {
		log.Println("Error deleting user:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user
// @Tags user
// @Accept json
// @Produce json
// @Param request body model.UpdateUserRequest true "Update user request"
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/user [put]
// @Security BearerAuth
func (h *UserHandler) UpdateUser(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}

	err := h.userService.UpdateUser(c.Request.Context(), user.ID, &req)
	if err != nil {
		log.Println("Error updating user:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}
