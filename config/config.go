package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Inicialize SQLite
	db, err = InicializeSQLite()

	if err != nil {
		return fmt.Errorf("error inicializing sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Inicialize Logger
	logger = NewLogger(p)
	return logger
}
