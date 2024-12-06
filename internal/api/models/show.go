package models

import "gorm.io/gorm"

type Show struct {
	gorm.Model
	Title        string
	InfoURL      string
	ThumbnailURL string
	EpisodeCount int

	Episodes []Episode
}

type Episode struct {
	gorm.Model
	Title    string
	Number   int
	VideoURL string

	ShowID uint
}
