package main

import (
	"chillit-store/configuration"
	"chillit-store/environment"
	"log"
	"os"
)

var Env *environment.Env

func init() {
	initInfoLogger := log.New(os.Stdout, "Init info  | ", 0)
	initErrorLogger := log.New(os.Stderr, "Init error | ", 0)
	initInfoLogger.Println("Initialization started...")

	// Build config and env
	conf, err := configuration.NewConfig("config.yml")
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

}
