package handler

import (
	"log"
	"net/http"
	"strconv"

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

func (h *UserHandler) RegisterRoutes(router gin.IRouter) {
	g := router.Group("/user")
	g.GET("/all", h.ListAllUser)
	g.POST("/create", h.CreateUser)
	g.GET("/:id", h.GetUser)
	g.DELETE("/:id", h.DeleteUser)
	g.PUT("/:id", h.UpdateUser)
}

// ListAllUser godoc
// @Summary List all user
// @Description List all user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} gin.H{total=int,users=[]model.User}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/user/all [get]
func (h *UserHandler) ListAllUser(c *gin.Context) {
	resp, err := h.userService.ListAllUser(c.Request.Context())
	if err != nil {
		log.Println("Error listing all users:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": resp,
		"total": len(resp),
	})
}

// CreateUser godoc
// @Summary Create user
// @Description Create user
// @Tags user
// @Accept json
// @Produce json
// @Param request body model.CreateUserRequest true "Create user request"
// @Success 201 {object} gin.H{message=string,user=model.User}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/user/create [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
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
		"message": "User created successfully",
		"user":    user,
	})
}

// GetUser godoc
// @Summary Get user
// @Description Get user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} gin.H{user=model.User}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/user/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("Error converting ID to int:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}
	user, err := h.userService.GetUserByID(c.Request.Context(), id)
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
// @Param id path int true "User ID"
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/user/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("Error converting ID to int:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}

	err = h.userService.DeleteUser(c.Request.Context(), id)
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
// @Param id path int true "User ID"
// @Param request body model.UpdateUserRequest true "Update user request"
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/user/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("Error converting ID to int:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}

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

	err = h.userService.UpdateUser(c.Request.Context(), id, &req)
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
