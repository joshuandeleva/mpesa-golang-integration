package db

import (
	"github.com/go-mpesa-integration/interfacex"
	"github.com/go-mpesa-integration/internals/model"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(
		&interfacex.STKPushRequest{},
		&model.CallbackRequest{},
	)
}