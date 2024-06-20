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

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) error {
	input := make(map[string]string)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return api.NewApiError(http.StatusBadRequest, "Invalid JSON")
	}

	token, err := c.service.Login(input["login"], input["password"])
	if err != nil {
		return err
	}

	c.writeJWTCookie(w, token)
	api.WriteJSON(w, http.StatusOK, true)
	return nil
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) error {
	c.writeJWTCookie(w, "")
	api.WriteJSON(w, http.StatusOK, true)
	return nil
}

func (c *AuthController) GetMe(w http.ResponseWriter, r *http.Request) error {
	return api.NewApiError(http.StatusNotImplemented, "Not Implemented")
}

func (c *AuthController) writeJWTCookie(w http.ResponseWriter, token string) {
	cookieKey := fmt.Sprintf("jwt=%s; Path=/; HttpOnly", token)
	w.Header().Add("Set-Cookie", cookieKey)
}
