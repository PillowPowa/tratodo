package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	"tratodo/internal/domain/models"
	"tratodo/internal/libs/jwt"
	"tratodo/internal/repository"
	"tratodo/pkg/api"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	GetUniqueUser(username string, email string) (*models.User, error)
	Create(user *models.User) (int64, error)
}

type AuthService struct {
	repo   AuthRepository
	jwtExp time.Duration
}

func NewAuthService(repo AuthRepository, jwtExp time.Duration) *AuthService {
	return &AuthService{
		repo:   repo,
		jwtExp: jwtExp,
	}
}

func (s *AuthService) Login(login, password string) (string, error) {
	u, err := s.repo.GetUniqueUser(login, login)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return "", api.NewApiError(http.StatusNotFound, fmt.Sprintf("User with login %v not found", login))
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PassHash), []byte(password))
	if err != nil {
		return "", api.NewApiError(http.StatusForbidden, "Invalid credentials")
	}

	return jwt.NewToken(u.ID, s.jwtExp)
}

func (s *AuthService) Register(user *models.POSTUser) (string, error) {
	_, err := s.repo.GetUniqueUser(user.Username, user.Email)
	if err == nil && !errors.Is(err, repository.ErrNotFound) {
		return "", api.NewApiError(http.StatusConflict, "User already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "nil", fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}

	u := &models.User{
		PassHash: hashedPassword,
		Username: user.Username,
		Email:    user.Email,
	}

	id, err := s.repo.Create(u)
	if err != nil {
		return "", err
	}

	return jwt.NewToken(id, s.jwtExp)
}
