package config

import (
	"encoding/json"
	"os"
)

type ConfigStruct struct {
	MattermostToken  string            `json:"mattermost_token"`
	Usernames        map[string]string `json:"usernames,omitempty"`
	Host             string            `json:"host,omitempty"`
	Port             uint16            `json:"port,omitempty"`
	Language         string            `json:"language,omitempty"`
	MattermostServer string            `json:"mattermost_server"`
}

var Config ConfigStruct
var configParsed bool = false

func Parse() (cfg ConfigStruct, err error) {
	configBytes, err := os.ReadFile("config.json")
	if err != nil {
		return ConfigStruct{}, err
	}
	err = json.Unmarshal(configBytes, &cfg)
	if err != nil {
		return ConfigStruct{}, err
	}
	if !configParsed {
		Config = cfg
		configParsed = true
	}
	return cfg, nil
}
