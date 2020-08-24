package migrations

import (
	"log"

	"github.com/11s14033/g1/services/room/commons/models"
	"github.com/jinzhu/gorm"
)

var rooms = []models.Room{
	models.Room{
		Address:     "Mampang",
		City:        "Jaksel",
		Description: "Kos",
		Owner:       "Mamang",
		Province:    "DKI",
		RoomName:    "Kamar6",
		Type:        1,
	},
	models.Room{
		Address:     "Mampang",
		City:        "Jaksel",
		Description: "Kos",
		Owner:       "Mamang",
		Province:    "DKI",
		RoomName:    "Kamar7",
		Type:        1,
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Room{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Room{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range rooms {
		err = db.Debug().Model(&models.Room{}).Create(&rooms[i]).Error
		if err != nil {
			log.Fatalf("cannot seed rooms table: %v", err)
		}
	}
}
