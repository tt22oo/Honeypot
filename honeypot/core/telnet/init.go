package telnet

import (
	"honeypot/core/config"
)

func Init(config *config.Configs) {
	if config.Telnet.Start {
		for _, listen := range config.Telnet.Listens {
			go Listen(config, listen)
		}
	}
}
