package postgres

import (
	"context"
	"fmt"
	"github.com/Tinddd28/TestTask/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SongPostgres struct {
	db *pgxpool.Pool
}

func NewSongPostgres(db *pgxpool.Pool) *SongPostgres {
	return &SongPostgres{db: db}
}

func (r *SongPostgres) GetAllSongs(offset int) ([]models.Song, error) {

	query := fmt.Sprintf("SELECT id, author_name, song_name, release_date, link FROM %s", songTable)
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	songs := make([]models.Song, 0)
	for rows.Next() {
		var song models.Song
		err := rows.Scan(&song.ID, &song.Name, &song.Song, &song.Year, &song.Link)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (r *SongPostgres) GetSong(id int) (models.ResponseSong, error) {

	return models.ResponseSong{}, nil
}

func (r *SongPostgres) DeleteSong(id int) error {
	return nil
}

func (r *SongPostgres) UpdateSong(id int, song models.Song) error {
	return nil
}

func (r *SongPostgres) CreateSong(song models.RequestSong) (int, error) {
	return 0, nil
}
