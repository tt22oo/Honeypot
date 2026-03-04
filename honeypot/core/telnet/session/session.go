package session

import (
	"fmt"
	"honeypot/core/config"
	"honeypot/core/logger"
	"net"

	"github.com/google/uuid"
)

type Session struct {
	ID   string // session id
	User string // username
	Pass string // password
	IP   string // client's ip
	Conn net.Conn
}

func New(cfgs *config.Config, conn net.Conn) (*Session, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("UUID Generate Error: %s", err.Error())
	}

	session := &Session{
		ID:   id.String(),
		IP:   conn.RemoteAddr().String(),
		Conn: conn,
	}
	logger.Session(cfgs, "telnet", session.IP, session.ID)

	return session, nil
}
