package telnet

import (
	"fmt"
	"honeypot/core/config"
	"honeypot/core/logger"
	"honeypot/core/telnet/session"
	"net"
)

// listen telnet
func Listen(config *config.Configs, host string) {
	ln, err := net.Listen("tcp", host)
	if err != nil {
		data := fmt.Sprintf("Listen Error: %s", err.Error())
		logger.Error("telnet", data)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			data := fmt.Sprintf("Accept Error: %s", err.Error())
			logger.Error("telnet", data)
			continue
		}

		s, err := session.New(config, conn)
		if err != nil {
			conn.Close()
			logger.Error("telnet", err.Error())
			continue
		}

		go s.Handler(config)
	}
}
