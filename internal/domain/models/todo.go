package models

type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserId    int64  `json:"-"`
}

type POSTTodo struct {
	Title string `json:"title" validate:"required,min=3,max=255"`
}

type PatchTodo struct {
	Title     *string `json:"title" validate:"omitempty,min=3,max=255"`
	Completed *bool   `json:"completed" validate:"omitempty"`
}

type TodoQuery struct {
	SortBy string `json:"sort_by" validate:"omitempty,oneof=newest oldest"`
	Filter string `json:"filter" validate:"omitempty,oneof=completed uncompleted"`
}
