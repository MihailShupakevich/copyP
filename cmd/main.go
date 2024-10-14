package main

import (
	"exp/config"
	"exp/internal/db"
	postHandler "exp/internal/handler/post_handler"
	"exp/internal/handler/user_handler"
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

	userRepoRouter := user_repo.New(database, redis)
	userUC := user_usecase.New(userRepoRouter, userRepoRouter)
	userH := user_handler.New(userUC)

	//post_handler
	postRepoRouter := post_repo.New(database)
	postUC := post_usecase.New(postRepoRouter)
	postH := postHandler.New(postUC)
	router := gin.Default()

	userGroup := router.Group("/user")
	userH.SetupRoutes(userGroup)

	postGroup := router.Group("/post")
	postH.SetupRoutes(postGroup)

	router.Run(":8080")

}
