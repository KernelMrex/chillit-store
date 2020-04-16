package main

import (
	"chillit-store/internal/app/configuration"
	"chillit-store/internal/app/models"
	"chillit-store/internal/app/placesstore"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Conf application configuration
var Conf *configuration.Configuration

var configPath string

func init() {
	flag.StringVar(&configPath, "config_path", "configs/config.yaml", "path for '.yaml' configuration file")
}

func main() {
	flag.Parse()

	log.Println("[ main ]", "Starting store service")

	// Parsing config
	conf, err := configuration.NewConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	// Connecting to datastore
	datastore, err := models.NewMysqlDB(conf.DB.GetUrl())
	if err != nil {
		log.Fatalln(err)
	}

	// Starting PlacesStore service
	go StartPlacesStore(conf.PlacesStore, datastore)
	log.Println("[ main ]", "Store service is now running!")

	WaitForTermSig()
	log.Println("[ main ]", "Store service stopped")
}

// StartPlacesStore starts places store server in separate go-routine
func StartPlacesStore(conf *placesstore.Config, datastore models.Datastore) {
	err := placesstore.Start(&placesstore.Config{
		Hostname: Conf.PlacesStore.Hostname,
	}, datastore)
	if err != nil {
		// TODO: send error to workflow chan
		log.Fatalln("[ main ]", err)
	}
}

// WaitForTermSig waits for termination signal from os
func WaitForTermSig() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
