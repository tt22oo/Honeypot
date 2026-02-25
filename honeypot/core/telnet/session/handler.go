package session

import (
	"honeypot/core/config"
	"honeypot/core/logger"
)

func (s *Session) Handler(config *config.Configs) {
	defer s.Conn.Close()

	err := s.Login(config)
	if err != nil {
		logger.Error("telnet", err.Error())
		return
	}

	for {
		input, err := s.WriteAndRead("# ")
		if err != nil {
			logger.Error("telnet", err.Error())
			return
		}

		logger.Cmd(config, "telnet", s.IP, s.ID, input)
	}
}
