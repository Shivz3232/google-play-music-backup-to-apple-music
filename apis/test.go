package api

import (
	"errors"
	"net/http"
)

func Test(developerToken string) (*http.Response, error) {
	if developerToken == "" {
		return nil, errors.New("Token is empty")
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.music.apple.com/v1/test", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+developerToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
