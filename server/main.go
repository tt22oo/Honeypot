package main

import (
	"fmt"
	"server/core/config"
	"server/core/database"
	"server/core/logger"
	"server/core/web"
)

func main() {
	cfgs, err := config.Read()
	if err != nil {
		data := fmt.Sprintf("Config Error: %s", err.Error())
		logger.Error(data)
		return
	}

	client, err := database.Connect(cfgs)
	if err != nil {
		data := fmt.Sprintf("MongoDB Error: %s", err.Error())
		logger.Error(data)
		return
	}

	web.Start(cfgs, client)
}
