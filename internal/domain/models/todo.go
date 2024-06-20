package models

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title" validate:"required,min=3,max=32"`
	Completed bool   `json:"completed"`
}
