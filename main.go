package main

import (
	"github.com/LuhanM/gopportunities/config"
	"github.com/LuhanM/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger := config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config intialization error: %v", err)
	}

	router.Initialize()
}
