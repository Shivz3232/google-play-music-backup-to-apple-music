package models

type Playlist struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Href       string `json:"href"`
	Attributes struct {
		PlayParams struct {
			Id        string `json:"id"`
			Kind      string `json:"kind"`
			IsLibrary bool   `json:"isLibrary"`
		} `json:"playParams"`
		Name       string `json:"name"`
		HasCatalog bool   `json:"hasCatalog"`
		CanEdit    bool   `json:"canEdit"`
		IsPublic   bool   `json:"isPublic"`
		DateAdded  string `json:"dateAdded"`
	} `json:"attributes"`
}
