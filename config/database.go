package config

import (
	"log"

	"github.com/debidarmawan/debozero-core/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func ConnectDatabase(maxOpenConnection int) *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(GetEnv(constants.DbUrl)),
		&gorm.Config{
			Logger:         logger.Default.LogMode(logger.Silent),
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxOpenConns(maxOpenConnection)

	return db
}
