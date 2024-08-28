package utils

import (
	"sync"
)

type Config struct {
	RulesFolder   string `yaml:"rulesFolder"`
	GeoIPFolder   string `yaml:"geoipFolder"`
	Elasticsearch string `yaml:"elasticsearch"`
	PostgreSQL    struct {
		Server   string `yaml:"server"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"postgresql"`
	ErrorLevel string `yaml:"errorLevel"`
	useCache string `yaml:"useCache"`
	useSystemRules string `yaml:"useSystemRules"`
	appendCompletedAlerts string `yaml:"appendCompletedAlerts"`
}

var oneConfigRead sync.Once
var cnf Config

func readConfig() {
	ReadYaml("config.yml", &cnf)
}

func GetConfig() Config {
	oneConfigRead.Do(func() { readConfig() })
	return cnf
}
