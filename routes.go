package main

import (
	"com.test.users_api_test/handler"
	"github.com/go-chi/chi"
)

func SetupRoutes(router *chi.Mux, db *handler.GormDB) {
	BrowserRoutes(router, db)
	UsersAPIRoutes(router, db)
	PostsAPIRoutes(router, db)
}

func BrowserRoutes(router *chi.Mux, db *handler.GormDB) {
	router.HandleFunc("/", handler.Welcome)
	// router.HandleFunc("/signin", signin)
	router.HandleFunc("/signup", handler.Signup)
}

func UsersAPIRoutes(router *chi.Mux, db *handler.GormDB) {
	router.Post("/createUser", db.HandlerCreateUserFromAPi)
	router.Get("/users", db.HandlerGetAllUsers)
	router.Get("/users/{id:[0-9]+}", db.HandlerGetUserById)
	// router.HandleFunc("/users/{id}", db.HandlerGetUserById)
	router.Patch("/users/{id:[0-9]+}", db.HandlerEditUser)
	router.Delete("/user/{id:[0-9]+}", db.HandlerDeleteUser)
}

func PostsAPIRoutes(router *chi.Mux, db *handler.GormDB) {
	router.Post("/createPost", db.HandlerCreatePost)
}
