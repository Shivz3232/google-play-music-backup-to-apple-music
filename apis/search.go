package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"shivu/google-play-music-backup-reader/models"
	"strings"
)

func SearchSong(developerToken string, musicUserToken string, storefront string, searchStr string, types []string, localization string, with []string) ([]models.Song, error) {
	term := getTermQueryParam(searchStr)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.music.apple.com/v1/catalog/%s/search", storefront), nil)
	if err != nil {
		return nil, err
	}

	queryParams := req.URL.Query()

	queryParams.Add("term", term)
	queryParams.Add("types", strings.Join(types, ","))
	queryParams.Add("l", localization)
	queryParams.Add("with", strings.Join(with, ","))

	req.URL.RawQuery = queryParams.Encode()

	req.Header.Set("Authorization", "Bearer "+developerToken)
	req.Header.Set("Music-User-Token", musicUserToken)

	fmt.Println(req.RequestURI)

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
