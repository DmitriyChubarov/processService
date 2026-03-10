package app

import (
	"github.com/DmitriyChubarov/processing/internal/http"
	"github.com/DmitriyChubarov/processing/internal/logic"
	"github.com/DmitriyChubarov/processing/internal/postgres"
	"github.com/DmitriyChubarov/processing/pkg/postgresql"
	"github.com/gocraft/dbr/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Run(serviceName string) {
	log.Info("start service: ", serviceName)
	config := postgresql.LoadConfig()
	log.Info("database connection")
	connection, err := dbr.Open("pgx", config.PostgresDSN, nil)
	if err != nil {
		log.Fatal("database connection error ", err)
	}
	defer connection.Close()

	session := connection.NewSession(nil)
	processRepository := postgres.NewProcessRepository(session)
	processLogic := logic.NewProcessLogic(processRepository)
	processHandlers := http.NewProcessingHandlers(processLogic)

	e := echo.New()

	e.HTTPErrorHandler = http.HTTPErrorHandler

	processHandlers.Register(e)

	log.Infof("service started: %s", serviceName)
	e.Logger.Fatal(e.Start(":" + config.HTTPPort))
	log.Info("service stopped")
}
