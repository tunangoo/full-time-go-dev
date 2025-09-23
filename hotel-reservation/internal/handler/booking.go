package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tunangoo/full-time-go-dev/hotel-reservation/internal/model"
	"github.com/tunangoo/full-time-go-dev/hotel-reservation/internal/service"
)

type BookingHandler struct {
	bookingService service.BookingService
}

func NewBookingHandler(
	bookingService service.BookingService,
) *BookingHandler {
	return &BookingHandler{bookingService: bookingService}
}

func (h *BookingHandler) RegisterRoutes(router gin.IRouter, authMiddleware gin.HandlerFunc) {
	g := router.Group("/booking", authMiddleware)

	g.GET("/all", h.ListAllBookings)
	g.POST("/:id/cancel", h.CancelBooking)
	g.GET("/:id", h.GetBookingByID)
}

// ListAllBookings godoc
// @Summary List all bookings
// @Description List all bookings
// @Tags booking
// @Accept json
// @Produce json
// @Success 200 {object} gin.H{total=int,bookings=[]model.Booking}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/booking/all [get]
// @Security BearerAuth
func (h *BookingHandler) ListAllBookings(c *gin.Context) {

	user := c.MustGet("user").(*model.User)

	userID := user.ID
	if user.Role == "admin" {
		userID = 0
	}

	resp, err := h.bookingService.ListAllBookings(c.Request.Context(), userID)
	if err != nil {
		log.Println("Error listing bookings:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Detail:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"bookings": resp, "total": len(resp)})
}

// CancelBooking godoc
// @Summary Cancel booking
// @Description Cancel booking
// @Tags booking
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/booking/{id}/cancel [post]
// @Security BearerAuth
func (h *BookingHandler) CancelBooking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("Error converting ID:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Detail:  err.Error(),
		})
		return
	}

	user := c.MustGet("user").(*model.User)

	err = h.bookingService.CancelBooking(c.Request.Context(), id, user.ID)
	if err != nil {
		log.Println("Error canceling booking:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error canceling booking",
			Detail:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking canceled successfully"})
}

// GetBookingByID godoc
// @Summary Get booking
// @Description Get booking
// @Tags booking
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} gin.H{booking=model.Booking}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/booking/{id} [get]
// @Security BearerAuth
func (h *BookingHandler) GetBookingByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("Error converting ID:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Detail:  err.Error(),
		})
		return
	}
	booking, err := h.bookingService.GetBookingByID(c.Request.Context(), id)
	if err != nil {
		log.Println("Error getting booking:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error getting booking",
			Detail:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"booking": booking})
}
