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

type ShellConfig struct {
	HomePath    string `json:"home_directory"`
	DirPath     string `json:"directory_path"`
	ProcessPath string `json:"process_path"`
	CpuinfoPath string `json:"cpuinfo_path"`
	MeminfoPath string `json:"meminfo_path"`
	VersionPath string `json:"version_path"`
}

type Config struct {
	Name      string      `json:"name"`
	Key       string      `json:"key"`
	ReportURL string      `json:"report_url"`
	Telnet    Telnet      `json:"telnet"`
	SSH       SSH         `json:"ssh"`
	Shell     ShellConfig `json:"shell"`
}

func parseConfigs(file *os.File) (*Config, error) {
	var cfgs Config
	err := json.NewDecoder(file).Decode(&cfgs)
	if err != nil {
		return nil, err
	}

	return &cfgs, nil
}
