package main

import (
	"github.com/10Narratives/golang-toolkit/pkg/logging"
	"github.com/10Narratives/order-tracker/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log, err := logging.New(
		logging.WithFormat(cfg.Logging.Format),
		logging.WithOutput(cfg.Logging.Output),
	)

	log.Error("starting orders service")
}
