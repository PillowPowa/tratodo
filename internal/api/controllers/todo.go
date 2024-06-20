package controllers

import (
	"net/http"
	"strconv"
	"tratodo/internal/domain/models"
	"tratodo/pkg/api"

	"github.com/gorilla/mux"
)

type TodoService interface {
	GetById(id int) (*models.Todo, error)
	Create(todo *models.Todo) (*models.Todo, error)
	Delete(id int) error
}

type TodoController struct {
	service TodoService
}

func NewTodoController(service TodoService) *TodoController {
	return &TodoController{
		service: service,
	}
}

func (c *TodoController) GetById(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return api.NewApiError(http.StatusBadRequest, "TODO ID should be an integer")
	}
	todo, err := c.service.GetById(id)

	if err != nil {
		return err
	}

	api.WriteJSON(w, http.StatusOK, todo)
	return nil
}
