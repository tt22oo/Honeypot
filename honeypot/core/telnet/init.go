package telnet

import (
	"honeypot/core/config"
)

func Init(configs *config.Configs) {
	if configs.Telnet.Start {
		for _, listen := range configs.Telnet.Listens {
			go Listen(configs, listen)
		}
	}
}
