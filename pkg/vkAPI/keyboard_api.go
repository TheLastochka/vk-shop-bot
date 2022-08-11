package vkAPI

import (
	"encoding/json"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"strconv"
	"strings"
	"vk_bot/pkg/keyboard"
)

func SendKeyboardNew(obj events.MessageNewObject, command string, msg string) {
	b := params.NewMessagesSendBuilder()
	b.PeerID(obj.Message.PeerID)
	b.Message(msg)
	b.RandomID(0)
	b.Keyboard(keyboard.NewMenu(command))
	_, err := vk.MessagesSend(b.Params)
	if err != nil {
		panic(err)
	}
}

func SendKeyboardUpdate(obj events.MessageEventObject, command string, msg string) {
	b := params.NewMessagesEditBuilder()
	b.PeerID(obj.PeerID)
	b.ConversationMessageID(obj.ConversationMessageID)
	b.Message(msg)
	b.Keyboard(keyboard.NewMenu(command))
	_, err := vk.MessagesEdit(b.Params)
	if err != nil {
		panic(err)
	}
}

func GetPayload(rawPayload json.RawMessage) keyboard.Payload {
	var payload keyboard.Payload
	unqStr, _ := strconv.Unquote(string(rawPayload))
	_ = json.Unmarshal([]byte(unqStr), &payload)
	return payload
}

func GetMenuLabel(commandPath string) string {
	commandSplit := strings.Split(commandPath, "_")
	label := ""
	for _, cmd := range commandSplit {
		if label != "" {
			label += " -> "
		}
		label += keyboard.ButtonsMap[cmd].Label
	}
	if label == "" {
		label = "магазин"
	}
	return label
}
