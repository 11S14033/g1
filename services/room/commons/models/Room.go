package models

import (
	"time"
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
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
