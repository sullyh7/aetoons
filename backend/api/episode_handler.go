package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sullyh7/aetoons/api/service"
	"github.com/sullyh7/aetoons/model"
)

const FILE_NAME = "content.mp4"
const DIR = "./uploads/"
const OUTPUT = "./uploads/content_with_subtitles.mp4"

type AddEpisodeRequest struct {
	Title         string `json:"title"`
	EpisodeNumber int    `json:"episode_number"`
	ShowID        uint   `json:"show_id"`
	VideoUrl      string `json:"video_url"`
}

func (s *Server) handleAddEpisode(c echo.Context) error {
	title := c.FormValue("title")
	episodeNumberStr := c.FormValue("episode_number")
	showIDStr := c.FormValue("show_id")

	if title == "" || episodeNumberStr == "" || showIDStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing required fields"})
	}

	episodeNumber, err := strconv.Atoi(episodeNumberStr)
	if err != nil || episodeNumber <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid episode number"})
	}

	showID, err := strconv.ParseUint(showIDStr, 10, 32)
	if err != nil || showID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid show ID"})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "File is required"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to open file"})
	}
	defer src.Close()

	if _, err := os.Stat("./uploads/"); os.IsNotExist(err) {
		if err := os.MkdirAll("./uploads/", os.ModePerm); err != nil {
			fmt.Println("Failed to create uploads directory:", err)
			return err
		}
	}
	filePath := fmt.Sprintf("./uploads/%s", FILE_NAME)
	dst, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save file"})
	}
	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("failed to save file contents: %w", err)
	}
	defer dst.Close()

	if err := service.Transcript(filePath, "./uploads"); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to transcribe file"})
	}

	video, err := s.vimeoService.UploadVideo(OUTPUT)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to upload file to vimeo"})
	}

	episode := model.Episode{
		Title:         title,
		EpisodeNumber: episodeNumber,
		VideoUrl:      video.Link,
		ShowID:        uint(showID),
	}

	if err := s.store.SaveEpisode(&episode); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save episode"})
	}

	return c.JSON(http.StatusOK, episode)
}

func (s *Server) handleAddEpisodeFromURL(c echo.Context) error {
	var request AddEpisodeRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload"})
	}

	if request.Title == "" || request.EpisodeNumber <= 0 || request.ShowID == 0 || request.VideoUrl == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing or invalid required fields"})
	}

	videoPath := fmt.Sprintf("./uploads/%d.mp4", request.EpisodeNumber) // Unique filename based on episode number
	resp, err := http.Get(request.VideoUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to download video"})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to download video: invalid URL or response"})
	}

	out, err := os.Create(videoPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create file for video"})
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save video file"})
	}

	if err := service.Transcript(videoPath, "./uploads"); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to transcribe video"})
	}

	video, err := s.vimeoService.UploadVideo(videoPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to upload video to Vimeo"})
	}

	episode := model.Episode{
		Title:         request.Title,
		EpisodeNumber: request.EpisodeNumber,
		VideoUrl:      video.Link,
		ShowID:        request.ShowID,
	}

	if err := s.store.SaveEpisode(&episode); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save episode"})
	}

	return c.JSON(http.StatusOK, episode)
}
