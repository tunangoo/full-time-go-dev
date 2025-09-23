package service

import (
	"context"
	"errors"
	"log"

	"github.com/tunangoo/full-time-go-dev/hotel-reservation/internal/model"
	"github.com/tunangoo/full-time-go-dev/hotel-reservation/internal/repository"
)

type RoomService interface {
	GetRoomByID(ctx context.Context, id int64) (*model.Room, error)
	CreateRoom(ctx context.Context, req *model.CreateRoomRequest) (*model.Room, error)
	ListAllRooms(ctx context.Context) ([]*model.Room, error)
	DeleteRoom(ctx context.Context, id int64) error
	UpdateRoom(ctx context.Context, id int64, req *model.UpdateRoomRequest) error
	BookRoom(ctx context.Context, userID int64, roomID int64, req *model.CreateBookingRequest) error
}

type roomService struct {
	roomRepository    repository.RoomRepository
	hotelRepository   repository.HotelRepository
	bookingRepository repository.BookingRepository
}

func NewRoomService(
	roomRepository repository.RoomRepository,
	hotelRepository repository.HotelRepository,
	bookingRepository repository.BookingRepository,
) RoomService {
	return &roomService{
		roomRepository:    roomRepository,
		hotelRepository:   hotelRepository,
		bookingRepository: bookingRepository,
	}
}

func (s *roomService) GetRoomByID(ctx context.Context, id int64) (*model.Room, error) {
	room, err := s.roomRepository.GetRoomByID(ctx, id)
	if err != nil {
		log.Println("Error getting room by ID:", err)
		return nil, err
	}
	return room, nil
}

func (s *roomService) CreateRoom(ctx context.Context, req *model.CreateRoomRequest) (*model.Room, error) {
	// Ensure referenced hotel exists
	if _, err := s.hotelRepository.GetHotelByID(ctx, req.HotelID); err != nil {
		log.Println("Referenced hotel not found:", err)
		return nil, err
	}

	room := &model.Room{
		Type:      req.Type,
		BasePrice: req.BasePrice,
		HotelID:   req.HotelID,
		Size:      req.Size,
	}
	if err := s.roomRepository.CreateRoom(ctx, room); err != nil {
		log.Println("Error creating room:", err)
		return nil, err
	}
	return room, nil
}

func (s *roomService) ListAllRooms(ctx context.Context) ([]*model.Room, error) {
	rooms, err := s.roomRepository.ListRooms(ctx)
	if err != nil {
		log.Println("Error listing rooms:", err)
		return nil, err
	}
	return rooms, nil
}

func (s *roomService) DeleteRoom(ctx context.Context, id int64) error {
	if err := s.roomRepository.DeleteRoom(ctx, id); err != nil {
		log.Println("Error deleting room:", err)
		return err
	}
	return nil
}

func (s *roomService) UpdateRoom(ctx context.Context, id int64, req *model.UpdateRoomRequest) error {
	room, err := s.roomRepository.GetRoomByID(ctx, id)
	if err != nil {
		log.Println("Error getting room by ID:", err)
		return err
	}

	if req.HotelID != 0 && req.HotelID != room.HotelID {
		if _, err := s.hotelRepository.GetHotelByID(ctx, req.HotelID); err != nil {
			log.Println("Referenced hotel not found:", err)
			return err
		}
		room.HotelID = req.HotelID
	}

	room.Type = req.Type
	room.BasePrice = req.BasePrice
	room.Size = req.Size

	if err := s.roomRepository.UpdateRoom(ctx, room); err != nil {
		log.Println("Error updating room:", err)
		return err
	}
	return nil
}

func (s *roomService) BookRoom(ctx context.Context, userID int64, roomID int64, req *model.CreateBookingRequest) error {
	// Step 1: Check if the room is exists
	_, err := s.roomRepository.GetRoomByID(ctx, roomID)
	if err != nil {
		log.Println("Error getting room:", err)
		return err
	}

	// Step 2: Check if the room is available
	bookings, err := s.bookingRepository.ListBookings(ctx, model.ListBookingsRequest{
		RoomID:   roomID,
		FromDate: &req.FromDate,
		TillDate: &req.TillDate,
	})
	if err != nil {
		log.Println("Error getting bookings:", err)
		return err
	}
	if len(bookings) > 0 {
		log.Println("Room is not available")
		return errors.New("room is not available")
	}

	// Step 3: Create the booking
	booking := &model.Booking{
		UserID:        userID,
		RoomID:        roomID,
		NumberPersons: req.NumberPersons,
		FromDate:      req.FromDate,
		TillDate:      req.TillDate,
	}
	err = s.bookingRepository.CreateBooking(ctx, booking)
	if err != nil {
		log.Println("Error creating booking:", err)
		return err
	}
	return nil
}
