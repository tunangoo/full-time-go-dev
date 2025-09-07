package repository

import (
	"context"
	"time"

	"github.com/tunangoo/full-time-go-dev/internal/model"
	"github.com/uptrace/bun"
)

type HotelRepository interface {
	CreateHotel(ctx context.Context, hotel *model.Hotel) error
	ListHotels(ctx context.Context) ([]*model.Hotel, error)
	GetHotelByID(ctx context.Context, id int64) (*model.Hotel, error)
	UpdateHotel(ctx context.Context, hotel *model.Hotel) error
	DeleteHotel(ctx context.Context, id int64) error
}

type hotelRepository struct {
	db *bun.DB
}

func NewHotelRepository(db *bun.DB) HotelRepository {
	return &hotelRepository{db: db}
}

func (r *hotelRepository) CreateHotel(ctx context.Context, hotel *model.Hotel) error {
	hotel.CreatedAt = time.Now()
	hotel.UpdatedAt = time.Now()

	_, err := r.db.NewInsert().Model(hotel).Exec(ctx)
	return err
}

func (r *hotelRepository) ListHotels(ctx context.Context) ([]*model.Hotel, error) {
	var hotels []*model.Hotel
	err := r.db.NewSelect().Model(&hotels).Scan(ctx)
	return hotels, err
}

func (r *hotelRepository) GetHotelByID(ctx context.Context, id int64) (*model.Hotel, error) {
	var hotel model.Hotel
	err := r.db.NewSelect().Model(&hotel).Where("id = ?", id).Relation("Rooms").Scan(ctx)
	return &hotel, err
}

func (r *hotelRepository) UpdateHotel(ctx context.Context, hotel *model.Hotel) error {
	hotel.UpdatedAt = time.Now()
	_, err := r.db.NewUpdate().Model(hotel).Where("id = ?", hotel.ID).Exec(ctx)
	return err
}

func (r *hotelRepository) DeleteHotel(ctx context.Context, id int64) error {
	_, err := r.db.NewDelete().Model(&model.Hotel{}).Where("id = ?", id).Exec(ctx)
	return err
}
