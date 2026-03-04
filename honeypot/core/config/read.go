package config

import (
	"os"
)

const configPATH string = "config/config.json"

func Read() (*Config, error) {
	file, err := os.Open(configPATH)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseConfigs(file)
}
