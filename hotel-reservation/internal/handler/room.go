package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tunangoo/full-time-go-dev/internal/model"
	"github.com/tunangoo/full-time-go-dev/internal/service"
)

type RoomHandler struct {
	roomService service.RoomService
}

func NewRoomHandler(
	roomService service.RoomService,
) *RoomHandler {
	return &RoomHandler{roomService: roomService}
}

func (h *RoomHandler) RegisterRoutes(router gin.IRouter, authMiddleware gin.HandlerFunc) {
	g := router.Group("/room", authMiddleware)
	g.GET("/all", h.ListAllRooms)
	g.POST("/create", h.CreateRoom)
	g.GET("/:id", h.GetRoom)
	g.DELETE("/:id", h.DeleteRoom)
	g.PUT("/:id", h.UpdateRoom)
	g.POST("/:id/book", h.BookRoom)
}

// ListAllRooms godoc
// @Summary List all rooms
// @Description List all rooms
// @Tags room
// @Accept json
// @Produce json
// @Success 200 {object} gin.H{total=int,rooms=[]model.Room}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/room/all [get]
// @Security BearerAuth
func (h *RoomHandler) ListAllRooms(c *gin.Context) {
	resp, err := h.roomService.ListAllRooms(c.Request.Context())
	if err != nil {
		log.Println("Error listing rooms:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rooms": resp, "total": len(resp)})
}

// CreateRoom godoc
// @Summary Create room
// @Description Create room
// @Tags room
// @Accept json
// @Produce json
// @Param request body model.CreateRoomRequest true "Create room request"
// @Success 201 {object} gin.H{message=string,room=model.Room}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/room/create [post]
// @Security BearerAuth
func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var req model.CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}
	room, err := h.roomService.CreateRoom(c.Request.Context(), &req)
	if err != nil {
		log.Println("Error creating room:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Room created successfully", "room": room})
}

// GetRoom godoc
// @Summary Get room
// @Description Get room
// @Tags room
// @Accept json
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} gin.H{room=model.Room}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/room/{id} [get]
// @Security BearerAuth
func (h *RoomHandler) GetRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("Error converting ID:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}
	room, err := h.roomService.GetRoomByID(c.Request.Context(), id)
	if err != nil {
		log.Println("Error getting room:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"room": room})
}

// DeleteRoom godoc
// @Summary Delete room
// @Description Delete room
// @Tags room
// @Accept json
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/room/{id} [delete]
// @Security BearerAuth
func (h *RoomHandler) DeleteRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("Error converting ID:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}
	if err := h.roomService.DeleteRoom(c.Request.Context(), id); err != nil {
		log.Println("Error deleting room:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}

// UpdateRoom godoc
// @Summary Update room
// @Description Update room
// @Tags room
// @Accept json
// @Produce json
// @Param id path int true "Room ID"
// @Param request body model.UpdateRoomRequest true "Update room request"
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/room/{id} [put]
// @Security BearerAuth
func (h *RoomHandler) UpdateRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("Error converting ID:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}
	var req model.UpdateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}
	if err := h.roomService.UpdateRoom(c.Request.Context(), id, &req); err != nil {
		log.Println("Error updating room:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Room updated successfully"})
}

// BookRoom godoc
// @Summary Book room
// @Description Book room
// @Tags room
// @Accept json
// @Produce json
// @Param id path int true "Room ID"
// @Param request body model.CreateBookingRequest true "Book room request"
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/room/{id}/book [post]
// @Security BearerAuth
func (h *RoomHandler) BookRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("Error converting ID:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}
	var req model.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  err.Error(),
		})
		return
	}

	if req.FromDate.After(req.TillDate) {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad request",
			Detail:  "From date must be before till date",
		})
		return
	}

	user := c.MustGet("user").(*model.User)

	if err := h.roomService.BookRoom(c.Request.Context(), user.ID, id, &req); err != nil {
		log.Println("Error booking room:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room booked successfully"})
}
