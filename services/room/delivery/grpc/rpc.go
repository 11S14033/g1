package grpc

import (
	"github.com/11s14033/g1/services/room/usecase"
)

type RoomRPCService struct {
	usecase usecase.RoomUseCase
}

func NewRoomRPCService(u usecase.RoomUseCase) RoomRPCService {
	return RoomRPCService{
		usecase: u,
	}
}
