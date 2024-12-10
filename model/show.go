package model

import "gorm.io/gorm"

type Show struct {
	gorm.Model
	MALID       int          `json:"id" gorm:"unique"`
	Title       string       `json:"title"`
	MainPicture PictureTypes `gorm:"embedded;embeddedPrefix:main_picture_" json:"main_picture"`

	Episodes []Episode
}

type PictureTypes struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

// {
// 	"id": 21,
// 	"title": "One Piece",
// 	"main_picture": {
// 	  "medium": "https://cdn.myanimelist.net/images/anime/1244/138851.jpg",
// 	  "large": "https://cdn.myanimelist.net/images/anime/1244/138851l.jpg"
// 	}
//   }
