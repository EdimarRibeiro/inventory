package models

type Account struct {
	Name         string         `json:"fantasia"`
	Document     string         `json:"document"`
	Registration string         `json:"registration"`
	Address      AccountAddress `json:"AccountAddress"`
	NameUser     string         `json:"name"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	Suframa      string         `json:"suframa"`
}
