package telnet

import (
	"honeypot/core/config"
)

func Init(cfgs *config.Config) {
	if cfgs.Telnet.Start {
		for _, addr := range cfgs.Telnet.Addrs {
			go Listen(cfgs, addr)
		}
	}
}
