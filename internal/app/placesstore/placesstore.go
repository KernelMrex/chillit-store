package placesstore

import (
	"chillit-store/internal/app/models"
	"chillit-store/internal/app/places"
	"net"

	"google.golang.org/grpc"
)

// Start starts StorePlaces gRPC server
func Start(config *Config, datastore models.Datastore) error {
	listener, err := net.Listen("tcp", config.Hostname)
	if err != nil {
		return err
	}
	defer listener.Close()
	storeServer := grpc.NewServer()
	places.RegisterPlacesStoreServer(storeServer, newServer(datastore))
	return storeServer.Serve(listener)
}
