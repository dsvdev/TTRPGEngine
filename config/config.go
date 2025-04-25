package config

import (
	"os"
	"sync"
)

type Config struct {
	MongoConfig MongoConfig
	BotConfig   BotConfig
}

type BotConfig struct {
	Token string
}

type MongoConfig struct {
	Uri    string
	DbName string
}

var cfg Config

var once sync.Once

func Load() *Config {
	once.Do(func() {
		cfg = Config{
			MongoConfig: MongoConfig{
				Uri:    os.Getenv("MONGO_URI"),
				DbName: os.Getenv("MONGO_DATABASE"),
			},
			BotConfig: BotConfig{Token: os.Getenv("BOT_TOKEN")},
		}
	})

	return &cfg
}
