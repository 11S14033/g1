package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Room struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	RoomName    string    `gorm:"size:255;not null;unique" json:"room_name"`
	Owner       string    `gorm:"size:255;not null" json:"owner"`
	Address     string    `gorm:"not null" json:"address"`
	City        string    `gorm:"not null" json:"city"`
	Province    string    `gorm:"not null" json:"province"`
	Description string    `gorm:"not null" json:"description"`
	Type        uint32    `gorm:"not null" json:"type"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}

func (r *Room) prepare() {
	r.ID = 0
	r.RoomName = html.EscapeString(strings.TrimSpace(r.RoomName))
	r.Owner = html.EscapeString(strings.TrimSpace(r.Owner))
	r.Address = html.EscapeString(strings.TrimSpace(r.Address))
	r.City = html.EscapeString(strings.TrimSpace(r.City))
	r.Province = html.EscapeString(strings.TrimSpace(r.Province))
	r.Description = html.EscapeString(strings.TrimSpace(r.Description))
	r.Type = 0
	r.CreatedAt = time.Now()
	r.UpdateAt = time.Now()

}

func (r *Room) validate() error {
	if r.RoomName == "" {
		return errors.New("Requeired room name")
	}
	if r.Owner == "" {
		return errors.New("Requeired owner")
	}
	if r.Address == "" {
		return errors.New("Requeired address")
	}
	if r.City == "" {
		return errors.New("Requeired city")
	}
	if r.Province == "" {
		return errors.New("Requeired province")
	}
	if r.Description == "" {
		return errors.New("Requeired description")
	}
	if r.Type < 0 {
		return errors.New("Requeired type")
	}

	return nil

}

func (r *Room) SaveRoom(db *gorm.DB) (*Room, error) {
	var err error

	//gorm save
	err = db.Debug().Model(&Room{}).Create(&r).Error
	if err != nil {
		return &Room{}, err
	}

	return r, nil
}

func (r *Room) FindAllRoom(db *gorm.DB) (*[]Room, error) {
	var err error
	rooms := []Room{}

	//gorn find Limit
	err = db.Debug().Model(&Room{}).Limit(100).Find(&rooms).Error
	if err != nil {
		return &[]Room{}, err
	}
	return &rooms, nil
}

func (r *Room) FindRoomByID(db *gorm.DB, rid uint64) (*Room, error) {
	var err error

	//gorm find
	err = db.Debug().Model(&Room{}).Where("id = ?", rid).Take(&r).Error
	if err != nil {
		return &Room{}, err
	}
	return r, nil
}

func (r *Room) UpdatePost(db *gorm.DB, rid int64) (*Room, error) {
	var err error

	//attribute to be update
	attribute := Room{
		RoomName:    r.RoomName,
		Address:     r.Address,
		City:        r.City,
		Province:    r.Province,
		Description: r.Description,
		Type:        r.Type,
		UpdateAt:    time.Now(),
	}
	//gorm update
	err = db.Debug().Model(&Room{}).Where("id = ?", r.ID).Update(attribute).Error
	if err != nil {
		return &Room{}, err
	}

	return r, nil
}
func (r *Room) DeletePost(db *gorm.DB, rid uint64) (int64, error) {
	var err error

	//gorm delete
	result := db.Debug().Model(&Room{}).Where("id = ? ").Take(&Room{}).Delete(&Room{})

	//handling error
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("Room not found")
		}
		return 0, err
	}
	return result.RowsAffected, nil
}
