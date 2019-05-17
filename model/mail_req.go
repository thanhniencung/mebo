package model

type Mail struct {
	To string `json:"to"`
	Subject string `json:"subject"`
	Name string `json:"name"`
}
