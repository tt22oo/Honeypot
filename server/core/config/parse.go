package config

import (
	"encoding/json"
	"os"
)

type Configs struct {
	Listen  string `json:"listen"`
	MongoDB string `json:"mongo_db"`
	DBName  string `json:"database_name"`
}

func parseConfigs(file *os.File) (*Configs, error) {
	var configs Configs
	err := json.NewDecoder(file).Decode(&configs)
	if err != nil {
		return nil, err
	}

	return &configs, nil
}
