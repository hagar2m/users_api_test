package routing

import (
	"com.test.users_api_test/user"
	"github.com/gin-gonic/gin"
)

func NonTokenRoutes(r *gin.RouterGroup) {
	r.POST("/createUser", user.CreateUserAPiHandler)
	r.POST("/signin", user.SingInAPiHandler)

}

func TokenRoutes(r *gin.RouterGroup) {
	r.GET("/users", user.GetAllUsersHandler)
	r.GET("/users/:id", user.GetUserByIdhandler)
	r.PATCH("/users/:id", user.EditUserhandler)
}

// func BrowserRoutes(appRouting config.AppRouter) {
// appRouting.Router.HandleFunc("/", handler.Welcome)
// router.HandleFunc("/signin", signin)
// router.HandleFunc("/signup", handler.Signup)
// }

// func UsersAPIRoutes(appRouting AppRouting, db *handler.GormDB) {

// 	router.Get("/users", db.HandlerGetAllUsers)
// 	router.Get("/users/{id:[0-9]+}", db.HandlerGetUserById)
// 	router.Patch("/users/{id:[0-9]+}", db.HandlerEditUser)
// 	router.Delete("/user/{id:[0-9]+}", db.HandlerDeleteUser)
// }

// func PostsAPIRoutes(router *chi.Mux, db *handler.GormDB) {
// 	router.Post("/createPost", db.HandlerCreatePost)
// }
