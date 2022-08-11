package keyboard

//start
//[купить]
//	->[гусей]
//		->[1, 2, 3]
//	->[куриц]
//		->[1, 2, 3]
//[продать]
//	->[гусей]
//		->[1, 2, 3]
//	->[куриц]
//		->[1, 2, 3]
//start_buy_chicken_3

type button struct {
	Label   string
	command string
	color   string
}

type Payload struct {
	Command string `json:"command"`
}

var ButtonsMap = map[string]button{
	"buy":         {"купить", "buy", "positive"},
	"sell":        {"продать", "sell", "negative"},
	"goose":       {"гусей", "goose", "secondary"},
	"chicken":     {"куриц", "chicken", "secondary"},
	"dog":         {"собак", "dog", "secondary"},
	"order":       {"заказ", "order", "primary"},
	"orderText":   {"описание", "orderText", "primary"},
	"removeOrder": {"отменить", "removeOrder", "negative"},
}

var animalList = []string{"goose", "chicken", "dog"}

var animalShop = map[string][][]string{
	"start":   {{"buy", "sell"}, {"order"}},
	"buy":     {animalList},
	"sell":    {animalList},
	"order":   {{"orderText", "removeOrder"}},
	"goose":   {{"num=3,row=3"}}, // numpad
	"chicken": {{"num=7,row=3"}},
	"dog":     {{"num=5,row=2"}},
}
