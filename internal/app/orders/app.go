package ordersapp

import (
	pgapp "github.com/10Narratives/golang-toolkit/components/databases/postgres"
	grpcsrv "github.com/10Narratives/golang-toolkit/components/network/grpc/server"
)

type App struct {
	grpc     *grpcsrv.App
	postgres *pgapp.App
}

func New() (*App, error) {
	return nil, nil
}

func Start() error {
	return nil
}
