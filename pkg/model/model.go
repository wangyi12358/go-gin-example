package model

import (
	"fmt"
	"go-gin-example/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func Setup() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Config.Database.Host,
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Name,
		config.Config.Database.Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

}
