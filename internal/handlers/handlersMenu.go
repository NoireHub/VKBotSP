package handlers

import (
	"math/rand"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/object"
)

func HandleMenu(obj events.MessageNewObject, kb object.MessagesKeyboard, msg string) api.Params {
	p := params.NewMessagesSendBuilder()
	reqParams(p, obj)

	p.Keyboard(kb)
	p.Message(msg)
	return p.Params
}

func reqParams(p *params.MessagesSendBuilder, obj events.MessageNewObject) {
	p.PeerID(obj.Message.PeerID)
	p.RandomID(int(rand.Int31()))
}
