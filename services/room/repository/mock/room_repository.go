package mock

import (
	"github.com/11s14033/g1/services/room/commons/models"

	"context"

	"github.com/stretchr/testify/mock"
)

type RoomRepositoryMock struct {
	mock.Mock
}

func (m *RoomRepositoryMock) GetRoomByID(ctx context.Context, rid uint64) (room models.Room, err error) {
	args := m.Called(rid)
	return args.Get(0).(models.Room), args.Error(1)
}

func (m *RoomRepositoryMock) SaveRoom(ctx context.Context, room models.Room) (err error) {
	args := m.Called(room)
	return args.Error(0)
}

func (m *RoomRepositoryMock) UpdateRoom(ctx context.Context, room models.Room) (newRoom models.Room, err error) {
	args := m.Called(room)
	return args.Get(0).(models.Room), args.Error(1)
}
func (m *RoomRepositoryMock) DeleteRoom(ctx context.Context, rid uint64) (err error) {
	args := m.Called(rid)
	return args.Error(1)
}

func (m *RoomRepositoryMock) GetRooms(ctx context.Context) (rooms []models.Room, err error) {
	var allRooms []models.Room
	args := m.Called(allRooms)
	return args.Get(0).([]models.Room), args.Error(1)
}
