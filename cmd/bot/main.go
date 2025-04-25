package main

import (
	"TRPGEngine/config"
	"TRPGEngine/internal/engine"
	"TRPGEngine/scripts"
	"github.com/dsvdev/telego/pkg/bot"
	"sync"
)

func main() {
	cfg := config.Load()
	scripts.InitStateDb()

	TTVBot := bot.NewLongpollingTelegramBot(cfg.BotConfig.Token)
	TTVBot.StartProcessUpdates(engine.ProcessUpdate)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
