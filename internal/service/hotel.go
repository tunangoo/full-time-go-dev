package service

import (
	"context"
	"log"

	"github.com/tunangoo/full-time-go-dev/internal/model"
	"github.com/tunangoo/full-time-go-dev/internal/repository"
)

type HotelService interface {
	GetHotelByID(ctx context.Context, id int64) (*model.Hotel, error)
	CreateHotel(ctx context.Context, req *model.CreateHotelRequest) (*model.Hotel, error)
	ListAllHotels(ctx context.Context) ([]*model.Hotel, error)
	DeleteHotel(ctx context.Context, id int64) error
	UpdateHotel(ctx context.Context, id int64, req *model.UpdateHotelRequest) error
}

type hotelService struct {
	hotelRepository repository.HotelRepository
}

func NewHotelService(
	hotelRepository repository.HotelRepository,
) HotelService {
	return &hotelService{
		hotelRepository: hotelRepository,
	}
}

func (s *hotelService) GetHotelByID(ctx context.Context, id int64) (*model.Hotel, error) {
	hotel, err := s.hotelRepository.GetHotelByID(ctx, id)
	if err != nil {
		log.Println("Error getting hotel by ID:", err)
		return nil, err
	}
	return hotel, nil
}

func (s *hotelService) CreateHotel(ctx context.Context, req *model.CreateHotelRequest) (*model.Hotel, error) {
	hotel := &model.Hotel{
		Name:     req.Name,
		Location: req.Location,
		Rating:   req.Rating,
	}

	if err := s.hotelRepository.CreateHotel(ctx, hotel); err != nil {
		log.Println("Error creating hotel:", err)
		return nil, err
	}
	return hotel, nil
}

func (s *hotelService) ListAllHotels(ctx context.Context) ([]*model.Hotel, error) {
	hotels, err := s.hotelRepository.ListHotels(ctx)
	if err != nil {
		log.Println("Error listing hotels:", err)
		return nil, err
	}
	return hotels, nil
}

func (s *hotelService) DeleteHotel(ctx context.Context, id int64) error {
	if err := s.hotelRepository.DeleteHotel(ctx, id); err != nil {
		log.Println("Error deleting hotel:", err)
		return err
	}
	return nil
}

func (s *hotelService) UpdateHotel(ctx context.Context, id int64, req *model.UpdateHotelRequest) error {
	hotel, err := s.hotelRepository.GetHotelByID(ctx, id)
	if err != nil {
		log.Println("Error getting hotel by ID:", err)
		return err
	}

	hotel.Name = req.Name
	hotel.Location = req.Location
	hotel.Rating = req.Rating

	if err := s.hotelRepository.UpdateHotel(ctx, hotel); err != nil {
		log.Println("Error updating hotel:", err)
		return err
	}
	return nil
}
