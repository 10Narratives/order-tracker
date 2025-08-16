package transportcfg

import "time"

type Config struct {
	Host    string        `yaml:"host" env-required:"true"`
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-default:"10s"`
}
