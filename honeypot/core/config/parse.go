package config

import (
	"encoding/json"
	"os"
)

type Telnet struct {
	Start   bool     `json:"start"`
	Listens []string `json:"listens"`
}

type Configs struct {
	Name      string `json:"name"`
	ReportURL string `json:"report_url"`
	Telnet    Telnet `json:"telnet"`
}

func parseConfigs(file *os.File) (*Configs, error) {
	var configs Configs
	err := json.NewDecoder(file).Decode(&configs)
	if err != nil {
		return nil, err
	}

	return &configs, nil
}
