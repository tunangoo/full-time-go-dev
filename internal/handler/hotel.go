package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tunangoo/full-time-go-dev/internal/model"
	"github.com/tunangoo/full-time-go-dev/internal/service"
)

type HotelHandler struct {
	hotelService service.HotelService
}

func NewHotelHandler(
	hotelService service.HotelService,
) *HotelHandler {
	return &HotelHandler{hotelService: hotelService}
}

func (h *HotelHandler) RegisterRoutes(router gin.IRouter) {
	g := router.Group("/hotel")
	g.GET("/all", h.ListAllHotels)
	g.POST("/create", h.CreateHotel)
	g.GET("/:id", h.GetHotel)
	g.DELETE("/:id", h.DeleteHotel)
	g.PUT("/:id", h.UpdateHotel)
}

// ListAllHotels godoc
// @Summary List all hotels
// @Description List all hotels
// @Tags hotel
// @Accept json
// @Produce json
// @Success 200 {object} gin.H{total=int,hotels=[]model.Hotel}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/hotel/all [get]
func (h *HotelHandler) ListAllHotels(c *gin.Context) {
	resp, err := h.hotelService.ListAllHotels(c.Request.Context())
	if err != nil {
		log.Println("Error listing hotels:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: http.StatusInternalServerError, Message: "Internal server error", Detail: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"hotels": resp, "total": len(resp)})
}

// CreateHotel godoc
// @Summary Create hotel
// @Description Create hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param request body model.CreateHotelRequest true "Create hotel request"
// @Success 201 {object} gin.H{message=string,hotel=model.Hotel}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/hotel/create [post]
func (h *HotelHandler) CreateHotel(c *gin.Context) {
	var req model.CreateHotelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: "Bad request", Detail: err.Error()})
		return
	}
	hotel, err := h.hotelService.CreateHotel(c.Request.Context(), &req)
	if err != nil {
		log.Println("Error creating hotel:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: http.StatusInternalServerError, Message: "Internal server error", Detail: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Hotel created successfully", "hotel": hotel})
}

// GetHotel godoc
// @Summary Get hotel
// @Description Get hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path int true "Hotel ID"
// @Success 200 {object} gin.H{hotel=model.Hotel}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/hotel/{id} [get]
func (h *HotelHandler) GetHotel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("Error converting ID:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: "Bad request", Detail: err.Error()})
		return
	}
	hotel, err := h.hotelService.GetHotelByID(c.Request.Context(), id)
	if err != nil {
		log.Println("Error getting hotel:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: http.StatusInternalServerError, Message: "Internal server error", Detail: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"hotel": hotel})
}

// DeleteHotel godoc
// @Summary Delete hotel
// @Description Delete hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path int true "Hotel ID"
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/hotel/{id} [delete]
func (h *HotelHandler) DeleteHotel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("Error converting ID:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: "Bad request", Detail: err.Error()})
		return
	}
	if err := h.hotelService.DeleteHotel(c.Request.Context(), id); err != nil {
		log.Println("Error deleting hotel:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: http.StatusInternalServerError, Message: "Internal server error", Detail: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted successfully"})
}

// UpdateHotel godoc
// @Summary Update hotel
// @Description Update hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path int true "Hotel ID"
// @Param request body model.UpdateHotelRequest true "Update hotel request"
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} model.ErrorResponse
// @Router /v1/hotel/{id} [put]
func (h *HotelHandler) UpdateHotel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("Error converting ID:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: "Bad request", Detail: err.Error()})
		return
	}
	var req model.UpdateHotelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Code: http.StatusBadRequest, Message: "Bad request", Detail: err.Error()})
		return
	}
	if err := h.hotelService.UpdateHotel(c.Request.Context(), id, &req); err != nil {
		log.Println("Error updating hotel:", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Code: http.StatusInternalServerError, Message: "Internal server error", Detail: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hotel updated successfully"})
}
