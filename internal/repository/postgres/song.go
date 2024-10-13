package postgres

import (
	"context"
	"fmt"
	"github.com/Tinddd28/TestTask/internal/models"
	"github.com/Tinddd28/TestTask/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SongPostgres struct {
	db *pgxpool.Pool
}

func NewSongPostgres(db *pgxpool.Pool) *SongPostgres {
	return &SongPostgres{db: db}
}

func (r *SongPostgres) GetAllSongs(group, songF, startDate, endDate string, offset int) ([]models.Song, error) {
	query := fmt.Sprintf("SELECT id, group_name, song_name, release_date, link FROM %s WHERE 1=1", songTable)

	var args []interface{}
	idx := 1
	offset -= 1
	if group != "" {
		query += fmt.Sprintf(" AND group_name ILIKE $%d", idx)
		args = append(args, "%"+group+"%")
		idx++
	}

	if songF != "" {
		query += fmt.Sprintf(" AND song_name ILIKE $%d", idx)
		args = append(args, "%"+songF+"%")
		idx++
	}

	if startDate != "" {
		query += fmt.Sprintf(" AND release_date >= $%d", idx)
		args = append(args, startDate)
		idx++
	}

	if endDate != "" {
		query += fmt.Sprintf(" AND release_date <= $%d", idx)
		args = append(args, endDate)
		idx++
	}
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", idx, idx+1)
	args = append(args, limitSong, offset)

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	songs := make([]models.Song, 0)
	for rows.Next() {
		var song models.Song
		err := rows.Scan(&song.ID, &song.Group, &song.Name, &song.Year, &song.Link)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func (r *SongPostgres) GetSong(id, page, pageSize int) ([]models.Verse, error) {
	offset := (page - 1) * pageSize

	query := fmt.Sprintf("SELECT id, text FROM %s WHERE song_id = $1 ORDER BY id LIMIT $2 OFFSET $3", versesTable)
	rows, err := r.db.Query(context.Background(), query, id, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	verses := make([]models.Verse, 0)
	for rows.Next() {
		var verse models.Verse
		if err := rows.Scan(&verse.ID, &verse.Text); err != nil {
			return nil, err
		}

		verses = append(verses, verse)
	}

	return verses, nil
}

func (r *SongPostgres) DeleteSong(id int) error {
	query := fmt.Sprintf("delete from %s where id = $1", songTable)
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}

func (r *SongPostgres) UpdateSong(id int, song models.Song) error {
	query := fmt.Sprintf("UPDATE %s SET group_name=$2, song_name=$3, release_date=$4, link=$5 WHERE id=$1", songTable)
	cmdTag, err := r.db.Exec(context.Background(), query, id, song.Group, song.Name, song.Year, song.Link)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("song with id %d not found", id)
	}
	return nil
}

func (r *SongPostgres) CreateSong(song models.InsertSongDb) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (group_name, song_name, release_date, link) VALUES ($1, $2, $3, $4, $5)")

	parsedDate, err := utils.ParseDate(song.Year)
	if err != nil {
		return 0, err
	}

	dateString := parsedDate.Format("2006-01-02")
	var id int
	err = r.db.QueryRow(context.Background(), query, song.Group, song.Name, dateString, song.Link).Scan(id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SongPostgres) CreateVerse(idSong int, text string, num int) error {
	query := fmt.Sprintf("INSERT INTO %s (song_id, text, num) VALUES ($1, $2, $3)", versesTable)
	var id int
	err := r.db.QueryRow(context.Background(), query, idSong, text, num).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
