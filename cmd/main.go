package main

import (
	"jul14/controller"
	"log"
	"net/http"
)

func main() {
	server := controller.Server{}
	server.Connect()
	log.Fatal(http.ListenAndServe(":8080", server.Router))

}
