package ssh

import (
	"honeypot/core/config"
)

func Init(cfgs *config.Config) {
	if cfgs.SSH.Start {
		for _, addr := range cfgs.SSH.Addrs {
			go Listen(cfgs, addr)
		}
	}
}
