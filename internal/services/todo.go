package services

import (
	"errors"
	"fmt"
	"net/http"
	"tratodo/internal/domain/models"
	"tratodo/internal/repository"
	"tratodo/pkg/api"
)

var errOwnership = api.NewApiError(http.StatusForbidden, "You cannot view other users' todos")

type TodoRepository interface {
	GetAll(query *models.TodoQuery, userId int64) ([]*models.Todo, error)
	GetById(id int64) (*models.Todo, error)
	Create(todo *models.POSTTodo, userId int64) (*models.Todo, error)
	UpdateOne(id int64, todo *models.PatchTodo) (*models.Todo, error)
	Delete(id int64) error
}

type TodoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) GetAll(query *models.TodoQuery, userId int64) ([]*models.Todo, error) {
	return s.repo.GetAll(query, userId)
}

func (s *TodoService) GetById(id int64, userId int64) (*models.Todo, error) {
	todo, err := s.repo.GetById(id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, api.NewApiError(http.StatusNotFound, fmt.Sprintf("TODO with id %v not found", id))
	}

	if todo.UserId != userId {
		return nil, errOwnership
	}

	return todo, err
}

func (s *TodoService) Create(todo *models.POSTTodo, userId int64) (*models.Todo, error) {
	t, err := s.repo.Create(todo, userId)
	if err != nil && errors.Is(err, repository.ErrRef) {
		return nil, api.NewApiError(http.StatusBadRequest, "Provided todo author doesn't exists")
	}
	return t, err
}

func (s *TodoService) Delete(id int64, userId int64) error {
	todo, err := s.repo.GetById(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return api.NewApiError(http.StatusNotFound, fmt.Sprintf("TODO with id %v not found", id))
		}
		return err
	}

	if todo.UserId != userId {
		return errOwnership
	}

	return s.repo.Delete(id)
}

func (s *TodoService) UpdateOne(id int64, todo *models.PatchTodo, userId int64) (*models.Todo, error) {
	_, err := s.GetById(id, userId)
	if err != nil {
		return nil, err
	}

	return s.repo.UpdateOne(id, todo)
}
