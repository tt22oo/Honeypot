package session

import (
	"fmt"
	"honeypot/core/logger"

	"github.com/gliderlabs/ssh"
	"github.com/tt22oo/fakeshell/shell/command"
	"github.com/tt22oo/fakeshell/shell/parser"
)

func (cfgs *Config) Handler(s ssh.Session) {
	defer s.Close()

	shell, err := cfgs.initShell()
	if err != nil {
		logger.Error("ssh", err.Error())
		return
	}
	shell.User = s.User()
	sessionID := s.Context().Value("sessionID").(string)

	for {
		input, err := WriteAndRead(s, "# ")
		if err != nil {
			logger.Error("ssh", err.Error())
			return
		}

		tokens := parser.Input(input)
		result := command.RunCommands(shell, tokens)
		_, err = s.Write([]byte(result))
		if err != nil {
			data := fmt.Sprintf("Write Error: %s", err.Error())
			logger.Error("ssh", data)
			return
		}

		logger.Cmd(cfgs.Config, "ssh", s.RemoteAddr().String(), sessionID, input)
	}
}
