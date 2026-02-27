package config

import (
	"encoding/json"
	"os"
)

type Configs struct {
	Listen  string `json:"listen"`
	Key     string `json:"key"`
	MongoDB string `json:"mongo_db"`
	DBName  string `json:"database_name"`
}

func parseConfigs(file *os.File) (*Configs, error) {
	var cfgs Configs
	err := json.NewDecoder(file).Decode(&cfgs)
	if err != nil {
		return nil, err
	}

	return &cfgs, nil
}
