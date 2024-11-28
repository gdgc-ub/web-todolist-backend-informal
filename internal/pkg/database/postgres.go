package database

import (
	"fmt"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.Todo{},
	)
}

func NewPostgresDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalln("Error connecting to database: ", err)
		return nil
	}

	if err := autoMigrate(db); err != nil {
		log.Fatalln("Error migrating database: ", err)
		return nil
	}

	return db
}
