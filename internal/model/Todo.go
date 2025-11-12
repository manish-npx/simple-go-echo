package model

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed string `json:"completed"`
}
