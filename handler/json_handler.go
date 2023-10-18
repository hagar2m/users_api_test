package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		log.Println("Responding with 5xx error", msg)
	}

	type errResponse struct { 
		Error string `json:"error"`
	}

	ResponseWithJson(w, code, errResponse{
		Error: msg,
	})
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("faild to marshal %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
