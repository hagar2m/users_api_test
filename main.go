package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"com.test.users_api_test/handler"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("can't connect to database")
	}
	gormDb := handler.GormDB{
		DB: db,
	}

	conn, err := sql.Open("mysql", dbURL)
	apiConfig := handler.APIConfig{
		DB: conn,
	}
	router := chi.NewRouter()

	defer func() {
		srv := &http.Server{
			Handler: router,
			Addr:    ":" + portString,
		}
		log.Printf("Server starting on port %v", portString)
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
		conn.Close()
		// defer db.Close()
	}()
	// Brower //
	router.HandleFunc("/", welcome)
	// router.HandleFunc("/signin", signin)
	router.HandleFunc("/signup", signup)
	router.HandleFunc("/createUserForm", apiConfig.HandlerCreateUserFromBrowser)

	// Apis
	router.Post("/createUser", gormDb.HandlerCreateUserFromAPi)
	router.Get("/users", gormDb.HandlerGetAllUsers)
	router.Get("/users/{id:[0-9]+}", gormDb.HandlerGetUserById)
	// router.Get("/users/{id}", gormDb.HandlerGetUserById)
	router.Patch("/users/{id:[0-9]+}", gormDb.HandlerEditUser)
	router.Delete("/user/{id:[0-9]+}", gormDb.HandlerDeleteUser)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("error->", err.Error())
	}
}
func signup(writer http.ResponseWriter, request *http.Request) {
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
