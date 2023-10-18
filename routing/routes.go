package routing

import (
	"com.test.users_api_test/app"
	"com.test.users_api_test/handler"
	"github.com/gin-gonic/gin"
)

func InitNonTokenRoutes(app *app.App, userHandler *handler.UserHandler) {
	app.Router.POST("/createUser", userHandler.CreateUserAPiHandler)
	// r.POST("/signin", userHandler.HandlerSignInAPiService)

}

func InitTokenRoutes(router *gin.RouterGroup) {
	// UsersAPIRoutes(router, db)
	// PostsAPIRoutes(router, db)
}

// func BrowserRoutes(appRouting config.AppRouter) {
// appRouting.Router.HandleFunc("/", handler.Welcome)
// router.HandleFunc("/signin", signin)
// router.HandleFunc("/signup", handler.Signup)
// }

// func SignUpRoutes(appRouting config.AppRouter) {
// 	ctxt := config.AppContext{}
// 	appRouting.Router.POST("/createUser", func(c *gin.Context) {
// 		ctxt.Context = c
// 		db.HandlerCreateUserFromAPi(&ctxt)
// 	})
// }

// func SignInRoutes(appRouting config.AppRouter) {
// 	ctxt := config.AppContext{}
// 	appRouting.Router.POST("/signin", func(c *gin.Context) {
// 		ctxt.Context = c
// 		config.DB.HandlerSignInAPi(&ctxt)
// 	})

// 	// appRouting.Router.Post("/createUser", db.HandlerCreateUserFromAPi)
// 	// router.Post("/signin", db.HandlerSignInAPi)
// }

// func UsersAPIRoutes(appRouting AppRouting, db *handler.GormDB) {

// 	router.Get("/users", db.HandlerGetAllUsers)
// 	router.Get("/users/{id:[0-9]+}", db.HandlerGetUserById)
// 	// router.HandleFunc("/users/{id}", db.HandlerGetUserById)
// 	router.Patch("/users/{id:[0-9]+}", db.HandlerEditUser)
// 	router.Delete("/user/{id:[0-9]+}", db.HandlerDeleteUser)
// }

// func PostsAPIRoutes(router *chi.Mux, db *handler.GormDB) {
// 	router.Post("/createPost", db.HandlerCreatePost)
// }
