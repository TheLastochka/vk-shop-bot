package controllers

var jsonFileName = "./data/users.json"

type UserData struct {
	Order UserOrder `json:"order"`
}

type UserOrder struct {
	Action string `json:"action"`
	Animal string `json:"animal"`
	Amount int    `json:"amount"`
}
