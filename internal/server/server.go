package server

import (
	"fmt"
	"github.com/LittleMikle/ozon/config"
	"github.com/LittleMikle/ozon/internal/handlers"
	"github.com/LittleMikle/ozon/internal/models"
	"github.com/LittleMikle/ozon/storage"
	"github.com/gofiber/fiber/v2"
)

func Run() error {
	conf, err := config.CreateConfig()
	if err != nil {
		return fmt.Errorf("failed to create a Config: %w", err)
	}

	app := fiber.New()

	if config.PgOrMem == true {

		db, err := storage.NewConnectionPostgres(conf)
		if err != nil {
			return fmt.Errorf("failed to connect to Postgres: %w", err)
		}

		err = models.MigrateURLs(db)
		if err != nil {
			return fmt.Errorf("failed to migrate URLs: %w", err)
		}

		r := handlers.Repository{
			DB: db,
		}

		r.SetupRoutes(app)
	} else {
		storage.SetupRoutes(app)
	}

	err = app.Listen(":8080")
	if err != nil {
		return fmt.Errorf("failed to listen server: %w", err)
	}
	return nil
}
