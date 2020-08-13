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

func (roomRepository *gormRepository) GetRoomByID(ctx context.Context, rid uint64) (room models.Room, err error) {
	db := roomRepository.DB.Debug().Model(&models.Room{}).Where("id = ?", rid).Take(&room)
	if db.Error != nil {
		log.Printf("[Error when get room from database][%v]", db.Error)
		return room, db.Error
	}

	return room, nil
}

func (roomRepository *gormRepository) SaveRoom(ctx context.Context, room models.Room) (err error) {
	err = roomRepository.DB.Debug().Model(&models.Room{}).Create(&room).Error
	if err != nil {
		log.Printf("[Error when save room to database][%v]", err)
		return err
	}
	return nil
}
func (roomRepository *gormRepository) UpdateRoom(ctx context.Context, room models.Room) (newRoom models.Room, err error) {
	err = roomRepository.DB.Debug().Model(&models.Room{}).Update(&room).Error
	if err != nil {
		log.Printf("[Error when update room to database][%v]", err)
		return room, err
	}

	return room, nil
}

func (roomRepository *gormRepository) DeleteRoom(ctx context.Context, rid uint64) (err error) {

	db := roomRepository.DB.Debug().Model(&models.Room{}).Where("id = ? ", rid).Take(&models.Room{}).Delete(&models.Room{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			log.Printf("[Error when delete room from database][%v]", db.Error)
			return db.Error
		}
		return db.Error
	}

	return nil
}

func (roomRepository *gormRepository) GetRooms(ctx context.Context) (rooms []models.Room, err error) {

	err = roomRepository.DB.Debug().Model(&models.Room{}).Limit(100).Find(&rooms).Error
	if err != nil {
		log.Printf("[Error when get rooms from database][%v]", err)
		return nil, err
	}
	return rooms, nil
}
