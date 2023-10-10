package main

import (
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
		// defer db.Close()
	}()

	// Pass the router and database connection to SetupRoutes
	SetupRoutes(router, &gormDb)
}


