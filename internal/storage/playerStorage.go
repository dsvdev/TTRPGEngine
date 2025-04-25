package storage

import (
	"TRPGEngine/config"
	"TRPGEngine/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

const PlayerCollection = "Player"

var playerCollection *mongo.Collection

var onceInitPlayerCollection sync.Once

func GetPlayer(id int64) (*model.Player, error) {
	initPlayerCollection()
	var player model.Player
	err := playerCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&player)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err // Другая ошибка
		}
		defState, err := DefaultState()
		if err != nil {
			return nil, err
		}

		player.Id = id
		player.State = defState.Id

		_, err = playerCollection.InsertOne(context.Background(), player)
		if err != nil {
			return nil, err
		}
	}
	return &player, nil
}

func SavePlayer(player *model.Player) (*model.Player, error) {
	initPlayerCollection()
	_, err := playerCollection.ReplaceOne(context.Background(), bson.M{"_id": player.Id}, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func initPlayerCollection() {
	onceInitPlayerCollection.Do(func() {
		cfg := config.Load().MongoConfig
		client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.Uri))
		playerCollection = client.Database(cfg.DbName).Collection(PlayerCollection)
	})
}
