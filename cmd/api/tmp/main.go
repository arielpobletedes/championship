package main

import (
	"github.com/gin-gonic/gin"
	"go.mod/cmd/api/handlers/player"
	PlayerService "go.mod/internal/services/player"

	"log"
)

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	ginEngine := gin.Default()

	playerService := PlayerService.Service{}
	playerHandler := player.Handler{
		PlayerService: playerService,
	}

	ginEngine.POST("/players", playerHandler.CreatePlayerHandler)

	log.Fatalln(ginEngine.Run(":8081"))
}
