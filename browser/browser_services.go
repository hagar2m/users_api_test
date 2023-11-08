package services

import (
	"fmt"
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("error->", err.Error())
	}
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("register.html"))
	tmpl.Execute(writer, nil)
}

// func signin(writer http.ResponseWriter, request *http.Request) {
// 	username := request.FormValue("exampleInputEmail1")
// 	password := request.FormValue("exampleInputPassword1")

// 	passwordFromDB := ""
// 	query := `SELECT password FROM users WHERE username = ?`
// 	err := db.QueryRow(query, username).Scan(&passwordFromDB)
// 	print(err)
// 	if password == passwordFromDB {
// 		fmt.Fprintf(writer, "Congratulations "+username+" You are successfully signed in.")
// 	} else {
// 		fmt.Fprintf(writer, "Oops! Username and password did not match.")
// 	}
// }
