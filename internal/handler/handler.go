package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userHandler  *UserHandler
	hotelHandler *HotelHandler
	roomHandler  *RoomHandler
}

func NewHandler(
	userHandler *UserHandler,
	hotelHandler *HotelHandler,
	roomHandler *RoomHandler,
) *Handler {
	return &Handler{
		userHandler:  userHandler,
		hotelHandler: hotelHandler,
		roomHandler:  roomHandler,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	apiv1 := router.Group("/api/v1")
	h.userHandler.RegisterRoutes(apiv1)
	h.hotelHandler.RegisterRoutes(apiv1)
	h.roomHandler.RegisterRoutes(apiv1)
}
