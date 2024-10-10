package main

import (
	"exp/config"
	"exp/internal/db"
	postHandler "exp/internal/handler/post_handler"
	"exp/internal/handler/user_handler"
	"exp/internal/middlewares"
	"exp/internal/repository/post_repo"
	"exp/internal/repository/user_repo"
	"exp/internal/usecase/post_usecase"
	"exp/internal/usecase/user_usecase"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.Config{
		RedisAddr:     "localhost:6379",
		RedisPassword: "",
		RedisDB:       0,
	}
	redis := db.RedisInit(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	if redis == nil {
		log.Fatal("Ошибка инициализации Redis")
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Ошибка подключения к PostgreSQL: %v", err)
	}
	//db.RedisAddData()

	userRepoRouter := user_repo.New(database, redis)
	userUC := user_usecase.New(userRepoRouter, userRepoRouter)
	userH := user_handler.New(userUC)

	//post_handler
	postRepoRouter := post_repo.New(database)
	postUC := post_usecase.New(postRepoRouter)
	postH := postHandler.New(postUC)
	router := gin.Default()

	v0 := router.Group("user")
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
	v1 := router.Group("/post")
	{
		v1.PATCH("/:id", postH.UpdatePost)
		v1.GET("/:id", postH.GetPost)
		v1.POST("/post", postH.CreatePost)
		v1.DELETE("/:id", postH.DeletePost)
	}
	router.Run(":8080")
}
