package environment

import (
	"chillit-store/configuration"
	"chillit-store/models"
	"errors"
	"log"
	"os"
)

type Env struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DB          models.Datastore
}

func BuildEnv(conf *configuration.Configuration) (*Env, error) {
	dbConn, err := models.NewMysqlDB(conf.DB.GetUrl())
	if err != nil {
		return nil, errors.New("[ BuildEnv ]" + err.Error())
	}

	return &Env{
		InfoLogger:  log.New(os.Stdout, "INFO: ", 0),
		ErrorLogger: log.New(os.Stderr, "ERROR:", 0),
		DB:          dbConn,
	}, nil
}
