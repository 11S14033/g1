package grpc

import (
	"context"
	"log"
	"time"

	"github.com/11s14033/g1/services/room/commons/models"
	roompb "github.com/11s14033/g1/services/room/commons/pb"
)

func dataToPBRoom(data models.Room) *roompb.Room {
	return &roompb.Room{

		Id:           data.ID,
		RoomName:     data.RoomName,
		Addres:       data.Address,
		City:         data.City,
		Province:     data.Province,
		Descriptions: data.Description,
		Owner:        data.Owner,
	}
}

func (rs RoomRPCService) GetRooms(req *roompb.GetRoomsRequest, stream roompb.RoomService_GetRoomsServer) error {
	datas, err := rs.usecase.GetRooms(context.Background())

	if err != nil {
		log.Printf("[Error][roomGRPCService][GetRooms]-[when calling][roomUseCase][GetRooms]-[%v]", err)
	}
	for _, data := range datas {
		stream.Send(&roompb.GetRoomsResponse{
			Room: dataToPBRoom(data),
		})
	}
	return nil
}
func (rs RoomRPCService) GetRoomByID(ctx context.Context, req *roompb.GetRoomByIDRequest) (*roompb.GetRoomByIDResponse, error) {
	var err error
	rid := req.GetRoomId()

	room, err := rs.usecase.GetRoomByID(ctx, rid)
	if err != nil {
		log.Printf("[Error][roomGRPCService][GetRoomByID]-[when calling][roomUseCase][GetRoomByID]-[%v]", err)
	}
	return &roompb.GetRoomByIDResponse{
		Room: dataToPBRoom(room),
	}, nil

}
func (rs RoomRPCService) UpdateRoomByID(ctx context.Context, req *roompb.UpdateRoomByIDRequest) (*roompb.UpdateRoomByIDResponse, error) {
	room := req.GetRoom()
	newRoom := models.Room{
		ID:          room.GetId(),
		RoomName:    room.GetRoomName(),
		Address:     room.GetAddres(),
		City:        room.GetCity(),
		Description: room.GetDescriptions(),
		Owner:       room.GetOwner(),
		Province:    room.GetProvince(),
		Type:        room.GetType(),
	}
	data, err := rs.usecase.UpdateRoom(ctx, newRoom)
	if err != nil {
		log.Printf("[Error][roomGRPCService][UpdateRoomByID]-[when calling][roomUseCase][UpdateRoomByID]-[%v]", err)
	}
	return &roompb.UpdateRoomByIDResponse{
		Room: dataToPBRoom(data),
	}, nil
}

func (rs RoomRPCService) SaveRoom(ctx context.Context, req *roompb.CreateRoomRequest) (*roompb.CreateRoomResponse, error) {
	room := req.GetRoom()
	data := models.Room{
		ID:          room.GetId(),
		Address:     room.GetAddres(),
		City:        room.GetCity(),
		Description: room.GetDescriptions(),
		Owner:       room.GetOwner(),
		Province:    room.GetProvince(),
		RoomName:    room.GetRoomName(),
		Type:        room.GetType(),
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}

	err := rs.usecase.SaveRoom(ctx, data)
	if err != nil {
		log.Printf("[Error][roomGRPCService][SaveRoom]-[when calling][roomUseCase][SaveRoom]-[%v]", err)
	}

	return &roompb.CreateRoomResponse{
		Room: dataToPBRoom(data),
	}, nil
}

func (rs RoomRPCService) DeleteRoom(ctx context.Context, req *roompb.DeleteRoomRequest) (res *roompb.DeleteRoomResponse, err error) {
	rid := req.GetRoomId()
	err = rs.usecase.DeleteRoom(context.Background(), rid)

	if err != nil {
		log.Printf("[Error][roomGRPCService][DeleteRoom]-[when calling][roomUseCase][DeleteRoom]-[%v]", err)
	}
	return &roompb.DeleteRoomResponse{
		RoomId: rid,
	}, nil
}
