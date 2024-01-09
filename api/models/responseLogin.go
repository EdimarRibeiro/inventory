package models

type ResponseUser struct {
	Login string `json:"login"`
	Image string `json:"image"`
}

type ResponseLogin struct {
	Token         string       `json:"token"`
	User          ResponseUser `json:"user"`
	Authenticated bool         `json:"authenticated"`
}
