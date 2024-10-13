package api

import (
	"encoding/json"
	"github.com/Tinddd28/TestTask/internal/models"
	"net/http"
	"net/url"
)

func GetInfo(baseUrl string, s models.RequestSong) (*models.SongDetail, error) {
	apiUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	query := apiUrl.Query()
	query.Set("group", s.Group)
	query.Set("song", s.Name)

	resChan := make(chan *models.SongDetail)
	errChan := make(chan error)
	go func() {
		apiUrl.RawQuery = query.Encode()

		resp, err := http.Get(apiUrl.String())
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			errChan <- err
			return
		}

		var songDetail models.SongDetail
		if err = json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
			errChan <- err
			return
		}

		resChan <- &songDetail
	}()

	select {
	case result := <-resChan:
		return result, nil
	case err := <-errChan:
		return nil, err
	}
}
