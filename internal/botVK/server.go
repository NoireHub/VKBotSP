package botvk

import (
	"context"
	"fmt"
	"math/rand"
	"os"

	"github.com/NoireHub/VKBotSP/internal/handlers"
	"github.com/NoireHub/VKBotSP/internal/model"
	"github.com/NoireHub/VKBotSP/internal/store"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/sirupsen/logrus"
)

type server struct {
	apiVK    *api.VK
	store    store.Store
	logger   *logrus.Logger
	longPoll *longpoll.LongPoll
	saveQoutesMessageID map[int]int //map[PeerID]ConversationMessageID
}

func NewServer(api *api.VK, store store.Store, groupID int) (*server, error) {
	s := &server{
		store:  store,
		apiVK:  api,
		logger: logrus.New(),
		saveQoutesMessageID: make(map[int]int),
	}

	lp, err := longpoll.NewLongPoll(api, groupID)
	if err != nil {
		return nil, err
	}

	s.longPoll = lp
	s.ConfigureLongPoll()

	return s, nil
}

func (s *server) ConfigureLongPoll() {
	s.longPoll.MessageNew(s.MessageNewHandler)
	//s.longPoll.MessageEvent(s.MessageEventHandler)
}

func (s *server) MessageNewHandler(_ context.Context, obj events.MessageNewObject) {
	var resp int
	var err error

	if s.saveQoutesMessageID[obj.Message.PeerID] == obj.Message.ConversationMessageID {
		params := handlers.HandleSaveQuote(obj,s.store)
		resp, err = s.apiVK.MessagesSend(params)
	}else{
		switch obj.Message.Text {
		case "Начать":
			params := handlers.HandleStart(obj)
			resp, err = s.apiVK.MessagesSend(params)
		case "Об авторе":
			params := handlers.HandleMenu(obj, model.NewAboutAuthorKeyboard(), "Раздел об авторе")
			resp, err = s.apiVK.MessagesSend(params)
		case "О боте":
			params := handlers.HandleMenu(obj, model.NewAboutBotKeyboard(), "Раздел о боте")
			resp, err = s.apiVK.MessagesSend(params)
		case "Цитаты":
			params := handlers.HandleMenu(obj, model.NewQuoteKeyboard(), "Раздел великих цитат")
			resp, err = s.apiVK.MessagesSend(params)
		case "Рандом":
			params := handlers.HandleMenu(obj, model.NewRandomMenuKeyboard(), "Уголок неопределенности")
			resp, err = s.apiVK.MessagesSend(params)
		case "Назад":
			params := handlers.HandleMenu(obj, model.NewMainMenuKeyboard(), "Главное меню")
			resp, err = s.apiVK.MessagesSend(params)
		case "Возможности":
			params := handlers.HandleAboutBot(obj)
			resp, err = s.apiVK.MessagesSend(params)
		case "Покрутить колесо удачи":
			params := handlers.HandleRoll(obj)
			resp, err = s.apiVK.MessagesSend(params)
		case "Получить предсказание":
			params := handlers.HandleRandomProphesy(obj)
			resp, err = s.apiVK.MessagesSend(params)
		case "Добавить цитату":
			s.saveQoutesMessageID[obj.Message.PeerID] = obj.Message.ConversationMessageID + 2
			params := handlers.HandleReadyToSaveQuote(obj)
			resp, err = s.apiVK.MessagesSend(params)
		case "Получить случайную цитату":
			params := handlers.HandleGetRandomQuote(obj, s.store)
			file, _ := os.Open(fmt.Sprintf("./static/stat%d.jpg",rand.Intn(3)))
			response, resError := s.apiVK.UploadMessagesPhoto(obj.Message.PeerID, file)
			if resError == nil {
				params["attachment"] = fmt.Sprintf("photo%d_%d", response[0].OwnerID, response[0].ID)
			}
			resp, err = s.apiVK.MessagesSend(params)
		default:
			params := handlers.HandleDefault(obj)
			resp, err = s.apiVK.MessagesSend(params)
		}
	}	
	
	if err != nil{
		s.logger.Info(fmt.Sprintf("request to user %d failed with response code %d and error: %s",obj.Message.PeerID,resp,err.Error()))
	}else{
		s.logger.Info(fmt.Sprintf("request to user %d completed with response code %d",obj.Message.PeerID,resp))
	}
}
