package model

type AppDescription struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UpdatedAt string `json:"updated_at"`
}
