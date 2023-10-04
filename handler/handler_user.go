package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (apiCfg *APIConfig) HandlerCreateUserFromBrowser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("exampleInputEmail1")
	password := r.FormValue("exampleInputPassword1")
	name := r.FormValue("exampleInputName1")

	result, error := apiCfg.DB.Exec(`INSERT INTO users (name, email, password, created_at) VALUES (?, ?, ?,?)`, name, email, password, time.Now())
	if error != nil {
		log.Fatal("Error", error)
		fmt.Fprintf(w, "Oops! try again later.")
	} else {
		fmt.Fprintf(w, "result = %v", result)
		fmt.Fprintf(w, "Congratulations "+name+" You are successfully regsitered.")
	}
}

func (apiCfg *APIConfig) HandlerCreateUserFromAPi(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&params)
	fmt.Printf("body is: %v\n", params)
	if error != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", error))
		return
	}

	result, error := apiCfg.DB.Exec(`INSERT INTO users (name, email, password, created_at) VALUES (?, ?, ?,?)`, params.Name, params.Email, params.Password, time.Now())
	if error != nil {
		log.Fatal("Error", error)
		fmt.Fprintf(w, "Oops! try again later.")
	} else {
		fmt.Fprintf(w, "result = %v", result)
		fmt.Fprintf(w, "Congratulations "+params.Name+" You are successfully regsitered.")
	}
	//responseWithJson(w, http.StatusOK, struct{}{})
}

