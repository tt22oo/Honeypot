package ssh

import (
	"fmt"
	"honeypot/core/config"
	"honeypot/core/logger"
	"honeypot/core/ssh/session"

	"github.com/gliderlabs/ssh"
)

// listen ssh honeypot
func Listen(cfgs *config.Configs, addr string) {
	c := &session.Config{
		Configs: cfgs,
	}

	server := &ssh.Server{
		Addr:            addr,
		Handler:         c.Handler,
		Version:         cfgs.SSH.Version,
		PasswordHandler: c.Login,
	}

	logger.Info(fmt.Sprintf("SSH Honeypot Listening on %s", addr))
	err := server.ListenAndServe()
	if err != nil {
		data := fmt.Sprintf("SSH Error: %s", err.Error())
		logger.Error("ssh", data)
	}
}
