package service

import (
	"github.com/Tinddd28/TestTask/internal/models"
	"github.com/Tinddd28/TestTask/internal/repository"
)

type Song interface {
	GetAllSongs(group, songF, startDate, endDate string, offset int) ([]models.Song, error)
	GetSong(id, page, pageSize int) ([]models.Verse, error)
	DeleteSong(id int) error
	UpdateSong(id int, song models.Song) error
	CreateSong(song models.RequestSong) (int, error)
}

type Service struct {
	Song
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Song: NewSongService(repo.Song),
	}
}
