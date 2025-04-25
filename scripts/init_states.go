package scripts

import (
	"TRPGEngine/internal/model"
	"TRPGEngine/internal/storage"
	"log"
)

var states = []model.PlayerState{
	{
		Id:    "default",
		Text:  "Добро пожаловать в игру!\nВыберите персонажа",
		Image: "AgACAgIAAxkBAAMXaAtgysIhWuHPI55PzEcQr7q-4IMAAovtMRvKeWBI66LIb60HYOYBAAMCAANzAAM2BA",
		Transitions: map[string]model.StateID{
			"Капитан": model.CaptainStateID,
			"Медик":   model.MedicStateID,
		},
	},

	{
		Id:    model.MedicStateID,
		Text:  "Вы медик!",
		Image: "AgACAgIAAxkBAAMjaAtjcC2gZQqZ-jnip8nTUkYlFu0AAsntMRvKeWBIWv3NMAgsfJEBAAMCAANzAAM2BA",
		Transitions: map[string]model.StateID{
			"Назад": model.DefaultStateID,
		},
	},

	{
		Id:    model.CaptainStateID,
		Text:  "Вы капитан!",
		Image: "AgACAgIAAxkBAAMfaAtjI1XPkEMgv7eatepGugs3mHEAAsTtMRvKeWBIPqSwILcH5IcBAAMCAANzAAM2BA",
		Transitions: map[string]model.StateID{
			"Назад": model.DefaultStateID,
		},
	},
}

func InitStateDb() {
	for _, state := range states {
		_, err := storage.SaveState(&state)
		if err != nil {
			log.Printf("%v", err)
		}
	}
}
