package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Player struct {
	Name   			string `json:"name" binding:"required"`
	Age 			int `json:"age" binding:"required"`
	CreationTime 	time.Time `json:"creation_time"`
}

func main() {
	
	ginEngine:= gin.Default()

	ginEngine.POST("/players", func(c *gin.Context){
		var player Player
		if err := c.BindJSON(&player); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		player.CreationTime = time.Now().UTC()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		collection := client.Database("ChampionshipDB").Collection("players")
		insertResult, err := collection.InsertOne(ctx, player)

		if err != nil {
			log.Fatal(err)
		}
		
		c.JSON(200, gin.H{
			"player_id": insertResult.InsertedID,
		})
	})
	
	log.Fatalln(ginEngine.Run(":8081"))
}