package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mehmonov/movies-crud/config"
	"github.com/mehmonov/movies-crud/internal/models"
)

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(
		&models.Movie{},
		&models.MovieMedia{},
		&models.MovieMetadata{},
		&models.User{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
