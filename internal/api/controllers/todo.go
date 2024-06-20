package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tratodo/internal/domain/models"
	"tratodo/pkg/api"

	"github.com/go-playground/validator/v10"
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

func (c *TodoController) Create(w http.ResponseWriter, r *http.Request) error {
	todo := new(models.Todo)
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		return api.NewApiError(http.StatusBadRequest, "Invalid JSON")
	}

	if err := validator.New().Struct(todo); err != nil {
		return api.NewApiError(http.StatusBadRequest, err.Error())
	}

	t, err := c.service.Create(todo)
	if err != nil {
		return err
	}

	api.WriteJSON(w, http.StatusCreated, t)
	return nil
}

func (c *TodoController) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return api.NewApiError(http.StatusBadRequest, "TODO ID should be an integer")
	}

	err = c.service.Delete(id)
	if err != nil {
		return err
	}

	api.WriteJSON(w, http.StatusOK, true)
	return nil
}
