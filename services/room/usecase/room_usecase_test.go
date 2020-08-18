package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/11s14033/g1/services/room/commons/models"
	roomMock "github.com/11s14033/g1/services/room/repository/mock"
	"github.com/stretchr/testify/mock"
)

type fields struct {
	Room  models.Room
	Rooms []models.Room
	rid   uint64
}

func TestGetRoomByID(t *testing.T) {
	tMock := time.Now()

	room := models.Room{
		ID:          1,
		RoomName:    "Test",
		Owner:       "Test",
		Address:     "Test",
		City:        "Test",
		Province:    "Test",
		Description: "Test",
		Type:        1,
		CreatedAt:   tMock,
		UpdatedAt:   tMock,
	}

	succesFields := fields{
		Room: room,
		rid:  1,
	}

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: succesFields,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepo := new(roomMock.RoomRepositoryMock)

			mockRepo.On("GetRoomByID", mock.AnythingOfType("uint64")).Return(test.Fields.Room, nil)

			roomUC := NewRoomUseCase(mockRepo)

			res, err := roomUC.GetRoomByID(context.Background(), test.Fields.rid)

			assert.NoError(t, err)
			assert.NotNil(t, res)
			if !reflect.DeepEqual(res, test.Fields.Room) {
				t.Errorf("TestGetRoomByID() = %v, want %v", res, test.Fields.Room)
			}
		})
	}
}

func TestGetRooms(t *testing.T) {
	tMock := time.Now()

	rooms := []models.Room{
		models.Room{
			ID:          1,
			RoomName:    "Test",
			Owner:       "Test",
			Address:     "Test",
			City:        "Test",
			Province:    "Test",
			Description: "Test",
			Type:        1,
			CreatedAt:   tMock,
			UpdatedAt:   tMock,
		},

		models.Room{
			ID:          2,
			RoomName:    "Test2",
			Owner:       "Test2",
			Address:     "Test2",
			City:        "Test2",
			Province:    "Test2",
			Description: "Test2",
			Type:        0,
			CreatedAt:   tMock,
			UpdatedAt:   tMock,
		},
	}

	succesFields := fields{
		Rooms: rooms,
	}

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: succesFields,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepo := new(roomMock.RoomRepositoryMock)

			mockRepo.On("GetRooms", mock.Anything).Return(test.Fields.Rooms, nil)

			roomUC := NewRoomUseCase(mockRepo)

			res, err := roomUC.GetRooms(context.Background())

			if !reflect.DeepEqual(res, test.Fields.Rooms) {
				t.Errorf("TestGetRooms() = %v, want %v", res, test.Fields.Rooms)
			}
			assert.NoError(t, err)
			assert.Len(t, res, len(test.Fields.Rooms))
		})
	}
}

func TestSaveRoom(t *testing.T) {
	tMock := time.Now()
	room := models.Room{
		ID:          1,
		RoomName:    "Test",
		Owner:       "Test",
		Address:     "Test",
		City:        "Test",
		Province:    "Test",
		Description: "Test",
		Type:        1,
		CreatedAt:   tMock,
		UpdatedAt:   tMock,
	}

	succesFields := fields{
		Room: room,
	}

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: succesFields,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepo := new(roomMock.RoomRepositoryMock)

			mockRepo.On("SaveRoom", mock.AnythingOfType("models.Room")).Return(nil)

			roomUC := NewRoomUseCase(mockRepo)

			err := roomUC.SaveRoom(context.Background(), test.Fields.Room)

			assert.NoError(t, err)

		})
	}

}

func TestUpdateRoom(t *testing.T) {
	tMock := time.Now()
	room := models.Room{
		ID:          1,
		RoomName:    "Test",
		Owner:       "Test",
		Address:     "Test",
		City:        "Test",
		Province:    "Test",
		Description: "Test",
		Type:        1,
		CreatedAt:   tMock,
		UpdatedAt:   tMock,
	}

	succesFields := fields{
		Room: room,
	}

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: succesFields,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepo := new(roomMock.RoomRepositoryMock)

			mockRepo.On("GetRoomByID", mock.AnythingOfType("uint64")).Return(test.Fields.Room, nil)
			mockRepo.On("UpdateRoom", mock.AnythingOfType("models.Room")).Return(test.Fields.Room, nil)

			roomUC := NewRoomUseCase(mockRepo)

			res, err := roomUC.UpdateRoom(context.Background(), test.Fields.Room)

			assert.NoError(t, err)
			if !reflect.DeepEqual(res, test.Fields.Room) {
				t.Errorf("TestUpdateRoom() = %v, want %v", res, test.Fields.Room)
			}

		})
	}

}

func TestDeleteRoom(t *testing.T) {
	successFields := fields{
		rid: 1,
	}

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: successFields,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mockRepo := new(roomMock.RoomRepositoryMock)
			mockRepo.On("DeleteRoom", mock.AnythingOfType("uint64")).Return(true, nil)

			roomUC := NewRoomUseCase(mockRepo)
			err := roomUC.DeleteRoom(context.Background(), test.Fields.rid)

			assert.NoError(t, err)
		})
	}
}
