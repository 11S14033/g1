package usecase

import (
	"context"

	"github.com/11s14033/g1/services/room/commons/models"
)

type RoomUseCase interface {
	GetRoomByID(ctx context.Context, rid uint64) (room models.Room, err error)
	SaveRoom(ctx context.Context, room models.Room) (err error)
	UpdateRoom(ctx context.Context, room models.Room) (newRoom models.Room, err error)
	DeleteRoom(ctx context.Context, rid uint64) (err error)
	GetRooms(ctx context.Context) (rooms []models.Room, err error)
}
