package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Hotel struct {
	bun.BaseModel `bun:"hotels"`
	ID            int64     `json:"id" bun:"id,pk,autoincrement"`
	Name          string    `json:"name" bun:"name"`
	Location      string    `json:"location" bun:"location"`
	Rooms         []*Room   `json:"rooms" bun:"rooms,rel:has-many,join:id=hotel_id"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bun:"updated_at"`
	Rating        float64   `json:"rating" bun:"rating"`
}

type CreateHotelRequest struct {
	Name     string  `json:"name" binding:"required,min=3,max=255"`
	Location string  `json:"location" binding:"required,min=3,max=255"`
	Rating   float64 `json:"rating" binding:"required,min=0,max=5"`
}

type UpdateHotelRequest struct {
	Name     string  `json:"name" binding:"required,min=3,max=255"`
	Location string  `json:"location" binding:"required,min=3,max=255"`
	Rating   float64 `json:"rating" binding:"required,min=0,max=5"`
}

type ListHotelsRequest struct {
	Page   int64  `form:"page"`
	Limit  int64  `form:"limit"`
	Search string `form:"search"`
}
