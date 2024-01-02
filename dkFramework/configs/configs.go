package configs

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
)

type Config struct {
	Logging struct {
		OutputType string `yaml:"output_type"`
		OutputPath string `yaml:"output_path"`
		MaxSizeMB  int    `yaml:"max_size_mb"`
	} `yaml:"logging"`
	DNS struct {
		Port            int           `yaml:"port"`
		RecursiveDNS    string        `yaml:"recursive_dns"`
		Whitelist       string        `yaml:"whitelist"`
		CacheTimeoutStr string        `yaml:"cache_timeout"`
		CacheTimeout    time.Duration // Add cache timeout configuration
	} `yaml:"dns"`
	HTTPS struct {
		Port       int    `yaml:"port"`
		ServerPath string `yaml:"server_path"`
	} `yaml:"https"`
}

var AppConfigInstance *Config

func LoadConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	//转化CacheTimeOut
	if config.DNS.CacheTimeoutStr != "" {
		duration, err := time.ParseDuration(config.DNS.CacheTimeoutStr)
		if err != nil {
			return config, err
		}
		config.DNS.CacheTimeout = duration
	}

	AppConfigInstance = config
	return config, nil
}
