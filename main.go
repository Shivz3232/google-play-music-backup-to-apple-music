package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	api "shivu/google-play-music-backup-reader/apis"
	"shivu/google-play-music-backup-reader/apis/playlist"
	"shivu/google-play-music-backup-reader/apis/storefront"
	"shivu/google-play-music-backup-reader/auth"
	"shivu/google-play-music-backup-reader/models"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatalln("Failed to read env variables")
		}
	}

	// Check if directory path is provided
	if len(os.Args) < 2 {
		log.Fatalln("Please provide directory path")
	}

	if os.Getenv("APPLE_MUSIC_DEVELOPER_TOKEN") == "" {
		os.Setenv("APPLE_MUSIC_DEVELOPER_TOKEN", getDeveloperToken())
	}

	developerToken := os.Getenv("APPLE_MUSIC_DEVELOPER_TOKEN")
	musicUserToken := os.Getenv("APPLE_MUSIC_USER_TOKEN")

	if res, err := api.Test(developerToken); err != nil || res.StatusCode != 200 {
		log.Fatalf("Unable to reach Apple endpoints: %s\n", err)
	}

	directoryPath := os.Args[1]

	playlistName := "Google Play Playlist: " + getDirectoryName(directoryPath)

	pl, err := playlist.Create(developerToken, musicUserToken, playlistName, "")
	if err != nil {
		fmt.Printf("Failed creating playlist with name: %s\n", playlistName)
		fmt.Println(err)
		return
	}
	fmt.Printf("Created playlist: %s\n", playlistName)

	// Get all files in the directory
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		log.Fatalln("Error reading directory")
	}

	storefront, err := getStoreFront(developerToken, musicUserToken)
	if err != nil {
		log.Fatalf("Failed retrieving storefront: %s\n", err)
	} else if storefront == nil {
		log.Fatalln("Couldn't find the United States storefront")
	}

	fmt.Printf("Inserting %d tracks\n", len(files))
	for _, file := range files {
		fileName := file.Name()
		songName := fileName[:len(fileName)-len(filepath.Ext(fileName))]

		fmt.Printf("Searching song: %s\n", songName)
		songs, err := api.SearchSong(developerToken, musicUserToken, storefront.Id, songName, "songs")
		if err != nil {
			fmt.Printf("Error searching for song with name: %s\n", songName)
			log.Fatalln(err)
		}

		if len(songs) == 0 {
			fmt.Printf("Couldn't find song: %s\n", songName)
			continue
		}

		fmt.Printf("Inserting song: %s\n", songs[0].Attributes.Name)
		err = playlist.InsertSong(developerToken, musicUserToken, pl.Id, songs[0].Id)
		if err != nil {
			fmt.Printf("Error inserting song into playlist: %s\n", songs[0].Id)
			log.Fatalln(err)
		}
	}
}

func getDeveloperToken() string {
	jwt := &models.DeveloperToken{
		KeyID:          os.Getenv("APPLE_MUSIC_KEY_ID"),
		TeamID:         os.Getenv("APPLE_MUSIC_TEAM_ID"),
		PrivateKeyPath: os.Getenv("APPLE_MUSIC_PRIVATE_KEY_PATH"),
	}

	developerToken, err := auth.GenrateDeveloperToken(jwt, time.Hour*24)
	if err != nil {
		log.Fatalf("Error generating developer token: %s\n", err)
	}

	return developerToken
}

func getDirectoryName(directoryPath string) string {
	arr := strings.Split(directoryPath, "/")

	length := len(arr)

	return arr[length-1]
}

func getStoreFront(developerToken string, musicUserToken string) (*models.StoreFront, error) {
	storeFronts, err := storefront.GetAll(developerToken, musicUserToken)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(storeFronts); i++ {
		if storeFronts[i].Id == "in" {
			return &storeFronts[i], nil
		}
	}

	return nil, nil
}
