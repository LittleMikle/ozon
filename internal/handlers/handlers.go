package handlers

import (
	"fmt"
	"github.com/LittleMikle/ozon/internal/models"
	"github.com/LittleMikle/ozon/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
)

type URL struct {
	URL     string `json:"url"`
	TinyURL string `json:"tiny_url"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create", r.CreateURL)
	api.Get("/find/:tiny_url", r.GetURL)
}

func (r *Repository) CreateURL(context *fiber.Ctx) error {
	url := URL{}
	tinyURL, err := tools.GenerateToken()
	if err != nil {
		return err
	}
	//	tinyURL = "localhost:8080/api/find/" + tinyURL
	url.TinyURL = tinyURL

	err = context.BodyParser(&url)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	fmt.Println(url)
	err = r.DB.Create(&url).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create an url"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "url has been added",
		"tinyURL": tinyURL})
	return nil
}

func (r *Repository) GetURL(context *fiber.Ctx) error {
	tinyURL := context.Params("tiny_url")
	urlModel := &models.URLs{}
	if tinyURL == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "url cannot be empty",
		})
		return nil
	}
	log.Info().Msgf("tinyURL is", tinyURL)

	err := r.DB.Where("tiny_url = ?", tinyURL).First(urlModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get the tinyURL",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "tinyURL fetched successfully",
		"data":    urlModel,
	})
	return nil
}
