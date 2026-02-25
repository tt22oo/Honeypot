package config

import (
	"os"
)

const configPATH string = "configs.json"

func Read() (*Configs, error) {
	file, err := os.Open(configPATH)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseConfigs(file)
}
