package service

import (
	"context"
	"fmt"
	"os"

	"github.com/silentsokolov/go-vimeo/vimeo"
	"github.com/sullyh7/aetoons/config"
	"golang.org/x/oauth2"
)

type VimeoService struct {
	Client *vimeo.Client
}

func NewVimeoService(accessToken string) (*VimeoService, error) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(context.Background(), ts)

	client := vimeo.NewClient(tc, &vimeo.Config{
		Uploader: &config.Uploader{},
	})
	return &VimeoService{Client: client}, nil
}

func (vs *VimeoService) UploadVideo(filePath string) (*vimeo.Video, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	video, resp, err := vs.Client.Users.UploadVideo("", file)
	if err != nil {
		return nil, fmt.Errorf("failed to upload video: %w", err)
	}
	defer resp.Body.Close()

	return video, nil
}
