package player

import (
	"context"
	"log"
	"os"
	"time"

	"go.mod/internal/domain"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (s Service) InsertPlayer(player domain.Player) (id interface{}, err error) {
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

	return insertResult, nil
}
