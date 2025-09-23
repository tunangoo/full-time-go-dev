package model

import (
	"time"

	"github.com/uptrace/bun"
)

type RoomType int

const (
	SingleRoomType RoomType = iota + 1
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type Room struct {
	bun.BaseModel `bun:"rooms"`
	ID            int64     `json:"id" bun:"id,pk,autoincrement"`
	Type          RoomType  `json:"type" bun:"type"`
	BasePrice     float64   `json:"base_price" bun:"base_price"`
	HotelID       int64     `json:"hotel_id" bun:"hotel_id"`
	Size          string    `json:"size" bun:"size"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bun:"updated_at"`
}

type CreateRoomRequest struct {
	Type      RoomType `json:"type" binding:"required"`
	BasePrice float64  `json:"base_price" binding:"required,min=0"`
	HotelID   int64    `json:"hotel_id" binding:"required"`
	Size      string   `json:"size" binding:"required,oneof=small normal kingsize"`
}

type UpdateRoomRequest struct {
	Type      RoomType `json:"type" binding:"required"`
	BasePrice float64  `json:"base_price" binding:"required,min=0"`
	HotelID   int64    `json:"hotel_id" binding:"required"`
	Size      string   `json:"size" binding:"required,oneof=small normal kingsize"`
}
