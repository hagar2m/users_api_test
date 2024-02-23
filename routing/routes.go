package routing

import (
	"com.test.users_api_test/api/comment"
	"com.test.users_api_test/api/post"
	"com.test.users_api_test/api/user"
	"com.test.users_api_test/pkg/auth"
	"github.com/gin-gonic/gin"
)

func NonTokenRoutes(r *gin.RouterGroup) {
	r.POST("/createUser", user.CreateUserHandler)
	r.POST("/signin", user.SingInHandler)
}

func TokenRoutes(r *gin.RouterGroup) {
	r.Use(auth.AuthMiddleware())

	r.GET("/users", user.GetAllUsersHandler)
	r.GET("/users/:id", user.GetUserByIdhandler)
	r.PATCH("/users/:id", user.EditUserhandler)
	r.DELETE("/users/:id", user.DeleteUserHandler)
	r.POST("/createPost", post.CreatePostHandler)
	r.POST("/createComment", comment.CreateCommentHandler)
	r.GET("/comment/:id", comment.GetCommentByIDHandler)
	r.GET("/post/:id", post.GetPostCommentsHandler)
}

// 	router.Get("/users/{id:[0-9]+}", db.HandlerGetUserById
