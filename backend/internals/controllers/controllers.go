package controllers

import (
	"artist-app-backend/internals/services"
	"artist-app-backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetArtists(c *fiber.Ctx) error {
	artists := services.ReadArtistFile()
	if artists == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't read artists.json",
		})
	}

	fmt.Println("Artists retrieved")
	c.Type("json")
	return c.Status(http.StatusOK).JSON(artists)
}

func GetSingleArtist(c *fiber.Ctx) error {
	id := c.Params("id")

	artists := services.ReadArtistFile()

	var foundArtist models.Artist

	for _, artist := range artists {
		if artist.ID == id {
			fmt.Println("Artist retrieved")
			foundArtist = artist
		}
	}
	c.Type("json")
	return c.Status(http.StatusOK).JSON(foundArtist)
}

func PostArtist(c *fiber.Ctx) error {

	artists := services.ReadArtistFile()

	if artists == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't read artists.json",
		})
	}

	var newArtist models.Artist

	if err := c.BodyParser(&newArtist); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	newArtist.ID = uuid.New().String()

	artists = append(artists, newArtist)

	file, err := json.MarshalIndent(artists, "", " ")

	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't parse artists.json",
		})
	}

	err = os.WriteFile("data.json", file, 0644)

	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't write to artists.json",
		})
	}

	fmt.Println("Artist created")
	c.Type("json")
	return c.Status(http.StatusCreated).JSON(artists)

}

func UpdateArtist(c *fiber.Ctx) error {
	id := c.Params("id")

	artists := services.ReadArtistFile()

	var updatedArtist models.Artist

	if err := c.BodyParser(&updatedArtist); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	updatedArtist.ID = id

	for i, artist := range artists {
		if artist.ID == id {
			artists[i] = updatedArtist
			break
		}
	}

	file, err := json.MarshalIndent(artists, "", " ")

	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't parse artists.json",
		})
	}

	err = os.WriteFile("data.json", file, 0644)

	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't write to artists.json",
		})
	}

	fmt.Println("Artist updated")
	c.Type("json")
	return c.Status(http.StatusOK).JSON(artists)
}

func DeleteArtist(c *fiber.Ctx) error {
	id := c.Params("id")

	artists := services.ReadArtistFile()

	for i, artist := range artists {
		if artist.ID == id {
			artists = append(artists[:i], artists[i+1:]...)
			break
		}
	}

	file, err := json.MarshalIndent(artists, "", " ")

	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't parse artists.json",
		})
	}

	err = os.WriteFile("data.json", file, 0644)

	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't write to artists.json",
		})
	}

	fmt.Println("Artist deleted")
	c.Type("json")
	return c.Status(http.StatusOK).JSON(artists)
}
