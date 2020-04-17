package placesstore

import (
	"chillit-store/internal/app/models"
	"chillit-store/internal/app/places"
	"net"

	"google.golang.org/grpc"
)

var storeGRPCServer *grpc.Server

// Start starts StorePlaces gRPC server
func Start(config *Config, datastore models.Datastore) error {
	listener, err := net.Listen("tcp", config.Hostname)
	if err != nil {
		return err
	}
	defer listener.Close()
	storeGRPCServer = grpc.NewServer()
	places.RegisterPlacesStoreServer(storeGRPCServer, newServer(datastore))
	return storeGRPCServer.Serve(listener)
}

// Stop stops StorePlaces gRPC server
func Stop() {
	if storeGRPCServer != nil {
		storeGRPCServer.Stop()
	}
}
