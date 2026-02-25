package report

import (
	"bytes"
	"encoding/json"
	"honeypot/core/config"
	"log"
	"net/http"
)

type ReportData struct {
	Time      string `json:"time"`
	IP        string `json:"ip"`
	Action    string `json:"action"`
	Protocol  string `json:"protocol"`
	SessionID string `json:"session_id"`
	Data      string `json:"data"`
}

func (r *ReportData) Report(configs *config.Configs) {
	body, err := json.Marshal(r)
	if err != nil {
		log.Printf("\033[31m[ERROR]\033[0m JSON Marshal Error: %s", err.Error())
		return
	}

	resp, err := http.Post(configs.ReportURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("\033[31m[ERROR]\033[0m Report Error: %s", err.Error())
		return
	}
	defer resp.Body.Close()
}
