package main

import (
	"exp/internal/db"
	postHandler "exp/internal/handler/post"
	user3 "exp/internal/handler/user"
	"exp/internal/middlewares"
	post2 "exp/internal/repository/post"
	user2 "exp/internal/repository/user"
	"exp/internal/usecase/post"
	"exp/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

func main() {

	database, _ := db.Connect()

	userRepoRouter := user2.New(database)
	userUC := user.New(userRepoRouter)
	userH := user3.New(userUC)

	//post
	postRepoRouter := post2.New(database)
	postUC := post.New(postRepoRouter)
	postH := postHandler.New(postUC)
	router := gin.Default()

	v0 := router.Group("/user")
	{

		v0.GET("/", userH.FindUsers)
		v0.POST("/users", userH.CreateUsers)
		v0.GET("/:id", userH.FindUser)
		v0.DELETE("/:id", userH.DeleteUser)
		v0.PATCH("/:id", userH.UpdateUser)
		v0.POST("/register", userH.Registration)
		v0.GET("/login", middlewares.JwtMiddleware(), userH.Login)
	}

	//post router
	v1 := router.Group("/post")
	{
		v1.PATCH("/:id", postH.UpdatePost)
		v1.GET("/:id", postH.GetPost)
		v1.POST("/post", postH.CreatePost)
		v1.DELETE("/:id", postH.DeletePost)
	}
	router.Run(":8080")
}
