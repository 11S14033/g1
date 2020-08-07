package grpc

import (
	"context"
	"log"

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

func (rs RoomRPCService) SaveRoom(ctx context.Context, req *roompb.CreateRoomRequest) (res *roompb.CreateRoomResponse, err error) {
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
	}

	rs.usecase.SaveRoom(ctx, data)

	return res, nil
}
func (rs RoomRPCService) GetRooms(req *roompb.GetRoomsRequest, stream roompb.RoomService_GetRoomsServer) error {
	datas, err := rs.usecase.GetRooms(context.Background())

	if err != nil {
		log.Fatalln("Error at service [roomGRPCService][GetRooms] , when call service [roomUseCase][GetRooms]", err)
	}
	for _, data := range datas {
		stream.Send(&roompb.GetRoomsResponse{
			Room: dataToPBRoom(data),
		})
	}
	return nil
}
func (rs RoomRPCService) GetRoomByID(ctx context.Context, req *roompb.GetRoomByIDRequest) (*roompb.GetRoomByIDResponse, error) {
	rid := req.GetRoomId()

	room, err := rs.usecase.GetRoomByID(rid)
	if err != nil {
		log.Fatalln("Error at service [roomGRPCService][GetRoomByID] , when call service [roomUseCase][GetRoomByID]", err)
	}
	return &roompb.GetRoomByIDResponse{
		Room: dataToPBRoom(room),
	}, nil

}
func (rs RoomRPCService) UpdateRoomByID(ctx context.Context, req *roompb.UpdateRoomByIDRequest) (*roompb.UpdateRoomByIDResponse, error) {
	room := req.GetRoom()
	newRoom := models.Room{
		ID:          room.GetId(),
		Address:     room.GetAddres(),
		City:        room.GetCity(),
		Description: room.GetDescriptions(),
		Owner:       room.GetOwner(),
		Province:    room.GetProvince(),
		Type:        room.GetType(),
	}
	data, err := rs.usecase.UpdateRoom(ctx, newRoom)
	if err != nil {
		log.Fatalln("Error at service [roomGRPCService][UpdateRoomByID] , when call service [roomUseCase][UpdateRoomByID]", err)
	}
	return &roompb.UpdateRoomByIDResponse{
		Room: dataToPBRoom(data),
	}, nil
}
func (rs RoomRPCService) DeleteRoom(ctx context.Context, req *roompb.DeleteRoomRequest) (res *roompb.DeleteRoomResponse, err error) {
	rid := req.GetRoomId()
	err = rs.usecase.DeleteRoom(ctx, rid)
	if err != nil {
		log.Fatalln("Error at service [roomGRPCService][DeleteRoom] , when call service [roomUseCase][DeleteRoom]", err)
	}
	return &roompb.DeleteRoomResponse{
		RoomId: rid,
	}, nil
}
