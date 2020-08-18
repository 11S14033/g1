package rest

import (
	"github.com/11s14033/g1/services/room/usecase"
)

type RoomRestService struct {
	usecase usecase.RoomUseCase
}

func NewRoomRPCService(u usecase.RoomUseCase) RoomRestService {
	return RoomRestService{
		usecase: u,
	}
}
