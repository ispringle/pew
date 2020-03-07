package pew

import (
	"log"
	"net/http"
)

func Run() {
	address := ":8080"
	var storage AlertStorage

	http.HandleFunc("/alert", storage.AlertListener)
	log.Fatal(http.ListenAndServe(address, nil))
}
