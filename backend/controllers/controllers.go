package controllers

import (
	"artist-app-backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetArtists(c *fiber.Ctx) error {
	data, err := os.ReadFile("data.json")

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't read artists.json",
		})
	}

	var artists []models.Artist

	err = json.Unmarshal(data, &artists)

	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't parse artists.json",
		})
	}

	c.Type("json")
	return c.Status(http.StatusOK).JSON(artists)
}

func PostArtist(c *fiber.Ctx) error {
	data, err := os.ReadFile("data.json")

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't read artists.json",
		})
	}

	var artists []models.Artist

	err = json.Unmarshal(data, &artists)

	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't parse artists.json",
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

	c.Type("json")
	return c.Status(http.StatusCreated).JSON(artists)

}

func UpdateArtist(c *fiber.Ctx) error {

	id := c.Params("id")

	data, err := os.ReadFile("data.json")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't read artists.json",
		})
	}

	var artists []models.Artist

	err = json.Unmarshal(data, &artists)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't parse artists.json",
		})
	}

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

	c.Type("json")
	return c.Status(http.StatusOK).JSON(artists)
}
