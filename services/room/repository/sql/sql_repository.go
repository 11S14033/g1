package sql

import (
	"context"
	"log"

	"github.com/11s14033/g1/services/room/commons/models"
	"github.com/11s14033/g1/services/room/repository"

	"github.com/jinzhu/gorm"
)

type gormRepository struct {
	DB *gorm.DB
}

func NewGormRepository(db *gorm.DB) repository.RoomRepository {
	return &gormRepository{
		DB: db,
	}
}

func (roomRepository *gormRepository) GetRoomByID(rid uint64) (room models.Room, err error) {
	err = roomRepository.DB.Debug().Model(&models.Room{}).Where("id = ?", rid).Take(&room).Error
	if err != nil {
		log.Fatalln("Error when get room from database", err)
		return room, err
	}
	return room, nil
}

func (roomRepository *gormRepository) SaveRoom(ctx context.Context, room models.Room) (err error) {
	err = roomRepository.DB.Debug().Model(&models.Room{}).Create(&room).Error
	if err != nil {
		log.Fatalln("Error when create room to database", err)
		return err
	}
	return nil
}
func (roomRepository *gormRepository) UpdateRoom(ctx context.Context, room models.Room) (newRoom models.Room, err error) {
	err = roomRepository.DB.Debug().Model(&models.Room{}).Update(&room).Error
	if err != nil {
		log.Fatalln("Error when update room to database", err)
		return room, err
	}

	return room, nil
}

func (roomRepository *gormRepository) DeleteRoom(ctx context.Context, rid uint64) (err error) {
	err = roomRepository.DB.Debug().Model(&models.Room{}).Where("id = ? ", rid).Take(&models.Room{}).Delete(&models.Room{}).Error
	if err != nil {
		log.Fatalln("Error when delete room from database", err)
		return err
	}
	return nil
}

func (roomRepository *gormRepository) GetRooms(ctx context.Context) (rooms []models.Room, err error) {
	err = roomRepository.DB.Debug().Model(&models.Room{}).Limit(100).Find(&rooms).Error
	if err != nil {
		log.Fatalln("Error when get rooms from database", err)
		return nil, err
	}
	return rooms, nil
}
