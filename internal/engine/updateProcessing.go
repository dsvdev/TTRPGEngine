package engine

import (
	"TRPGEngine/internal/model"
	"TRPGEngine/internal/storage"
	"github.com/dsvdev/telego/pkg/common"
	"github.com/dsvdev/telego/pkg/common/sending"
)

func ProcessUpdate(inputUpdate *common.Message, outbox chan sending.TelegramSendable) {
	PlayerID := inputUpdate.ChatID
	if inputUpdate.PhotoID != "" {
		outbox <- &sending.SendPhotoById{
			ChatID:  inputUpdate.ChatID,
			Text:    inputUpdate.PhotoID,
			PhotoID: inputUpdate.PhotoID,
		}
		return
	}

	player, err := storage.GetPlayer(PlayerID)
	if err != nil {
		outbox <- &sending.SendMessage{
			ChatID: inputUpdate.ChatID,
			Text:   err.Error(),
		}
		return
	}

	state, err := storage.StateById(player.State)
	if err != nil {
		outbox <- &sending.SendMessage{
			ChatID: inputUpdate.ChatID,
			Text:   err.Error(),
		}
	}
	transitions := state.Transitions
	newStateId, ok := transitions[inputUpdate.Text]

	defState, err := storage.DefaultState()
	if err != nil {
		outbox <- &sending.SendMessage{
			ChatID: inputUpdate.ChatID,
			Text:   err.Error(),
		}
	}

	if ok == false && state.Id != defState.Id {
		outbox <- &sending.SendMessage{
			ChatID: inputUpdate.ChatID,
			Text:   "Пользуйтесь кнопками",
			Keyboard: sending.SendKeyboard{
				Type:    sending.ReplyKeyboard,
				Buttons: transitionsToButtons(transitions),
			},
		}
		return
	}

	player.State = newStateId

	_, err = storage.SavePlayer(player)
	if err != nil {
		outbox <- &sending.SendMessage{
			ChatID: inputUpdate.ChatID,
			Text:   err.Error(),
		}
		return
	}

	newState, err := storage.StateById(player.State)
	if err != nil {
		outbox <- &sending.SendMessage{
			ChatID: inputUpdate.ChatID,
			Text:   err.Error(),
		}
	}

	keyboard := sending.SendKeyboard{
		Type:    sending.ReplyKeyboard,
		Buttons: transitionsToButtons(newState.Transitions),
	}

	if newState.Image == "" {
		outbox <- &sending.SendMessage{
			ChatID:   inputUpdate.ChatID,
			Text:     newState.Text,
			Keyboard: keyboard,
		}
	} else {
		outbox <- &sending.SendPhotoById{
			ChatID:   inputUpdate.ChatID,
			Text:     newState.Text,
			PhotoID:  newState.Image,
			Keyboard: keyboard,
		}
	}
}

func transitionsToButtons(stateByTransition map[string]model.StateID) [][]string {
	if stateByTransition == nil {
		return nil
	}

	result := make([][]string, len(stateByTransition))
	i := 0
	for key := range stateByTransition {
		result[i] = make([]string, 1)
		result[i][0] = key
		i++
	}
	return result
}
