package main

import (
	"fmt"
	"honeypot/core/config"
	"honeypot/core/logger"
	"honeypot/core/telnet"
)

func main() {
	config, err := config.Read()
	if err != nil {
		data := fmt.Sprintf("Config Error: %s", err.Error())
		logger.Error("telnet", data)
		return
	}

	telnet.Init(config)

	select {}
}
