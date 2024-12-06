package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sullyh7/aetoons/internal/api/handlers"
	"github.com/sullyh7/aetoons/internal/config"
)

type API struct {
	*echo.Echo
	Config *config.Config
}

func NewAPI(cfg *config.Config, h *handlers.ShowHandler) *API {
	e := echo.New()
	e.Debug = true

	e.GET("/", h.CreateShow)

	return &API{
		Echo:   e,
		Config: cfg,
	}
}
