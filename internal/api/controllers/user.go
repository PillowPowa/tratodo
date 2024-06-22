package controllers

import (
	"net/http"
	"tratodo/internal/domain/models"
	"tratodo/internal/libs/context"
	"tratodo/pkg/api"
)

type UserService interface {
	GetById(id int64) (*models.User, error)
}

type UserController struct {
	service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// @Summary Get self user
// @Description Get user using infered ID from JWT cookie
// @ID get-user-by-id
// @Produce json
// @Security jwt
// @Success 200 {object} models.User
// @Failure 404,401 {object} api.Error
// @Router /user/ [get]
func (c *UserController) GetMe(w http.ResponseWriter, r *http.Request) error {
	id, err := context.InferAuthContext(r.Context())
	if err != nil {
		return err
	}

	u, err := c.service.GetById(id)
	if err != nil {
		return err
	}

	return api.WriteJSON(w, http.StatusOK, u)
}
