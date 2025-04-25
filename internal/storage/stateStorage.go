package storage

import (
	"TRPGEngine/config"
	"TRPGEngine/internal/model"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

const StateCollection = "State"

var stateCollection *mongo.Collection

var onceInitStateCollection sync.Once

func DefaultState() (*model.PlayerState, error) {
	return StateById("default")
}

func SaveState(state *model.PlayerState) (*model.PlayerState, error) {
	initStateCollection()
	_, err := stateCollection.InsertOne(context.Background(), state)
	if err != nil {
		return nil, err
	}
	return state, nil
}

func StateById(id model.StateID) (*model.PlayerState, error) {
	initStateCollection()
	var state model.PlayerState
	err := stateCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&state)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err // Другая ошибка
		}
		return DefaultState()
	}
	return &state, nil
}

func initStateCollection() {
	onceInitStateCollection.Do(func() {
		cfg := config.Load().MongoConfig
		client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.Uri))
		stateCollection = client.Database(cfg.DbName).Collection(StateCollection)
	})
}
