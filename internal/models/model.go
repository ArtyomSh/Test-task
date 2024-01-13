package models

type Rate struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Response struct {
	Message interface{}
	Error   string
}
