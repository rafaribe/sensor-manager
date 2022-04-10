package config

import (
	"io"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"

	. "github.com/rafaribe/planetwatch-awair-uploader/internal/logger"
)

type Config struct {
	Sensors []sensor `yaml:"sensors"`
	Store   store    `yaml:"store"`
}
type sensor struct {
	Model    string `yaml:"model"`
	Name     string `yaml:"name"`
	Endpoint string `yaml:"endpoint"`
}

// Type of store must be a pointer because it is optional
// https://github.com/go-yaml/yaml/issues/505#issuecomment-538453157
type store struct {
	Type     string    `yaml:"type"`
	Postgres *postgres `yaml:"postgres,omitempty"`
	InfluxDb *influxdb `yaml:"influxdb,omitempty"`
}
type postgres struct {
	ConnectionString string `yaml:"connection_string"`
}
type influxdb struct {
	host  string `yaml:"host"`
	token string `yaml:"token"`
}

func openAndReadFile(fileName string) (*Config, error) {
	log := zap.S()
	file, err := os.Open(fileName)
	if err != nil {
		log.Errorf("Failed to open file: %s", fileName)
	}
	cfg, err := unmarshalConfig(file)
	if err != nil {
		log.Errorf("Failed to read file: %s", fileName)
	}
	return cfg, nil
}

func unmarshalConfig(reader io.Reader) (*Config, error) {
	log := zap.S()
	cfg := &Config{}
	yd := yaml.NewDecoder(reader)
	err := yd.Decode(cfg)

	if err != nil {
		log.Errorf("Failed to unmarshal config: %s", err)
		return nil, err

	}

	return cfg, nil
}

func ParseConfiguration() *Config {
	// Initialize Logger
	logger := InitZapLog()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
	log := zap.S()
	configFilePath := os.Getenv("CONFIG_PATH")
	if configFilePath == "" {
		configFilePath = "~/.config.yaml"
		log.Infof("Loaded config at default location %s", configFilePath)
	} else {
		log.Infof("Loaded configuration file from non-default location %s", configFilePath)
	}
	cfg, err := openAndReadFile(configFilePath)
	if err != nil {
		log.Errorw("Error parsing configuration file %s", err)
	}
	return cfg
}
