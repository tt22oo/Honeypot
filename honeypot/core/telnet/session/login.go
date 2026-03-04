package session

import (
	"honeypot/core/config"
	"honeypot/core/logger"
)

func (s *Session) Login(cfgs *config.Config) error {
	var err error
	s.User, err = s.WriteAndRead("Username: ")
	if err != nil {
		return err
	}

	s.Pass, err = s.WriteAndRead("Password: ")
	if err != nil {
		return err
	}

	logger.Login(cfgs, "telnet", s.IP, s.ID, s.User, s.Pass)

	return nil
}
