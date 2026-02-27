package report

import (
	"bytes"
	"encoding/json"
	"fmt"
	"honeypot/core/config"
	"log"
	"net/http"
	"os"
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

	url := fmt.Sprintf("%s?key=%s", cfgs.ReportURL, cfgs.Key)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
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
