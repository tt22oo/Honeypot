package session

import (
	"honeypot/core/config"
	"honeypot/core/logger"
)

func (s *Session) Handler(cfgs *config.Configs) {
	defer s.Conn.Close()

	err := s.Login(cfgs)
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

		logger.Cmd(cfgs, "telnet", s.IP, s.ID, input)
	}
}
