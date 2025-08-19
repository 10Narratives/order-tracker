package ordersapp

import (
	pgapp "github.com/10Narratives/golang-toolkit/components/databases/postgres"
	grpcsrv "github.com/10Narratives/golang-toolkit/components/network/grpc/server"
)

type Config struct {
	Grpc     *grpcsrv.Config `yaml:"grpc" env-required:"true"`
	Database *pgapp.Config   `yaml:"database" env-required:"true"`
}
