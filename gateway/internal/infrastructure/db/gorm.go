package database

import (
	"fmt"

	"github.com/superdev/ecommerce/gateway/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(conf *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.DBHost,
		conf.DBUser,
		conf.DBPassword,
		conf.DBName,
		conf.DBPort,
		// viper.GetString("DB_HOST"),
		// viper.GetString("DB_USER"),
		// viper.GetString("DB_PASSWORD"),
		// viper.GetString("DB_NAME"),
		// viper.GetString("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Auto-migrate User model
	if conf.AutoMigrate {
		//TODO: Create a migrator service with injectors
		if err = db.AutoMigrate(&models.User{}); err != nil {
			return nil, err
		}
	}

	return db, err
}
