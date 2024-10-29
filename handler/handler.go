package handler

import (
	"github.com/Devtino7/goapi/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InicializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
