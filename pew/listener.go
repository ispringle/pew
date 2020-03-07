package pew

import (
	"encoding/json"
	"log"
	"net/http"
)

func (storage *AlertStorage) AlertListener(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		storage.getListener(res, req)
	case http.MethodPost:
		storage.postListener(res, req)
	default:
		http.Error(res, "unsupported method, have a 400", 400)
	}
}

func (storage *AlertStorage) getListener(res http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(res)
	res.Header().Set("Content-Type", "application/json")

	storage.Lock()
	defer storage.Unlock()

	if err := encoder.Encode(storage.AlertMessages); err != nil {
		log.Printf("error encoding messages %v", err)
	}

}

func (storage *AlertStorage) postListener(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	var message Message
	if err := decoder.Decode(&message); err != nil {
		log.Printf("error decoding message: %v", err)
		http.Error(res, "invalid request, have a 400", 400)
	}

	storage.Lock()
	defer storage.Unlock()

	storage.AlertMessages = append(storage.AlertMessages, &message)
}
