package main

import (
	"log"

	"github.com/pdhawan2001/Go-REST-API/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
