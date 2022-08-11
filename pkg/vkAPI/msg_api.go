package vkAPI

import (
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
)

func SendSimpleMsg(to int, msg string) {
	b := params.NewMessagesSendBuilder()
	b.Message(msg)
	b.RandomID(0)
	b.PeerID(to)
	_, err := vk.MessagesSend(b.Params)
	if err != nil {
		panic(err)
	}
}

func SendMessageEventAnswer(obj events.MessageEventObject, eventData string) {
	b := params.NewMessagesSendMessageEventAnswerBuilder()
	b.PeerID(obj.PeerID)
	b.UserID(obj.UserID)
	b.EventID(obj.EventID)
	b.EventData(eventData)
	_, err := vk.MessagesSendMessageEventAnswer(b.Params)
	if err != nil {
		panic(err)
	}
}
