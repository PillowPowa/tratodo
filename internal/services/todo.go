package services

import (
	"errors"
	"fmt"
	"net/http"
	"tratodo/internal/domain/models"
	"tratodo/internal/repository"
	"tratodo/pkg/api"
)

type TodoRepository interface {
	GetById(id int) (*models.Todo, error)
	Create(todo *models.Todo) (*models.Todo, error)
	Delete(id int) error
}

type TodoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) GetById(id int) (*models.Todo, error) {
	todo, err := s.repo.GetById(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, api.NewApiError(http.StatusNotFound, fmt.Sprintf("TODO with id %v not found", id))
		}

		return nil, err
	}

	return todo, nil
}

func (s *TodoService) Create(todo *models.Todo) (*models.Todo, error) {
	t, err := s.repo.Create(todo)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *TodoService) Delete(id int) error {
	if _, err := s.repo.GetById(id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return api.NewApiError(http.StatusNotFound, fmt.Sprintf("TODO with id %v not found", id))
		}
		return err
	}

	err := s.repo.Delete(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return api.NewApiError(http.StatusNotFound, fmt.Sprintf("TODO with id %v not found", id))
		}

		return err
	}

	return nil
}
