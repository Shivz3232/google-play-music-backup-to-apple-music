package playlist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"shivu/google-play-music-backup-reader/models"
)

func Create(developerToken string, musicUserToken string, name string, description string) (*models.Playlist, error) {
	reqBody := &CreatePlaylistRequest{
		Attributes: CreatePlaylistRequestAttributes{
			Name:        name,
			Description: description,
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.music.apple.com/v1/me/library/playlists", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+developerToken)
	req.Header.Set("Music-User-Token", musicUserToken)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 201 {
		resBody, e := io.ReadAll(res.Body)
		if e != nil {
			return nil, err
		}
		fmt.Printf("Response: %s\n", resBody)
		return nil, errors.New("Unable to create playlist, status code: " + res.Status)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	container := &CreatePlaylistsResponse{}
	if err := json.Unmarshal(resBody, container); err != nil {
		return nil, err
	}

	return &container.Data[0], nil
}

type CreatePlaylistRequest struct {
	Attributes CreatePlaylistRequestAttributes `json:"attributes"`
}

type CreatePlaylistRequestAttributes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreatePlaylistsResponse struct {
	Data []models.Playlist `json:"data"`
}
