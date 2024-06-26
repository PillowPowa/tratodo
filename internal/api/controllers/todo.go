package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"tratodo/internal/domain/models"
	"tratodo/internal/libs/context"
	"tratodo/pkg/api"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type TodoService interface {
	GetAll(query *models.TodoQuery, userId int64) ([]*models.Todo, error)
	GetById(id int64, userId int64) (*models.Todo, error)
	Create(todo *models.POSTTodo, userId int64) (*models.Todo, error)
	UpdateOne(id int64, todo *models.PatchTodo, userID int64) (*models.Todo, error)
	Delete(id int64, userId int64) error
}

type TodoController struct {
	service TodoService
}

func NewTodoController(service TodoService) *TodoController {
	return &TodoController{
		service: service,
	}
}

// @Summary Get all todos
// @Description Get all todos
// @Tags todo
// @Security jwt
// @Accept json
// @Produce json
// @Param sort_by query string false "Sort by" Enums(newest,oldest)
// @Param filter query string false "Filter" Enums(completed,uncompleted)
// @Success 200 {array} models.Todo
// @Failure 400,401,403,404,500 {object} api.Error
// @Router /todo/ [get]
func (c *TodoController) GetAll(w http.ResponseWriter, r *http.Request) error {
	query := new(models.TodoQuery)
	query.Filter = r.URL.Query().Get("filter")
	query.SortBy = r.URL.Query().Get("sort_by")

	if err := validator.New().Struct(query); err != nil {
		return api.NewApiError(http.StatusBadRequest, err.Error())
	}

	userId, authErr := context.InferAuthContext(r.Context())
	todos, err := c.service.GetAll(query, userId)

	if errors.Join(authErr, err) != nil {
		return err
	}

	return api.WriteJSON(w, http.StatusOK, todos)
}

// @Summary Get todo by ID
// @Description Get todo by ID
// @Tags todo
// @Security jwt
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Failure 400,401,403,404,500 {object} api.Error
// @Router /todo/{id} [get]
func (c *TodoController) GetById(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return api.NewApiError(http.StatusBadRequest, "TODO ID should be an integer")
	}

	userId, err := context.InferAuthContext(r.Context())
	if err != nil {
		return err
	}

	todo, err := c.service.GetById(id, userId)
	if err != nil {
		return err
	}

	api.WriteJSON(w, http.StatusOK, todo)
	return nil
}

// @Summary Create todo
// @Description Create todo
// @Tags todo
// @Security jwt
// @Accept json
// @Produce json
// @Param todo body models.POSTTodo true "Todo title"
// @Success 201 {object} models.Todo
// @Failure 400,401,500 {object} api.Error
// @Router /todo/ [post]
func (c *TodoController) Create(w http.ResponseWriter, r *http.Request) error {
	userId, err := context.InferAuthContext(r.Context())
	if err != nil {
		return err
	}

	todo := new(models.POSTTodo)
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		return api.NewApiError(http.StatusBadRequest, "Invalid JSON")
	}

	if err := validator.New().Struct(todo); err != nil {
		return api.NewApiError(http.StatusBadRequest, err.Error())
	}

	t, err := c.service.Create(todo, userId)
	if err != nil {
		return err
	}

	api.WriteJSON(w, http.StatusCreated, t)
	return nil
}

// @Summary Delete todo
// @Description Delete todo
// @Tags todo
// @Security jwt
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {boolean} boolean
// @Failure 400,401,403,404,500 {object} api.Error
// @Router /todo/{id} [delete]
func (c *TodoController) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return api.NewApiError(http.StatusBadRequest, "TODO ID should be an integer")
	}

	userId, err := context.InferAuthContext(r.Context())
	if err != nil {
		return err
	}

	err = c.service.Delete(id, userId)
	if err != nil {
		return err
	}

	return api.WriteJSON(w, http.StatusOK, true)
}

// @Summary Update todo
// @Description Update todo
// @Tags todo
// @Security jwt
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body models.PatchTodo true "Todo title"
// @Success 200 {object} models.Todo
// @Failure 400,401,403,404,500 {object} api.Error
// @Router /todo/{id} [patch]
func (c *TodoController) UpdateOne(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return api.NewApiError(http.StatusBadRequest, "TODO ID should be an integer")
	}

	userId, err := context.InferAuthContext(r.Context())
	if err != nil {
		return err
	}

	todo := new(models.PatchTodo)
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		return api.NewApiError(http.StatusBadRequest, "Invalid JSON")
	}

	if err := validator.New().Struct(todo); err != nil {
		return api.NewApiError(http.StatusBadRequest, err.Error())
	}

	t, err := c.service.UpdateOne(id, todo, userId)
	if err != nil {
		return err
	}

	api.WriteJSON(w, http.StatusOK, t)
	return nil
}
