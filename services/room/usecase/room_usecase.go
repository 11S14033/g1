package usecase

import (
	"context"
	"log"

	"github.com/11s14033/g1/services/room/commons/models"
	"github.com/11s14033/g1/services/room/repository"
)

type roomUseCase struct {
	roomRepository repository.RoomRepository
}

func NewRoomUseCase(repo repository.RoomRepository) RoomUseCase {
	return &roomUseCase{
		roomRepository: repo,
	}
}

func (roomUseCase *roomUseCase) GetRoomByID(ctx context.Context, rid uint64) (room models.Room, err error) {

	room, err = roomUseCase.roomRepository.GetRoomByID(ctx, rid)

	if err != nil {
		log.Printf("[Error][roomUseCase][GetRoomByID]-[when calling][roomRepository][GetRoomByID]-[%v]", err)
	}
	return room, nil
}
func (roomUseCase *roomUseCase) SaveRoom(ctx context.Context, room models.Room) (err error) {

	err = roomUseCase.roomRepository.SaveRoom(ctx, room)
	if err != nil {
		log.Printf("[Error][roomUseCase][SaveRoom]-[when calling][roomRepository][SaveRoom]-[%v]", err)
	}
	return nil
}
func (roomUseCase *roomUseCase) UpdateRoom(ctx context.Context, room models.Room) (newRoom models.Room, err error) {

	newRoom, err = roomUseCase.roomRepository.UpdateRoom(ctx, room)
	if err != nil {
		log.Printf("[Error][roomUseCase][UpdateRoom]-[when calling][roomRepository][UpdateRoom]-[%v]", err)
	}
	return newRoom, nil
}
func (roomUseCase *roomUseCase) DeleteRoom(ctx context.Context, rid uint64) (err error) {
	err = roomUseCase.roomRepository.DeleteRoom(ctx, rid)
	if err != nil {
		log.Printf("[Error][roomUseCase][DeleteRoom]-[when calling][roomRepository][DeleteRoom]-[%v]", err)
	}
	return nil
}
func (roomUseCase *roomUseCase) GetRooms(ctx context.Context) (rooms []models.Room, err error) {
	rooms, err = roomUseCase.roomRepository.GetRooms(ctx)
	if err != nil {
		log.Printf("[Error][roomUseCase][GetRooms]-[when calling][roomRepository][GetRooms]-[%v]", err)
	}
	return rooms, nil
}
