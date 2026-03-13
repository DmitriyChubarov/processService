package main

import (
	"github.com/labstack/gommon/log"

	"github.com/DmitriyChubarov/processing/pkg/postgresql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config := postgresql.LoadConfig()
	m, err := migrate.New(
		"file://migrations",
		config.PostgresDSN,
	)
	if err != nil {
		log.Fatal("failed to initialize migrations ", err)
	}

	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("error applying migration", err)
	}

	log.Info("migrations created")
}
