package main

import (
	"chillit-store/configuration"
	"chillit-store/environment"
	"chillit-store/places"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
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
	flag.StringVar(&confPath, "config_path", "config.yaml", "path for '.yaml' configuration file")
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
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		Env.ErrorLogger.Fatalln("[ main ] could not start listen:", err)
	}
	storeServer := grpc.NewServer()
	places.RegisterPlacesStoreServer(storeServer, &places.StoreServer{
		Env: Env,
	})
	go func() {
		if err := storeServer.Serve(listener); err != nil {
			Env.ErrorLogger.Fatalln("[ main ] error while serving 'PlacesStoreServer':", err)
		}
	}()
	// Wait here until CTRL-C or other term signal is received.
	Env.InfoLogger.Println("[ main ]", "Store service is now running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	Env.InfoLogger.Println("[ main ]", "Store service is now shutting down...")
}
