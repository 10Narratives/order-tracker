package config

import (
	"github.com/10Narratives/golang-toolkit/pkg/config/loader"
	"github.com/10Narratives/golang-toolkit/pkg/logging"
	databasecfg "github.com/10Narratives/order-tracker/internal/config/database"
	transportcfg "github.com/10Narratives/order-tracker/internal/config/transport"
)

type Config struct {
	Transport transportcfg.Config   `yaml:"transport"`
	Database  databasecfg.Config    `yaml:"database"`
	Logging   logging.LoggingConfig `yaml:"logging"`
}

var l = loader.ConfigLoader[Config]{}

func Load() (*Config, error) {
	return l.Load()
}

func LoadFromFile(path string) (*Config, error) {
	return l.LoadFromFile(path)
}
