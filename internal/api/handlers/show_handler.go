package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sullyh7/aetoons/internal/db/repositories"
)

type ShowHandler struct {
	repository *repositories.ShowRepository
}

func NewShowHandler(r *repositories.ShowRepository) *ShowHandler {
	return &ShowHandler{
		repository: r,
	}
}

func (h *ShowHandler) CreateShow(c echo.Context) error {
	c.JSON(200, map[string]string{"working": "yes"})
	return nil
}
