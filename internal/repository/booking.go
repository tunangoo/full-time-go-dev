package repository

import (
	"context"
	"time"

	"github.com/tunangoo/full-time-go-dev/internal/model"
	"github.com/uptrace/bun"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, booking *model.Booking) error
	ListBookings(ctx context.Context, req model.ListBookingsRequest) ([]*model.Booking, error)
}

type bookingRepository struct {
	db *bun.DB
}

func NewBookingRepository(db *bun.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) CreateBooking(ctx context.Context, booking *model.Booking) error {
	booking.CreatedAt = time.Now()
	booking.UpdatedAt = time.Now()

	_, err := r.db.NewInsert().Model(booking).Exec(ctx)
	return err
}

func (r *bookingRepository) ListBookings(ctx context.Context, req model.ListBookingsRequest) ([]*model.Booking, error) {
	var bookings []*model.Booking
	query := r.db.NewSelect().Model(&bookings)
	if req.RoomID != 0 {
		query = query.Where("room_id = ?", req.RoomID)
	}
	if req.FromDate != nil && req.TillDate != nil {
		query = query.Where("from_date <= ?", req.TillDate).Where("till_date >= ?", req.FromDate)
	}
	err := query.Scan(ctx)
	return bookings, err

}
