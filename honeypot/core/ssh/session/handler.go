package session

import (
	"honeypot/core/logger"

	"github.com/gliderlabs/ssh"
)

func (c *Config) Handler(s ssh.Session) {
	defer s.Close()

	sessionID := s.Context().Value("sessionID").(string)
	for {
		cmd, err := WriteAndRead(s, "# ")
		if err != nil {
			logger.Error("ssh", err.Error())
			return
		}

		logger.Cmd(c.Configs, "ssh", s.RemoteAddr().String(), sessionID, cmd)
	}
}
