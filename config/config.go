package config

import (
	"io"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"

	. "github.com/rafaribe/sensor-manager/internal/logger"
)

type Config struct {
	Sensors []Sensor `yaml:"sensors"`
	Store   Store    `yaml:"store"`
}
type Sensor struct {
	Model    string `yaml:"model"`
	Name     string `yaml:"name"`
	Endpoint string `yaml:"endpoint"`
}

// Type of store must be a pointer because it is optional
// https://github.com/go-yaml/yaml/issues/505#issuecomment-538453157
type Store struct {
	Type     string    `yaml:"type"`
	Postgres *Postgres `yaml:"postgres,omitempty"`
	InfluxDb *Influxdb `yaml:"influxdb,omitempty"`
}
type Postgres struct {
	ConnectionString string `yaml:"connection_string"`
}
type Influxdb struct {
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
