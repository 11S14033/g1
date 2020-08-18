package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/11s14033/g1/services/room/commons/models"
	utils "github.com/11s14033/g1/services/room/delivery/utils"
	"github.com/gorilla/mux"
)

func (rs RoomRestService) GetRoomByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}
	room, err := rs.usecase.GetRoomByID(context.Background(), rid)

	if err != nil {
		log.Printf("[Error][roomRestService][GetRoomByID]-[when calling][roomUseCase]-[GetRoomByID]-[%v]", err)
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, room)

}
func (rs RoomRestService) SaveRoom(w http.ResponseWriter, r *http.Request) {
	var err error
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}
	room := models.Room{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = json.Unmarshal(body, &room)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = rs.usecase.SaveRoom(context.Background(), room)
	if err != nil {
		log.Printf("[Error][roomRestService][SaveRoom]-[when calling][roomUseCase]-[SaveRoom]-[%v]", err)
		formattedError := utils.FormatError(err.Error())
		utils.ErrorJSON(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, room.ID))
	utils.JSON(w, http.StatusCreated, room)
}
func (rs RoomRestService) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	room := models.Room{
		UpdatedAt: time.Now(),
	}

	err = json.Unmarshal(body, &room)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	newRoom, err := rs.usecase.UpdateRoom(context.Background(), room)
	if err != nil {
		log.Printf("[Error][roomRestService][UpdateRoom]-[when calling][roomUseCase]-[UpdateRoom]-[%v]", err)
		utils.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSON(w, http.StatusCreated, newRoom)
}
func (rs RoomRestService) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	rid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}
	err = rs.usecase.DeleteRoom(context.Background(), rid)
	if err != nil {
		log.Printf("[Error][roomRestService][DeleteRoom]-[when calling][roomUseCase]-[DeleteRoom]-[%v]", err)
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", rid))
	utils.JSON(w, http.StatusNoContent, "")
}
func (rs RoomRestService) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := rs.usecase.GetRooms(context.Background())
	if err != nil {
		log.Printf("[Error][roomRestService][GetRooms]-[when calling][roomUseCase]-[GetRooms]-[%v]", err)
		utils.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSON(w, http.StatusCreated, rooms)
}
