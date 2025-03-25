package models

type StoreFront struct {
	Id         string               `json:"id"`
	Type       string               `json:"type"`
	Href       string               `json:"href"`
	Attributes StoreFrontAttributes `json:"attributes"`
}

type StoreFrontAttributes struct {
	DefaultLanguageTag    string   `json:"defaultLanguageTag"`
	ExplicitContentPolicy string   `json:"explicitContentPolicy"`
	Name                  string   `json:"name"`
	SupportedLanguageTags []string `json:"supportedLanguageTags"`
}
