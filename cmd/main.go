package main

import (
	"errors"
	"net/http"

	"github.com/Arijeet-webonise/chatTest/app"
	"github.com/Arijeet-webonise/chatTest/app/config"
	"github.com/Arijeet-webonise/chatTest/pkg/database"
	"github.com/Arijeet-webonise/chatTest/pkg/session"
	"github.com/Arijeet-webonise/chatTest/pkg/templates"
	"github.com/azer/logger"
	"github.com/go-zoo/bone"
)

func main() {
	cfg, cfgErr := config.InitConfiguration()
	if cfgErr != nil {
		panic(cfgErr)
	}

	db := &database.DatabaseConfig{
		Protocol:         cfg.DBProtocol,
		Username:         cfg.DBUsername,
		Password:         cfg.DBPassword,
		Host:             cfg.DBHost,
		DatabaseName:     cfg.DBName,
		ConnectionParams: cfg.DBConnParams,
	}

	dbConn, dbErr := db.InitialiseConnection()
	if dbErr != nil || dbConn == nil {
		panic(errors.New("could not initialise the DB"))
	}

	sessionManager, err := session.CreateSessionManager("hiuh")
	if err != nil {
		panic(err)
	}

	a := &app.App{
		Router:         bone.New(),
		TplParser:      &templates.TemplateParser{},
		SessionManager: sessionManager,
		Log:            logger.New("chatTest"),
	}

	a.InitRoute()
	a.Log.Error("gdh;lh")

	if err := http.ListenAndServe(cfg.GetPort(), a.Router); err != nil {
		panic(err)
	}
}
