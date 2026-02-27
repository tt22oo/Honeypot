package config

import (
	"encoding/json"
	"os"
)

type SSH struct {
	Start   bool     `json:"start"`
	Version string   `json:"version"`
	Addrs   []string `json:"addrs"`
}

type Telnet struct {
	Start bool     `json:"start"`
	Addrs []string `json:"addrs"`
}

type Configs struct {
	Name      string `json:"name"`
	Key       string `json:"key"`
	ReportURL string `json:"report_url"`
	Telnet    Telnet `json:"telnet"`
	SSH       SSH    `json:"ssh"`
}

func parseConfigs(file *os.File) (*Configs, error) {
	var cfgs Configs
	err := json.NewDecoder(file).Decode(&cfgs)
	if err != nil {
		return nil, err
	}

	return &cfgs, nil
}
