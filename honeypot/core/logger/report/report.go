package report

import (
	"bytes"
	"encoding/json"
	"honeypot/core/config"
	"log"
	"net/http"
	"os"
	"time"
)

type ReportData struct {
	Name      string `json:"name"`
	Time      string `json:"time"`
	IP        string `json:"ip"`
	Action    string `json:"action"`
	Protocol  string `json:"protocol"`
	SessionID string `json:"session_id"`
	Data      string `json:"data"`
}

type Response struct {
	Stat    string `json:"stat"`
	Message string `json:"message"`
}

var errorLogger = log.New(os.Stdout, "\033[31m[ERROR]\033[0m ", log.LstdFlags)

func (r *ReportData) Report(cfgs *config.Configs) {
	body, err := json.Marshal(r)
	if err != nil {
		errorLogger.Printf("JSON Marshal Error: %s", err.Error())
		return
	}

	req, err := http.NewRequest("POST", cfgs.ReportURL, bytes.NewBuffer(body))
	if err != nil {
		errorLogger.Printf("New Request Error: %s", err.Error())
		return
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("key", cfgs.Key)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		errorLogger.Printf("Report Error: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	var rp Response
	err = json.NewDecoder(resp.Body).Decode(&rp)
	if err != nil {
		errorLogger.Printf("JSON Decode Error: %s", err.Error())
		return
	}

	if rp.Stat != "success" {
		errorLogger.Printf("Report Error: %s", rp.Message)
	}
}
