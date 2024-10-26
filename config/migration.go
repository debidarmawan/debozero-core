package config

import (
	"log"

	"debozero-core/model"

	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Oauth2Client{},
		&model.RoleKey{},
		&model.RoleKeyMapping{},
	)

	if err != nil {
		log.Fatal(err)
	}
}
