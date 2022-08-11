package vkAPI

import (
	"github.com/SevereCloud/vksdk/v2/events"
	"strings"
	"vk_bot/pkg/controllers"
)

func SendOrder(peerID int) {
	SendSimpleMsg(peerID, controllers.GetOrderStr(peerID))
}

func SendOrderInline(obj events.MessageEventObject, commandPath string) {
	commandSplit := strings.Split(commandPath, "_")
	prevCommandPath := strings.Join(commandSplit[:len(commandSplit)-1], "_")
	SendKeyboardUpdate(obj, prevCommandPath, controllers.GetOrderStr(obj.PeerID))
}
