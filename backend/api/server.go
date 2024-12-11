package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sullyh7/aetoons/api/service"
	"github.com/sullyh7/aetoons/config"
	"github.com/sullyh7/aetoons/model"
)

type Server struct {
	*echo.Echo
	store        model.Store
	config       *config.Config
	vimeoService *service.VimeoService
}

func NewServer(store model.Store, vimeoService *service.VimeoService, config *config.Config) *Server {
	mux := echo.New()
	return &Server{
		mux,
		store,
		config,
		vimeoService,
	}
}

func (s *Server) SetupRoutes() {
	s.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"test": "success"})
	})

	s.GET("/shows", s.handleShows)

	s.POST("/add-show", s.handleAddShow)
	s.POST("/add-episode", s.handleAddEpisode)
	s.POST("/add-episode-from-url", s.handleAddEpisodeFromURL)
}
