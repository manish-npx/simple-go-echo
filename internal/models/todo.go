package models

type Todo struct {
	ID    int64  `json:"id"`
	Title string `json:"title" validate:"required"`
	Done  bool   `json:"done"`
}
