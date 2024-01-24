package services

import (
	"artist-app-backend/models"
	"encoding/json"
	"fmt"
	"os"
)

func ReadArtistFile() []models.Artist {
	data, err := os.ReadFile("data.json")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var artists []models.Artist

	err = json.Unmarshal(data, &artists)

	if err != nil {
		fmt.Println(err)
		return nil

	}

	return artists
}
