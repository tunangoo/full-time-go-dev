package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userHandler *UserHandler
}

func NewHandler(
	userHandler *UserHandler,
) *Handler {
	return &Handler{
		userHandler: userHandler,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	apiv1 := router.Group("/api/v1")
	h.userHandler.RegisterRoutes(apiv1)
}
