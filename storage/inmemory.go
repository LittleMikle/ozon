package storage

import (
	"fmt"
	"github.com/LittleMikle/ozon/tools"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type URL struct {
	Url     string `json:"url"`
	TinyURL string `json:"tiny_url"`
}

var mapURL map[string]string

func SetupRoutes(app *fiber.App) {
	mapURL = map[string]string{}
	api := app.Group("/api")
	api.Post("/create", CreateUrl)
	api.Get("/find/:tiny_url", GetURL)
}

func CreateUrl(context *fiber.Ctx) error {
	Url := URL{}
	tinyURL, err := tools.GenerateToken()
	if err != nil {
		return err
	}

	Url.TinyURL = tinyURL
	err = context.BodyParser(&Url)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	mapURL[Url.TinyURL] = Url.Url

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "url has been added",
		"tinyURL": Url.TinyURL})

	return nil
}

func GetURL(context *fiber.Ctx) error {
	tinyURL := context.Params("tiny_url")

	Z := ""
	if x, found := mapURL[tinyURL]; found {
		Z = x
	}
	fmt.Println(mapURL)
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":  "Full URL fetched successfully",
		"full url": Z,
	})

	return nil
}
