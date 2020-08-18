package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FormatError(err string) error {

	if strings.Contains(err, "rooms_pkey") {
		return errors.New("ID Already Taken")
	}

	if strings.Contains(err, "room_name") {
		return errors.New("RoomName Already Taken")
	}
	return errors.New("Incorrect Details")
}

//Handling error response when mode delivery is REST
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ErrorJSON(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error   string `json:"error"`
			Code    int    `json:"code"`
			Message string `json:""message`
		}{
			Error:   err.Error(),
			Code:    statusCode,
			Message: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}

//Handling error response when mode delivery is gRPC
func ErrorRPC(statusCode codes.Code, desc string) error {
	return status.Errorf(codes.InvalidArgument, desc)
}
