package config

import (
	"io/ioutil"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Config data
type Config struct {
	DBHost       string `yaml:"DBHOST"`
	DBName       string `yaml:"DBNAME"`
	DBUsername   string `yaml:"DBUSERNAME"`
	DBPassword   string `yaml:"DBPASSWORD"`
	Port         int    `yaml:"PORT"`
	DBConnParams string `yaml:"DBCONNPARAMS"`
	DBProtocol   string `yaml:"DBPROTOCOL"`
}

// GetPort return port for Server
func (cfg *Config) GetPort() string {
	return ":" + strconv.Itoa(cfg.Port)
}

// InitConfiguration initilize configuration
func InitConfiguration() (*Config, error) {
	cfg := Config{}
	configYml, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(configYml, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
