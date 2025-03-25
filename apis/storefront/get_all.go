package storefront

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"shivu/google-play-music-backup-reader/models"
)

func GetAll(developerToken string, musicUserToken string) ([]models.StoreFront, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.music.apple.com/v1/storefronts", nil)
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

	container := &StorefrontsResponse{}
	if err := json.Unmarshal(resBody, container); err != nil {
		return nil, err
	}

	return container.Data, nil
}

type StorefrontsResponse struct {
	Data []models.StoreFront `json:"data"`
}
