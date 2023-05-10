package botvk

import (
	"context"

	"github.com/NoireHub/VKBotSP/internal/store"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/sirupsen/logrus"
)

type Server struct {
	apiVK *api.VK
	store store.Store
	logger *logrus.Logger
	longPoll *longpoll.LongPoll
}

func (s *Server) NewServer(api *api.VK, store store.Store) {

}

func (s *Server) ConfigureLongPoll() {
	s.longPoll.MessageNew(s.MessageNewHandler)
}


func (s *Server) MessageNewHandler(_ context.Context, obj events.MessageNewObject) {



}

