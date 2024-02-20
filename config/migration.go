package config

import (
	"log"

	"github.com/debidarmawan/debozero-core/model"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Oauth2Client{},
	)

	if err != nil {
		log.Fatal(err)
	}
}
