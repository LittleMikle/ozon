package storage

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
	SSLMode  string
}

func NewConnectionPostgres(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s",
		config.User, config.Password, config.DBname, config.Host, config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf("failed with gorm connection: %w", err)
	} else {
		log.Info().Msg("Connection to Postgres successful")
	}

	return db, nil
}
