package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tunangoo/full-time-go-dev/internal/middleware"
)

type Handler struct {
	userHandler   *UserHandler
	hotelHandler  *HotelHandler
	roomHandler   *RoomHandler
	authHandler   *AuthHandler
	jwtMiddleware *middleware.JWTMiddleware
}

func NewHandler(
	userHandler *UserHandler,
	hotelHandler *HotelHandler,
	roomHandler *RoomHandler,
	authHandler *AuthHandler,
	jwtMiddleware *middleware.JWTMiddleware,
) *Handler {
	return &Handler{
		userHandler:   userHandler,
		hotelHandler:  hotelHandler,
		roomHandler:   roomHandler,
		authHandler:   authHandler,
		jwtMiddleware: jwtMiddleware,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	apiv1 := router.Group("/api/v1")
	h.userHandler.RegisterRoutes(apiv1, h.jwtMiddleware.Handle())
	h.hotelHandler.RegisterRoutes(apiv1, h.jwtMiddleware.Handle())
	h.roomHandler.RegisterRoutes(apiv1, h.jwtMiddleware.Handle())
	h.authHandler.RegisterRoutes(apiv1)
}
