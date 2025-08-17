package database

import (
	"L0_task/internal/config"
	"L0_task/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var toAutoMigrate = []interface{}{
	&models.Delivery{},
	&models.Order{},
	&models.Payment{},
	&models.Item{},
}

func buildDSN(cfg *config.Config) string {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
	)
	return dsn
}

func MustLoad(config *config.Config) *gorm.DB {
	dsn := buildDSN(config)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	if err := db.AutoMigrate(toAutoMigrate...); err != nil {
		panic(fmt.Errorf("failed to auto migrate database: %w", err))
	}
	return db
}
