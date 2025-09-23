package service

import (
	"context"
	"errors"
	"log"

	"github.com/tunangoo/full-time-go-dev/hotel-reservation/internal/model"
	"github.com/tunangoo/full-time-go-dev/hotel-reservation/internal/repository"
)

type BookingService interface {
	GetBookingByID(ctx context.Context, id int64) (*model.Booking, error)
	ListAllBookings(ctx context.Context, userID int64) ([]*model.Booking, error)
	CancelBooking(ctx context.Context, id int64, userID int64) error
}

type bookingService struct {
	bookingRepository repository.BookingRepository
}

func NewBookingService(
	bookingRepository repository.BookingRepository,
) BookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
	}
}

func (s *bookingService) GetBookingByID(ctx context.Context, id int64) (*model.Booking, error) {
	booking, err := s.bookingRepository.GetBookingByID(ctx, id)
	if err != nil {
		log.Println("Error getting booking by ID:", err)
		return nil, err
	}
	return booking, nil
}

func (s *bookingService) ListAllBookings(ctx context.Context, userID int64) ([]*model.Booking, error) {
	bookings, err := s.bookingRepository.ListBookingByUserID(ctx, userID)
	if err != nil {
		log.Println("Error listing bookings by user ID:", err)
		return nil, err
	}

	return bookings, nil
}

func (s *bookingService) CancelBooking(ctx context.Context, id int64, userID int64) error {
	// Step 1: Get Booking
	booking, err := s.bookingRepository.GetBookingByID(ctx, id)
	if err != nil {
		log.Println("Error getting booking by ID:", err)
		return err
	}

	// Step 2: Check if user is the owner of the booking
	if booking.UserID != userID {
		log.Println("User is not the owner of the booking")
		return errors.New("user is not the owner of the booking")
	}

	// Step 3: Cancel the booking
	err = s.bookingRepository.CancelBooking(ctx, id)
	if err != nil {
		log.Println("Error canceling booking:", err)
		return err
	}
	return nil
}
