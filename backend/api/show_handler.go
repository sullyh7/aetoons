package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sullyh7/aetoons/model"
)

const MAL_AUTH_PARAM_NAME = "X-MAL-CLIENT-ID"

type AddShowRequest struct {
	MALID int `json:"mal_id"`
}

func (s *Server) handleAddShow(c echo.Context) error {
	var request AddShowRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if request.MALID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "MAL ID is required"})
	}
	url := fmt.Sprintf("https://api.myanimelist.net/v2/anime/%d", request.MALID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create request"})
	}
	req.Header.Set(MAL_AUTH_PARAM_NAME, s.config.MALClientID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data from MAL API"})
	}
	defer resp.Body.Close()
	var show model.Show
	json.NewDecoder(resp.Body).Decode(&show)
	s.store.SaveShow(&show)
	return c.JSON(http.StatusOK, show)
}

func (s *Server) handleShows(c echo.Context) error {
	shows, err := s.store.FetchShows()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, shows)
}
