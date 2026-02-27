package logger

import (
	"fmt"
	"honeypot/core/config"
	"honeypot/core/logger/report"
	"log"
	"os"
	"time"
)

var (
	infoLogger    = log.New(os.Stdout, "\033[36m[INFO]\033[0m ", log.LstdFlags)
	sessionLogger = log.New(os.Stdout, "\033[36m[SESSION]\033[0m ", log.LstdFlags)
	loginLogger   = log.New(os.Stdout, "\033[36m[LOGIN]\033[0m ", log.LstdFlags)
	cmdLogger     = log.New(os.Stdout, "\033[34m[CMD]\033[0m ", log.LstdFlags)
	errorLogger   = log.New(os.Stdout, "\033[31m[ERROR]\033[0m ", log.LstdFlags)
)

func Info(data string) {
	infoLogger.Print(data)
}

func Session(cfgs *config.Configs, protocol, ip, id string) {
	sessionLogger.Printf("protocol=%s, ip=%s, session_id=%s\r\n", protocol, ip, id)
	r := &report.ReportData{
		Name:      cfgs.Name,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		IP:        ip,
		Action:    "new_session",
		Protocol:  protocol,
		Data:      "",
		SessionID: id,
	}
	r.Report(cfgs)
}

func Login(cfgs *config.Configs, protocol, ip, id, user, pass string) {
	loginLogger.Printf("protocol=%s, username=%s, password=%s, session_id=%s\r\n", protocol, user, pass, id)
	r := &report.ReportData{
		Name:      cfgs.Name,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		IP:        ip,
		Action:    "login",
		Protocol:  protocol,
		Data:      fmt.Sprintf("%s:%s", user, pass),
		SessionID: id,
	}
	r.Report(cfgs)
}

func Cmd(cfgs *config.Configs, protocol, ip, id, cmd string) {
	cmdLogger.Printf("protocol=%s, cmd=%s, session_id=%s\r\n", protocol, cmd, id)
	r := &report.ReportData{
		Name:      cfgs.Name,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		IP:        ip,
		Action:    "cmd",
		Protocol:  protocol,
		Data:      cmd,
		SessionID: id,
	}
	r.Report(cfgs)
}

// print error log
func Error(protocol, data string) {
	errorLogger.Printf("protocol=%s, %s\r\n", protocol, data)
}
