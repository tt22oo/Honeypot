package session

import (
	"honeypot/core/logger"

	"github.com/gliderlabs/ssh"
)

func (cfgs *Config) Login(ctx ssh.Context, pass string) bool {
	id, err := NewID()
	if err != nil {
		logger.Error("ssh", err.Error())
		return false
	}

	ctx.SetValue("sessionID", id)

	logger.Session(cfgs.Config, "ssh", ctx.RemoteAddr().String(), id)
	logger.Login(cfgs.Config, "ssh", ctx.RemoteAddr().String(), id, ctx.User(), pass)

	return true
}
