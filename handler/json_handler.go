package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		log.Println("Responding with 5xx error", msg)
	}

	type errResponse struct { 
		Error string `json:"error"`
	}

	responseWithJson(w, code, errResponse{
		Error: msg,
	})
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("faild to marshal %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
