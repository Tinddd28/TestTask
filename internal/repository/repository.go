package repository

import (
	"github.com/Tinddd28/TestTask/internal/models"
	"github.com/Tinddd28/TestTask/internal/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Song interface {
	GetAllSongs(group, songF, startDate, endDate string, offset int) ([]models.Song, error)
	GetSong(id, page, pageSize int) ([]models.Verse, error)
	DeleteSong(id int) error
	UpdateSong(id int, song models.Song) error
	CreateSong(song models.InsertSongDb) (int, error)
	CreateVerse(id int, text string, num int) error
}

type Repository struct {
	Song
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Song: postgres.NewSongPostgres(db),
	}
}
