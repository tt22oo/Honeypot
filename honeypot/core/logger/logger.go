package logger

import (
	"fmt"
	"honeypot/core/config"
	"honeypot/core/logger/report"
	"log"
	"time"
)

const (
	TypeSession string = "\033[36m[SESSION]\033[0m"
	TypeLogin   string = "\033[36m[LOGIN]\033[0m"
	TypeCMD     string = "\033[34m[CMD]\033[0m"
	TypeError   string = "\033[31m[ERROR]\033[0m"
)

func Session(configs *config.Configs, protocol, ip, id string) {
	log.Printf("%s protocol=%s, ip=%s, session_id=%s)\r\n", TypeSession, protocol, ip, id)
	r := &report.ReportData{
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		IP:        ip,
		Action:    "new_session",
		Protocol:  protocol,
		Data:      "",
		SessionID: id,
	}
	r.Report(configs)
}

func Login(configs *config.Configs, protocol, ip, id, user, pass string) {
	log.Printf("%s protocol=%s, username=%s, password=%s, session_id=%s)\r\n", TypeLogin, protocol, user, pass, id)
	r := &report.ReportData{
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		IP:        ip,
		Action:    "login",
		Protocol:  protocol,
		Data:      fmt.Sprintf("%s:%s", user, pass),
		SessionID: id,
	}
	r.Report(configs)
}

func Cmd(configs *config.Configs, protocol, ip, id, cmd string) {
	log.Printf("%s protocol=%s, cmd=%s, session_id=%s)\r\n", TypeCMD, protocol, cmd, id)
	r := &report.ReportData{
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		IP:        ip,
		Action:    "cmd",
		Protocol:  protocol,
		Data:      cmd,
		SessionID: id,
	}
	r.Report(configs)
}

// print error log
func Error(protocol, data string) {
	log.Printf("%s protocol=%s, %s\r\n", TypeError, protocol, data)
}
