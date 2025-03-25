package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"shivu/google-play-music-backup-reader/models"
	"strings"
)

func SearchSong(developerToken string, musicUserToken string, storefront string, searchStr string, types string) ([]models.Song, error) {
	term := getTermQueryParam(searchStr)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.music.apple.com/v1/catalog/%s/search?term=%s&types=%s", storefront, url.QueryEscape(term), url.QueryEscape(types)), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+developerToken)
	req.Header.Set("Music-User-Token", musicUserToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
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

	container := &SearchResponse{}
	if err := json.Unmarshal(resBody, container); err != nil {
		return nil, err
	}

	return container.Results.Songs.Data, nil
}

func getTermQueryParam(searchStr string) string {
	terms := strings.Split(searchStr, " ")

	return strings.Join(terms, "+")
}

type SearchResponse struct {
	Results SearchResponseResults `json:"results"`
}

type SearchResponseResults struct {
	Songs SongResult `json:"songs"`
}

type SongResult struct {
	Data []models.Song `json:"data"`
	Href string        `json:"href"`
	Next string        `json:"next"`
}
