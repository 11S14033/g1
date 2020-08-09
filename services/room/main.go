package main

import (
	"log"
)

func main() {
	roomService := RoomService{}
	err := roomService.StartService()
	if err != nil {
		log.Fatalln("Cannont start server cause: ", err)
	}
}
