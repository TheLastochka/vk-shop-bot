package keyboard

import (
	"github.com/SevereCloud/vksdk/v2/object"
	"strconv"
	"strings"
)

func NewMenu(commandPath string) *object.MessagesKeyboard {
	kb := object.NewMessagesKeyboard(false)
	kb.Inline = true
	commandList := strings.Split(commandPath, "_")
	lastCommand := commandList[len(commandList)-1]
	if lastCommand != "start" {
		prevCommandPath := strings.Join(commandList[:len(commandList)-1], "_")
		kb.AddRow()
		kb.AddCallbackButton("<-Назад", "{\"command\":\""+prevCommandPath+"\"}", "positive")
	}

	if len(animalShop[lastCommand]) != 0 && animalShop[lastCommand][0][0][:3] == "num" {
		part := strings.Split(animalShop[lastCommand][0][0], ",")
		num, _ := strconv.Atoi(strings.Split(part[0], "=")[1])
		rowLen, _ := strconv.Atoi(strings.Split(part[1], "=")[1])
		generateNumpad(kb, commandPath, num, rowLen)
		return kb
	}

	generateDefaultKeyboard(kb, commandPath, animalShop[lastCommand])

	return kb

}

func generateDefaultKeyboard(kb *object.MessagesKeyboard, curCommand string, btnMx [][]string) {
	for _, btnList := range btnMx {
		kb.AddRow()
		for _, btnID := range btnList {
			btn := ButtonsMap[btnID]
			kb.AddCallbackButton(btn.Label, "{\"command\":\""+curCommand+"_"+btn.command+"\"}", btn.color)
		}

	}
}

func generateNumpad(kb *object.MessagesKeyboard, commandPath string, n int, rowLen int) {
	iMax := n / rowLen
	if n%rowLen != 0 {
		iMax += 1
	}
	for i := 0; i < iMax; i++ {
		jMax := rowLen
		if i*rowLen+jMax-1 >= n {
			jMax = n % rowLen
		}
		kb.AddRow()
		for j := 0; j < jMax; j++ {
			strNum := strconv.Itoa(i*rowLen + j + 1)
			kb.AddCallbackButton(strNum, "{\"command\":\""+commandPath+"_"+strNum+"\"}", "secondary")
		}
	}
}
