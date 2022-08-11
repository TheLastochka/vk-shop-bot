package vkAPI

import "github.com/SevereCloud/vksdk/v2/api"

var vk *api.VK

func InitNewVK(tokens ...string) *api.VK {
	vk = api.NewVK(tokens...)
	return vk
}
