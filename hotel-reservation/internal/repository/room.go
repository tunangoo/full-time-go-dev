package repository

import (
	"context"
	"time"

	"github.com/tunangoo/full-time-go-dev/internal/model"
	"github.com/uptrace/bun"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *model.Room) error
	ListRooms(ctx context.Context) ([]*model.Room, error)
	GetRoomByID(ctx context.Context, id int64) (*model.Room, error)
	UpdateRoom(ctx context.Context, room *model.Room) error
	DeleteRoom(ctx context.Context, id int64) error
}

type roomRepository struct {
	db *bun.DB
}

func NewRoomRepository(db *bun.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) CreateRoom(ctx context.Context, room *model.Room) error {
	room.CreatedAt = time.Now()
	room.UpdatedAt = time.Now()

	_, err := r.db.NewInsert().Model(room).Exec(ctx)
	return err
}

func (r *roomRepository) ListRooms(ctx context.Context) ([]*model.Room, error) {
	var rooms []*model.Room
	err := r.db.NewSelect().Model(&rooms).Scan(ctx)
	return rooms, err
}

func (r *roomRepository) GetRoomByID(ctx context.Context, id int64) (*model.Room, error) {
	var room model.Room
	err := r.db.NewSelect().Model(&room).Where("id = ?", id).Scan(ctx)
	return &room, err
}

func (r *roomRepository) UpdateRoom(ctx context.Context, room *model.Room) error {
	room.UpdatedAt = time.Now()
	_, err := r.db.NewUpdate().Model(room).Where("id = ?", room.ID).Exec(ctx)
	return err
}

func (r *roomRepository) DeleteRoom(ctx context.Context, id int64) error {
	_, err := r.db.NewDelete().Model(&model.Room{}).Where("id = ?", id).Exec(ctx)
	return err
}
