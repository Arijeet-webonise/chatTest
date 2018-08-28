package main

import (
	"errors"
	"net/http"

	"github.com/Arijeet-webonise/chatTest/app"
	"github.com/Arijeet-webonise/chatTest/app/config"
	"github.com/Arijeet-webonise/chatTest/app/models"
	"github.com/Arijeet-webonise/chatTest/pkg/database"
	"github.com/Arijeet-webonise/chatTest/pkg/logger"
	"github.com/Arijeet-webonise/chatTest/pkg/redis"
	"github.com/Arijeet-webonise/chatTest/pkg/session"
	"github.com/Arijeet-webonise/chatTest/pkg/storage"
	"github.com/Arijeet-webonise/chatTest/pkg/templates"
	"github.com/go-zoo/bone"
	"github.com/gorilla/csrf"
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

	sessionManager, err := session.CreateSessionManager(cfg.SessionSecretKey)
	if err != nil {
		panic(err)
	}

	log := &logger.RealLogger{}
	log.Initialise()
	redisService, pong, err := redis.InitRedis("localhost:6379", "", 0)
	if err != nil {
		panic(err)
	}
	log.Info(pong)
	a := &app.App{
		Router:         bone.New(),
		TplParser:      &templates.TemplateParser{},
		SessionManager: sessionManager,
		Log:            log,
		StorageManager: &storage.StorageManagerServiceImpl{
			AccessKey: cfg.AwsAccessKeyID,
			SecretKey: cfg.AwsAccessSecretKey,
			Bucket:    cfg.AwsBucket,
			Endpoint:  cfg.AwsEndPointURL,
			Region:    cfg.AwsRegion,
		},
		RedisService:            redisService,
		UserService:             &models.PortalUserServiceImpl{DB: dbConn},
		CustomPortalUserService: &models.CustomPortalUserImpl{DB: dbConn},
	}
	CSRF := csrf.Protect([]byte(cfg.CSRFSecretKey), csrf.Secure(cfg.CSRFSecure))
	a.InitRoute()

	if err := http.ListenAndServe(cfg.GetPort(), CSRF(a.Router)); err != nil {
		panic(err)
	}
}
