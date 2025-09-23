package repository

import (
	"context"
	"time"

	"github.com/tunangoo/full-time-go-dev/hotel-reservation/internal/model"
	"github.com/uptrace/bun"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, booking *model.Booking) error
	ListBookings(ctx context.Context, req model.ListBookingsRequest) ([]*model.Booking, error)
	GetBookingByID(ctx context.Context, id int64) (*model.Booking, error)
	CancelBooking(ctx context.Context, id int64) error
	ListBookingByUserID(ctx context.Context, userID int64) ([]*model.Booking, error)
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

func (r *bookingRepository) GetBookingByID(ctx context.Context, id int64) (*model.Booking, error) {
	var booking model.Booking
	err := r.db.NewSelect().Model(&booking).Where("id = ?", id).Scan(ctx)
	return &booking, err
}

func (r *bookingRepository) CancelBooking(ctx context.Context, id int64) error {
	_, err := r.db.NewUpdate().Model(&model.Booking{}).Where("id = ?", id).Set("cancelled = TRUE").Exec(ctx)
	return err
}

func (r *bookingRepository) ListBookingByUserID(ctx context.Context, userID int64) ([]*model.Booking, error) {
	var bookings []*model.Booking
	query := r.db.NewSelect().Model(&bookings)

	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}

	err := query.Scan(ctx)
	return bookings, err
}
