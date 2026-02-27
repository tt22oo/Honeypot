package main

import (
	"fmt"
	"honeypot/core/config"
	"honeypot/core/logger"
	"honeypot/core/ssh"
	"honeypot/core/telnet"
)

func main() {
	cfgs, err := config.Read()
	if err != nil {
		data := fmt.Sprintf("Config Error: %s", err.Error())
		logger.Error("telnet", data)
		return
	}

	ssh.Init(cfgs)
	telnet.Init(cfgs)

	select {}
}
