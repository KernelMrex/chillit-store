package environment

import (
	"chillit-store/internal/app/configuration"
	"chillit-store/internal/app/models"
	"errors"
	"log"
	"os"
)

// Env provides easy-in-use dependencies structure(deprecated)
type Env struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DB          models.Datastore
}

// BuildEnv auto build Env
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
