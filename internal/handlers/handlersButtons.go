package handlers

import (
	"math/rand"

	"github.com/NoireHub/VKBotSP/internal/model"
	"github.com/NoireHub/VKBotSP/internal/store"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
)

func HandleDefault(obj events.MessageNewObject) api.Params {
	p := params.NewMessagesSendBuilder()
	reqParams(p, obj)

	p.Message(msgUnknownCommand)
	return p.Params
}

func HandleStart(obj events.MessageNewObject) api.Params {
	p := params.NewMessagesSendBuilder()
	reqParams(p, obj)

	p.Keyboard(model.NewMainMenuKeyboard())
	p.Message(msgHello)
	return p.Params
}

func HandleReadyToSaveQuote(obj events.MessageNewObject) api.Params {
	p := params.NewMessagesSendBuilder()
	reqParams(p, obj)

	p.Message(msgReadyToSave)
	return p.Params
}

func HandleSaveQuote(obj events.MessageNewObject, store store.Store) api.Params {
	p := params.NewMessagesSendBuilder()
	reqParams(p, obj)

	if err := store.Quote().Create(&model.Quote{
		Text:   obj.Message.Text,
		PeerID: obj.Message.PeerID,
	}); err != nil {
		if err.Error() == "такая цитата уже существует" {
			p.Message(msgAlreadyExists)
			return p.Params
		}

		p.Message(msgError)
		return p.Params
	}

	p.Message(msgSaved)
	return p.Params
}

func HandleGetRandomQuote(obj events.MessageNewObject, store store.Store) api.Params {
	p := params.NewMessagesSendBuilder()
	reqParams(p, obj)

	quotes, err := store.Quote().FindById(obj.Message.PeerID)
	if err != nil {
		p.Message(msgError)
		return p.Params
	}

	if len(quotes) == 0 {
		p.Message(msgNoSavedQuotes)
		return p.Params
	}

	p.Message(quotes[rand.Intn(len(quotes))].Text + "\n\n © Джейсон Стейтем")
	return p.Params
}

func HandleAboutBot(obj events.MessageNewObject) api.Params {
	p := params.NewMessagesSendBuilder()
	reqParams(p, obj)

	p.Message(msgHelp)
	return p.Params
}

func HandleRoll(obj events.MessageNewObject) api.Params {
	p := params.NewMessagesSendBuilder()
	reqParams(p, obj)

	rand := rand.Intn(101)
	if rand > 99 {
		p.Message(msgSuperLucky)
		return p.Params
	}
	if rand > 50 {
		p.Message(msgLucky)
		return p.Params
	}

	p.Message(msgNoLucky)
	return p.Params
}

func HandleRandomProphesy(obj events.MessageNewObject) api.Params {
	p := params.NewMessagesSendBuilder()
	reqParams(p, obj)

	p.Message(Prophesy[rand.Intn(len(Prophesy))])
	return p.Params
}
