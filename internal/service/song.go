package service

import (
	"github.com/Tinddd28/TestTask/internal/api"
	"github.com/Tinddd28/TestTask/internal/models"
	"github.com/Tinddd28/TestTask/internal/repository"
)

type SongService struct {
	repo repository.Song
}

func NewSongService(repo repository.Song) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) GetAllSongs(group, songF, startDate, endDate string, offset int) ([]models.Song, error) {
	return s.repo.GetAllSongs(group, songF, startDate, endDate, offset)
}

func (s *SongService) GetSong(id, page, pageSize int) ([]models.Verse, error) {
	return s.repo.GetSong(id, page, pageSize)
}

func (s *SongService) DeleteSong(id int) error {
	return s.repo.DeleteSong(id)
}

func (s *SongService) UpdateSong(id int, song models.Song) error {
	return s.repo.UpdateSong(id, song)
}

func (s *SongService) CreateSong(song models.RequestSong) (int, error) {
	songDetail, err := api.GetInfo("example.com", song)
	if err != nil {
		return 0, err
	}

	songToInsert := models.InsertSongDb{
		Group: song.Group,
		Name:  song.Name,
		Year:  songDetail.Year,
		Link:  songDetail.Link,
	}

	songId, err := s.repo.CreateSong(songToInsert)
	if err != nil {
		return 0, err
	}
	idx := 0
	for _, text := range songDetail.Text {
		idx += 1
		if err := s.repo.CreateVerse(songId, text, idx); err != nil {
			return 0, err
		}
	}

	return songId, nil
}
