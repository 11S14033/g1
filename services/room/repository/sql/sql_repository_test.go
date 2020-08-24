package sql_test

import (
	"context"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/11s14033/g1/services/room/commons/models"
	reposql "github.com/11s14033/g1/services/room/repository/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

type fields struct {
	DB    *gorm.DB
	Room  models.Room
	Rooms []models.Room
	Query string
}

func TestGetRoomByID(t *testing.T) {
	tMock := time.Now()

	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Error("Failed to mock db")
	}

	pg, err := gorm.Open("postgres", mockDb)
	pg.LogMode(true)
	if err != nil {
		t.Error("Failed to mock posgre db")
	}

	sql := `SELECT * FROM "rooms" WHERE (id = $1) LIMIT 1`
	rgxQuery := regexp.QuoteMeta(sql)

	successFields := fields{
		DB: pg,
		Room: models.Room{
			ID:          1,
			RoomName:    "Test",
			Owner:       "Test",
			Address:     "Test",
			City:        "Test",
			Province:    "Test",
			Description: "Test",
			Type:        1,
			CreatedAt:   tMock,
			UpdatedAt:   tMock,
		},
		Query: rgxQuery,
	}

	tests := []struct {
		Name   string
		Fields fields
		Want   []models.Room
	}{
		{
			Name:   "success-test",
			Fields: successFields,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := reposql.NewGormRepository(test.Fields.DB)

			rows := sqlmock.NewRows([]string{"id", "room_name", "owner", "address", "city", "province", "description", "type", "created_at", "updated_at"}).
				AddRow(1, "Test", "Test", "Test", "Test", "Test", "Test", 1, tMock, tMock)

			mock.ExpectQuery(test.Fields.Query).WillReturnRows(rows)

			if got, _ := repo.GetRoomByID(context.Background(), 5); !reflect.DeepEqual(got, test.Fields.Room) {
				t.Errorf("Room_repository.GetRoomByID(rid) = %v, want %v", got, test.Fields.Room)
			}
		})
	}
}

func TestGetRooms(t *testing.T) {
	tMock := time.Now()

	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Error("Failed to mock db")
	}

	pg, err := gorm.Open("postgres", mockDb)
	pg.LogMode(true)
	if err != nil {
		t.Error("Failed to mock posgre db")
	}

	sql := `SELECT * FROM "rooms" `
	rgxQuery := regexp.QuoteMeta(sql)

	successFields := fields{
		DB: pg,
		Rooms: []models.Room{
			models.Room{
				ID:          1,
				RoomName:    "Test",
				Owner:       "Test",
				Address:     "Test",
				City:        "Test",
				Province:    "Test",
				Description: "Test",
				Type:        1,
				CreatedAt:   tMock,
				UpdatedAt:   tMock,
			},
			models.Room{
				ID:          2,
				RoomName:    "Test2",
				Owner:       "Test2",
				Address:     "Test2",
				City:        "Test2",
				Province:    "Test2",
				Description: "Test2",
				Type:        2,
				CreatedAt:   tMock,
				UpdatedAt:   tMock,
			},
		},

		Query: rgxQuery,
	}

	tests := []struct {
		Name   string
		Fields fields
		Want   []models.Room
	}{
		{
			Name:   "success-test",
			Fields: successFields,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			repo := reposql.NewGormRepository(test.Fields.DB)

			rows := sqlmock.NewRows([]string{"id", "room_name", "owner", "address", "city", "province", "description", "type", "created_at", "updated_at"}).
				AddRow(1, "Test", "Test", "Test", "Test", "Test", "Test", 1, tMock, tMock).
				AddRow(2, "Test2", "Test2", "Test2", "Test2", "Test2", "Test2", 2, tMock, tMock)

			mock.ExpectQuery(test.Fields.Query).WillReturnRows(rows)

			if got, _ := repo.GetRooms(context.Background()); !reflect.DeepEqual(got, test.Fields.Rooms) {
				t.Errorf("Room_repository.GetRooms = %v, want %v", got, test.Fields.Rooms)
			}
		})
	}

}
