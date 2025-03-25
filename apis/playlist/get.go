package playlist

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"shivu/google-play-music-backup-reader/models"
)

func Get(developerToken string, musicUserToken string) ([]models.Playlist, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.music.apple.com/v1/me/library/playlists", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+developerToken)
	req.Header.Set("Music-User-Token", musicUserToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, errors.New("Unable to fetch playlists, status code: " + res.Status)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	container := &GetPlaylistsResponse{}
	if err := json.Unmarshal(resBody, container); err != nil {
		return nil, err
	}

	return container.Data, nil
}

type GetPlaylistsResponse struct {
	Next string            `json:"next"`
	Data []models.Playlist `json:"data"`
}
