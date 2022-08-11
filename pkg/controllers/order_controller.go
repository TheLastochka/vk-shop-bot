package controllers

import (
	"encoding/json"
	"github.com/SevereCloud/vksdk/v2/events"
	"os"
	"strconv"
	"strings"
	"vk_bot/pkg/keyboard"
)

func UpdateOrder(peerID int, updatedOrder UserOrder) {
	users := make(map[string]UserData)
	fileName := jsonFileName
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &users); err != nil {
		panic(err)
	}

	strPeerID := strconv.Itoa(peerID)
	u := users[strPeerID]
	u.Order = updatedOrder
	users[strPeerID] = u

	ordersJSON, err := json.Marshal(users)
	if err := os.WriteFile(fileName, ordersJSON, 0666); err != nil {
		panic(err)
	}
}

func SaveOrder(commandPath string, obj events.MessageEventObject) {
	commandSplit := strings.Split(commandPath, "_")
	amount, _ := strconv.Atoi(commandSplit[3])
	newOrder := UserOrder{Action: commandSplit[1], Animal: commandSplit[2], Amount: amount}

	UpdateOrder(obj.PeerID, newOrder)
}

func getOrder(peerID int) UserOrder {
	orders := make(map[string]UserData)
	fileName := jsonFileName
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(file, &orders)
	return orders[strconv.Itoa(peerID)].Order
}

func GetOrderStr(peerID int) string {
	userOrder := getOrder(peerID)
	var msg string
	if userOrder.Amount == 0 {
		msg = "Ваш заказ: нет заказа"
	} else {
		action := keyboard.ButtonsMap[userOrder.Action].Label
		animal := keyboard.ButtonsMap[userOrder.Animal].Label
		amount := strconv.Itoa(userOrder.Amount)
		msg = "Ваш заказ: " + action + " " + animal + " " + amount
	}
	return msg
}
