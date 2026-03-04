package session

import (
	"fmt"
	"honeypot/core/config"
	"honeypot/core/logger"

	"github.com/tt22oo/fakeshell/shell/command"
	"github.com/tt22oo/fakeshell/shell/parser"
)

func (s *Session) Handler(cfgs *config.Config) {
	defer s.Conn.Close()

	err := s.Login(cfgs)
	if err != nil {
		logger.Error("telnet", err.Error())
		return
	}

	shell, err := initShell(cfgs)
	if err != nil {
		logger.Error("telnet", err.Error())
		return
	}
	shell.User = s.User

	for {
		input, err := s.WriteAndRead("# ")
		if err != nil {
			logger.Error("telnet", err.Error())
			return
		}

		tokens := parser.Input(input)
		result := command.RunCommands(shell, tokens)
		_, err = s.Conn.Write([]byte(result))
		if err != nil {
			data := fmt.Sprintf("Write Error: %s", err.Error())
			logger.Error("telnet", data)
			return
		}

		logger.Cmd(cfgs, "telnet", s.IP, s.ID, input)
	}
}
