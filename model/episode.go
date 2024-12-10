package model

import "gorm.io/gorm"

type Episode struct {
	gorm.Model
	Title         string
	EpisodeNumber int
	VideoUrl      string

	ShowID uint
}
