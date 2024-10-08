package main

import (
	"exp/internal/db"
	postHandler "exp/internal/handler/post_handler"
	user3 "exp/internal/handler/user_handler"
	"exp/internal/middlewares"
	post2 "exp/internal/repository/post_repo"
	user2 "exp/internal/repository/user_repo"
	"exp/internal/usecase/post_usecase"
	"exp/internal/usecase/user_usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	database, _ := db.Connect()

	userRepoRouter := user2.New(database)
	userUC := user_usecase.New(userRepoRouter)
	userH := user3.New(userUC)

	//post_handler
	postRepoRouter := post2.New(database)
	postUC := post_usecase.New(postRepoRouter)
	postH := postHandler.New(postUC)
	router := gin.Default()

	v0 := router.Group("/user_repo")
	{

		v0.GET("/", userH.FindUsers)
		v0.POST("/users", userH.CreateUsers)
		v0.GET("/:id", userH.FindUser)
		v0.DELETE("/:id", userH.DeleteUser)
		v0.PATCH("/:id", userH.UpdateUser)
		v0.POST("/register", userH.Registration)
		v0.GET("/login", middlewares.JwtMiddleware(), userH.Login)
	}

	//post_handler router
	v1 := router.Group("/post_handler")
	{
		v1.PATCH("/:id", postH.UpdatePost)
		v1.GET("/:id", postH.GetPost)
		v1.POST("/post_handler", postH.CreatePost)
		v1.DELETE("/:id", postH.DeletePost)
	}
	router.Run(":8080")
}
