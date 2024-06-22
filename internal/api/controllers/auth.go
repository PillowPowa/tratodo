package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tratodo/internal/domain/models"
	"tratodo/pkg/api"

	"github.com/go-playground/validator/v10"
)

type AuthService interface {
	Register(user *models.POSTUser) (string, error)
	Login(login, password string) (string, error)
}

type AuthController struct {
	service AuthService
}

func NewAuthController(service AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

// @Summary Sign up
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.POSTUser true "User data"
// @Success 201 {boolean} boolean
// @Failure 400,500,409 {object} api.Error
// @Router /auth/register [post]
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) error {
	user := new(models.POSTUser)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return api.NewApiError(http.StatusBadRequest, "Invalid JSON")
	}

	if err := validator.New().Struct(user); err != nil {
		return api.NewApiError(http.StatusBadRequest, err.Error())
	}

	token, err := c.service.Register(user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	c.writeJWTCookie(w, token)
	api.WriteJSON(w, http.StatusCreated, true)
	return nil
}

type signInInput struct {
	Login    string `json:"login" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6"`
}

// @Summary Sign in
// @Description Authenticate user
// @Tags auth
// @Accept json
// @Produce json
// @Param login body signInInput true "User login"
// @Success 200 {boolean} boolean
// @Failure 400,404,500 {object} api.Error
// @Router /auth/login [post]
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) error {
	input := new(signInInput)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return api.NewApiError(http.StatusBadRequest, "Invalid JSON")
	}

	if err := validator.New().Struct(input); err != nil {
		return api.NewApiError(http.StatusBadRequest, err.Error())
	}

	token, err := c.service.Login(input.Login, input.Password)
	if err != nil {
		return err
	}

	c.writeJWTCookie(w, token)
	return api.WriteJSON(w, http.StatusOK, true)
}

// @Summary Log out
// @Description Invalidate JWT token
// @Tags auth
// @Security jwt
// @Success 200 {boolean} boolean
// @Failure 401,500 {object} api.Error
// @Router /auth/logout [delete]
func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) error {
	c.writeJWTCookie(w, "")
	return api.WriteJSON(w, http.StatusOK, true)
}

func (c *AuthController) writeJWTCookie(w http.ResponseWriter, token string) {
	cookieKey := fmt.Sprintf("jwt=%s; Path=/; HttpOnly", token)
	w.Header().Add("Set-Cookie", cookieKey)
}
