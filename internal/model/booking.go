package model

import "time"

type Booking struct {
	ID            int64     `json:"id" bun:"id,pk,autoincrement"`
	UserID        int64     `json:"user_id" bun:"user_id"`
	RoomID        int64     `json:"room_id" bun:"room_id"`
	NumberPersons int64     `json:"number_persons" bun:"number_persons"`
	FromDate      time.Time `json:"from_date" bun:"from_date"`
	TillDate      time.Time `json:"till_date" bun:"till_date"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bun:"updated_at"`
}

type CreateBookingRequest struct {
	NumberPersons int64     `json:"number_persons" bun:"number_persons"`
	FromDate      time.Time `json:"from_date" bun:"from_date"`
	TillDate      time.Time `json:"till_date" bun:"till_date"`
}

type ListBookingsRequest struct {
	RoomID   int64      `json:"room_id" bun:"room_id"`
	FromDate *time.Time `json:"from_date" bun:"from_date"`
	TillDate *time.Time `json:"till_date" bun:"till_date"`
}
