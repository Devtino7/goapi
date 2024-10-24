package main

import (
	"github.com/Devtino7/goapi/config"
	"github.com/Devtino7/goapi/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	// Inicialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config inicialization error: %v", err)
		return
	}
	// Inicialize router
	router.Inicialize()
}
