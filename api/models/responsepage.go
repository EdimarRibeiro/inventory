package models

type ResponsePage struct {
	Records   interface{} `json:"records,omitempty"`
	Pages     int64       `json:"pages"`
	TotalRows int64       `json:"totalRows"`
	Page      int64       `json:"page"`
	Rows      int64       `json:"rows"`
}
