package main

import (
	"chillit-store/internal/app/configuration"
	"chillit-store/internal/app/environment"
	"chillit-store/internal/app/placesstore"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var Env *environment.Env

func init() {
	// Logger initialization
	initInfoLogger := log.New(os.Stdout, "Init info: ", 0)
	initErrorLogger := log.New(os.Stderr, "Init error: ", 0)
	initInfoLogger.Println("Initialization started...")

	// Getting config path from flag
	var confPath string
	flag.StringVar(&confPath, "config_path", "configs/config.yaml", "path for '.yaml' configuration file")
	flag.Parse()

	// Build config and env
	conf, err := configuration.NewConfig(confPath)
	if err != nil {
		initErrorLogger.Fatalln(err)
	}
	initInfoLogger.Println("Configuration has loaded")

	Env, err = environment.BuildEnv(conf)
	if err != nil {
		initErrorLogger.Fatalln(err)
	}
	initInfoLogger.Println("Environment has built")
	initInfoLogger.Println("Initialization successful")
}

func main() {
	Env.InfoLogger.Println("[ main ]", "Starting store service is now running!")

	go func() {
		err := placesstore.Start(&placesstore.Config{
			Hostname: ":10100",
		}, Env.DB)
		if err != nil {
			// TODO: send error to workflow chan
			Env.ErrorLogger.Fatalln("[ main ]", err)
		}
	}()
	Env.InfoLogger.Println("[ main ]", "Store service is now running!")

	WaitForTermSig()
	Env.InfoLogger.Println("[ main ]", "Store service is now shutting down...")

	Env.InfoLogger.Println("[ main ]", "Store service stopped")
}

// WaitForTermSig waits for termination signal from os
func WaitForTermSig() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
