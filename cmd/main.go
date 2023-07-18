package main

import (
	"fmt"
	"jul14/controller"
	"log"
	"net/http"
	"os"
)

func main() {
	server := controller.Server{}
	server.Connect()

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%v", os.Getenv("HOST_PORT")),
			server.Router,
		),
	)

}
