package handlers

import (
	"context"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/events"
	"strconv"
	"strings"
	"vk_bot/pkg/controllers"
	"vk_bot/pkg/vkAPI"
)

func MessageNewHandler(ctx context.Context, obj events.MessageNewObject) {
	fmt.Println("====MSG====")
	msg := obj.Message.Text
	fmt.Println(msg)
	if msg == "Начать" || msg == "начать" {
		vkAPI.SendKeyboardNew(obj, "start", "магазин")
	}
	fmt.Println("===========")
}

func MessageEventHandler(ctx context.Context, obj events.MessageEventObject) {
	fmt.Println("===EVENT===")
	payload := vkAPI.GetPayload(obj.Payload)
	commandSplit := strings.Split(payload.Command, "_")
	newCommand := commandSplit[len(commandSplit)-1]
	fmt.Println(newCommand)
	var eventData = ``
	if newCommand == "orderText" {
		vkAPI.SendOrderInline(obj, payload.Command)
	} else if newCommand == "removeOrder" {
		controllers.UpdateOrder(obj.PeerID, controllers.UserOrder{})
		vkAPI.SendOrderInline(obj, payload.Command)
		eventData = `{
		   "type": "show_snackbar",
		   "text": "Заказ отменен"
		 }`
	} else if num, _ := strconv.Atoi(newCommand); num > 0 {
		controllers.SaveOrder(payload.Command, obj)
		vkAPI.SendKeyboardUpdate(obj, "start", "магазин")
		eventData = `{
		   "type": "show_snackbar",
		   "text": "Заказ принят"
		 }`
	} else {
		var msg = vkAPI.GetMenuLabel(payload.Command)
		vkAPI.SendKeyboardUpdate(obj, payload.Command, msg)
	}

	vkAPI.SendMessageEventAnswer(obj, eventData)

	fmt.Println("===========")
}
