package config

import (
	"os"

	"github.com/eventials/go-tus"
	"github.com/joho/godotenv"
	"github.com/silentsokolov/go-vimeo/vimeo"
)

type Config struct {
	MALClientID      string
	VimeoAccessToken string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	return &Config{
		MALClientID:      os.Getenv("MAL_CLIENT_ID"),
		VimeoAccessToken: os.Getenv("VIMEO_ACCESS_TOKEN"),
	}, nil
}

type Uploader struct{}

func (u Uploader) UploadFromFile(c *vimeo.Client, uploadURL string, f *os.File) error {
	tusClient, err := tus.NewClient(uploadURL, nil)
	if err != nil {
		return err
	}

	upload, err := tus.NewUploadFromFile(f)
	if err != nil {
		return err
	}

	uploader := tus.NewUploader(tusClient, uploadURL, upload, 0)

	return uploader.Upload()
}
