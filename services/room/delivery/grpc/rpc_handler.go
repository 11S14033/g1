package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"

	"github.com/11s14033/g1/services/room/commons/models"
	roompb "github.com/11s14033/g1/services/room/commons/pb"
	utils "github.com/11s14033/g1/services/room/delivery/utils"
)

func dataToPBRoom(data models.Room) *roompb.Room {
	return &roompb.Room{

		Id:          data.ID,
		RoomName:    data.RoomName,
		Addres:      data.Address,
		City:        data.City,
		Province:    data.Province,
		Description: data.Description,
		Owner:       data.Owner,
		CreatedAt:   data.CreatedAt.String(),
		UpdatedAt:   data.UpdatedAt.String(),
	}
}

func (rs RoomRPCService) GetRooms(req *roompb.GetRoomsRequest, stream roompb.RoomService_GetRoomsServer) error {
	datas, err := rs.usecase.GetRooms(context.Background())

	if err != nil {
		desc := fmt.Sprintf("[Error][roomGRPCService][GetRooms]-[when calling][roomUseCase][GetRooms]-[%v]", err)
		log.Println(desc)
		return utils.ErrorRPC(codes.Internal, desc)
	}
	for _, data := range datas {
		stream.Send(dataToPBRoom(data))
	}
	return nil
}
func (rs RoomRPCService) GetRoomByID(ctx context.Context, req *roompb.GetRoomByIDRequest) (*roompb.Room, error) {
	var err error
	rid := req.GetRoomId()

	room, err := rs.usecase.GetRoomByID(ctx, rid)
	if err != nil {
		desc := fmt.Sprintf("[Error][roomGRPCService][GetRoomByID]-[when calling][roomUseCase][GetRoomByID]-[%v]", err)
		log.Println(desc)
		return nil, utils.ErrorRPC(codes.Internal, desc)
	}
	return dataToPBRoom(room), nil

}
func (rs RoomRPCService) UpdateRoomByID(ctx context.Context, req *roompb.Room) (*roompb.Room, error) {

	newRoom := models.Room{
		ID:          req.GetId(),
		RoomName:    req.GetRoomName(),
		Address:     req.GetAddres(),
		City:        req.GetCity(),
		Description: req.GetDescription(),
		Owner:       req.GetOwner(),
		Province:    req.GetProvince(),
		Type:        req.GetType(),
		UpdatedAt:   time.Now(),
	}
	data, err := rs.usecase.UpdateRoom(ctx, newRoom)
	if err != nil {
		desc := fmt.Sprintf("[Error][roomGRPCService][UpdateRoomByID]-[when calling][roomUseCase][UpdateRoomByID]-[%v]", err)
		log.Println(desc)
		return nil, utils.ErrorRPC(codes.Internal, desc)
	}
	return dataToPBRoom(data), nil
}

func (rs RoomRPCService) SaveRoom(ctx context.Context, req *roompb.Room) (*roompb.Room, error) {
	data := models.Room{
		ID:          req.GetId(),
		Address:     req.GetAddres(),
		City:        req.GetCity(),
		Description: req.GetDescription(),
		Owner:       req.GetOwner(),
		Province:    req.GetProvince(),
		RoomName:    req.GetRoomName(),
		Type:        req.GetType(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := rs.usecase.SaveRoom(ctx, data)
	if err != nil {
		desc := fmt.Sprintf("[Error][roomGRPCService][SaveRoom]-[when calling][roomUseCase][SaveRoom]-[%v]", err)
		log.Println(desc)
		return nil, utils.ErrorRPC(codes.Internal, desc)
	}

	return dataToPBRoom(data), nil
}

func (rs RoomRPCService) DeleteRoom(ctx context.Context, req *roompb.DeleteRoomRequest) (*roompb.DeleteRoomResponse, error) {
	rid := req.GetRoomId()
	err := rs.usecase.DeleteRoom(context.Background(), rid)

	if err != nil {

		desc := fmt.Sprintf("[Error][roomGRPCService][DeleteRoom]-[when calling][roomUseCase][DeleteRoom]-[%v]", err)
		log.Println(desc)
		return nil, utils.ErrorRPC(codes.Internal, desc)
	}
	return &roompb.DeleteRoomResponse{
		RoomId: rid,
	}, nil
}
