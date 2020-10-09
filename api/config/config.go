package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// BotConfig is a type that holds the config for the app.
type BotConfig struct {
	DB struct {
		Name    string            `yaml:"name"`
		Mode    os.FileMode       `yaml:"mode"`
		Buckets map[string]string `yaml:"buckets"`
	} `yaml:"db"`
	Logging struct {
		Path string `yaml:"path"`
	} `yaml:"logging"`
	App struct {
		Name   string            `yaml:"name"`
		Assets map[string]string `yaml:"assets"`
	} `yaml:"app"`
}

// New makes a new instance of BotConfig.
func New(cfgPath string) (*BotConfig, error) {
	cfg := &BotConfig{}

	f, err := os.Open(cfgPath)

	if err != nil {
		log.Fatalf("Error opening config file sat %s: %v", cfgPath, err)
		return nil, err
	}

	defer f.Close()

	d := yaml.NewDecoder(f)

	if err := d.Decode(&cfg); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
		return nil, err
	}

	return cfg, nil
}
