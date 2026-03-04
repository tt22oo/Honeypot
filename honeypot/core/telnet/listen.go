package telnet

import (
	"fmt"
	"honeypot/core/config"
	"honeypot/core/logger"
	"honeypot/core/telnet/session"
	"net"
)

// listen telnet honeypot
func Listen(cfgs *config.Config, addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		data := fmt.Sprintf("Listen Error: %s", err.Error())
		logger.Error("telnet", data)
		return
	}
	defer ln.Close()

	logger.Info(fmt.Sprintf("Telnet Honeypot Listening on %s", addr))
	for {
		conn, err := ln.Accept()
		if err != nil {
			data := fmt.Sprintf("Accept Error: %s", err.Error())
			logger.Error("telnet", data)
			continue
		}

		s, err := session.New(cfgs, conn)
		if err != nil {
			conn.Close()
			logger.Error("telnet", err.Error())
			continue
		}

		go s.Handler(cfgs)
	}
}
