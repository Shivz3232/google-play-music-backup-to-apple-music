package models

type Song struct {
	Id         string         `json:"id"`
	Type       string         `json:"type"`
	Href       string         `json:"href"`
	Attributes SongAttributes `json:"attributes"`
}

type SongAttributes struct {
	AlbumName            string    `json:"albumName"`               // (Required)
	ArtistName           string    `json:"artistName"`              // (Required)
	ArtistURL            string    `json:"href,omitempty"`          // (Extended) (Artist URL from the "href" field)
	Artwork              Artwork   `json:"artwork"`                 // (Required)
	Attribution          string    `json:"attribution,omitempty"`   // (Classical music only)
	AudioVariants        []string  `json:"audioVariants,omitempty"` // (Extended)
	ComposerName         string    `json:"composerName,omitempty"`
	ContentRating        string    `json:"contentRating,omitempty"` // Possible values: clean, explicit
	DiscNumber           int       `json:"discNumber,omitempty"`
	DurationInMillis     int       `json:"durationInMillis"`     // (Required)
	GenreNames           []string  `json:"genreNames"`           // (Required)
	HasLyrics            bool      `json:"hasLyrics"`            // (Required)
	IsAppleDigitalMaster bool      `json:"isAppleDigitalMaster"` // (Required)
	ISRC                 string    `json:"isrc"`
	MovementCount        int       `json:"movementCount,omitempty"`  // (Classical music only)
	MovementName         string    `json:"movementName,omitempty"`   // (Classical music only)
	MovementNumber       int       `json:"movementNumber,omitempty"` // (Classical music only)
	Name                 string    `json:"name"`                     // (Required)
	ReleaseDate          string    `json:"releaseDate"`              // YYYY-MM-DD or YYYY
	TrackNumber          int       `json:"trackNumber"`              // Track number in the album
	URL                  string    `json:"url"`                      // (Required)
	WorkName             string    `json:"workName,omitempty"`       // (Classical music only)
	Previews             []Preview `json:"previews,omitempty"`       // (New field for previews)
}

type Preview struct {
	URL string `json:"url"`
}
