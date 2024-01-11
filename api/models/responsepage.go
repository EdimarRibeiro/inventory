package models

type ResponsePage struct {
	Records interface{} `json:"records,omitempty"`
	Pages   int64       `json:"pages"`
	Page    int64       `json:"page"`
	Rows    int64       `json:"rows"`
}
