package playlist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func InsertSong(developerToken string, musicUserToken string, playlistId string, songId string) error {
	reqBody := &LibraryPlaylistTracksRequest{
		Data: []LibraryPlaylistTracksRequestData{
			{
				Id:   songId,
				Type: "songs",
			},
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.music.apple.com/v1/me/library/playlists/%s/tracks", url.PathEscape(playlistId)), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+developerToken)
	req.Header.Set("Music-User-Token", musicUserToken)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	} else if res.StatusCode != 201 && res.StatusCode != 204 {
		resBody, e := io.ReadAll(res.Body)
		if e != nil {
			return err
		}
		fmt.Printf("Response: %s\n", resBody)
		return errors.New("Unable to insert song, status code: " + res.Status)
	}

	return nil
}

type LibraryPlaylistTracksRequest struct {
	Data []LibraryPlaylistTracksRequestData `json:"data"`
}

type LibraryPlaylistTracksRequestData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}
