package database

import (
	"fmt"
	"go_data_gouv_client/config"
	"go_data_gouv_client/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.DBUser, cfg.Database.DBPass, cfg.Database.DBHost, cfg.Database.DBPort, cfg.Database.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Echec connexion MySQL:", err)
	}

	db.AutoMigrate(&models.GovData{})

	log.Println("Connexion MySQL réussie et migration effectuée.")
	return db
}
