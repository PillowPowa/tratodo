package services

import (
	"errors"
	"net/http"
	"tratodo/internal/domain/models"
	"tratodo/internal/repository"
	"tratodo/pkg/api"
)

type UserRepository interface {
	GetById(id int64) (*models.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetById(id int64) (*models.User, error) {
	u, err := s.repo.GetById(id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, api.NewApiError(http.StatusNotFound, "User not found")
	}
	return u, err
}
