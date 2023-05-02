package config

import (
	"fmt"
	"github.com/LittleMikle/ozon/storage"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

//mem if false Postgres if true

var PgOrMem bool

func CreateConfig() (Conf *storage.Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("failed with creating config: %w", err)
	} else {
		log.Info().Msg("Config created successfully")
	}
	if os.Getenv("DB_PG") == "true" {
		PgOrMem = true
	}

	if os.Getenv("DB_PG") != "true" && os.Getenv("DB_PG") != "false" {
		log.Fatal().Msg("wrong .env DB_PG value: %w")
	}

	Conf = &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	return Conf, nil
}
