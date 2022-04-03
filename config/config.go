package config

import (
	"io"
	"os"

	. "github.com/rafaribe/planetwatch-awair-uploader/internal/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Inputs  Inputs  `yaml:"inputs"`
	Outputs Outputs `yaml:"outputs"`
	Store   Store   `yaml:"store"`
}
type Inputs struct {
	LocalAwair []string `yaml:"localAwair"`
}
type PlanetWatch struct {
	Endpoint     string `yaml:"endpoint"`
	AccessToken  string `yaml:"access_token"`
	RefreshToken string `yaml:"refresh_token"`
}
type Outputs struct {
	PlanetWatch PlanetWatch `yaml:"planetWatch"`
}
type Firebase struct {
	SecureCredentialsFile string `yaml:"secureCredentialsFile"`
}
type Store struct {
	Firebase Firebase `yaml:"firebase"`
}

func OpenAndReadFile(fileName string) (*Config, error) {
	log := zap.S()
	file, err := os.Open(fileName)
	if err != nil {
		log.Errorf("Failed to open file: %s", fileName)
	}
	cfg, err := UnmarshalConfig(file)
	if err != nil {
		log.Errorf("Failed to read file: %s", fileName)
	}
	return cfg, nil
}

func UnmarshalConfig(reader io.Reader) (*Config, error) {
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
	configFilePath := os.Getenv("CONFIG_FILE")
	if configFilePath == "" {
		configFilePath = "~/.config.yaml"
		log.Infof("Loaded config at default location %s", configFilePath)
	} else {
		log.Infof("Loaded configuration file from non-default location %s", configFilePath)
	}
	cfg, err := OpenAndReadFile(configFilePath)
	if err != nil {
		log.Errorw("Error parsing configuration file %s", err)
	}
	return cfg
}
