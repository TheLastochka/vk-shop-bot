package main

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/joho/godotenv"
	"vk_bot/pkg/handlers"
	"vk_bot/pkg/vkAPI"
)

func main() {
	envs, _ := godotenv.Read()
	vk := vkAPI.InitNewVK(envs["TOKEN"])
	lp, _ := longpoll.NewLongPollCommunity(vk)

	lp.MessageNew(handlers.MessageNewHandler)

	lp.MessageEvent(handlers.MessageEventHandler)

	fmt.Println("===Bot started===")
	if err := lp.Run(); err != nil {
		fmt.Println(err)
	}

}
