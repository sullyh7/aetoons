package model

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store interface {
	SaveShow(*Show) error
	SaveEpisode(*Episode) error
	FetchShows() ([]Show, error)
}

type MySQLStore struct {
	db *gorm.DB
}

func NewMySQLStore(url string) MySQLStore {
	db, err := gorm.Open(sqlite.Open("app.sqlite3"), &gorm.Config{})

	db.AutoMigrate(&Show{}, &PictureTypes{}, &Episode{})
	if err != nil {
		log.Fatal(err)
	}
	return MySQLStore{db: db}
}

func (s MySQLStore) SaveShow(show *Show) error {
	if err := s.db.Create(show).Error; err != nil {
		return err
	}
	return nil
}

func (s MySQLStore) SaveEpisode(episode *Episode) error {
	// Ensure the Episode has a valid ShowID
	if episode.ShowID == 0 {
		return fmt.Errorf("invalid episode: missing ShowID")
	}

	// Save the Episode to the database
	if err := s.db.Create(episode).Error; err != nil {
		return fmt.Errorf("failed to save episode: %w", err)
	}

	return nil
}

func (s MySQLStore) FetchShows() ([]Show, error) {
	var shows []Show

	if err := s.db.Preload("Episodes").Find(&shows).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch shows: %w", err)
	}

	return shows, nil
}
