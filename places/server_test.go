package places

import (
	"chillit-store/environment"
	"chillit-store/models"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

const mockServerPort = 8888

var env *environment.Env

var mockServer *grpc.Server

var mockClient PlacesStoreClient

func init() {
	env = &environment.Env{
		InfoLogger:  nil,
		ErrorLogger: nil,
		DB:          &models.MockDB{},
	}

	// Setting up grpc server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", mockServerPort))
	if err != nil {
		log.Fatalln("[ test init ] could not start listen:", err)
	}
	mockServer = grpc.NewServer()
	RegisterPlacesStoreServer(mockServer, &StoreServer{
		Env: env,
	})
	go func() {
		if err := mockServer.Serve(listener); err != nil {
			log.Fatalln("[ test init ] error while serving 'PlacesStoreServer':", err)
		}
	}()

	// Setting up grpc client
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", mockServerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalln("[ test init ] error connecting to mock server:", err)
	}
	mockClient = NewPlacesStoreClient(conn)
}

func TestAddPlaceRequest(t *testing.T) {
	for i := 0; i < 6; i++ {
		if resp, err := mockClient.AddPlace(context.Background(), &AddPlaceRequest{
			Place: &Place{
				Title: fmt.Sprintf("Test title %d", i),
				Address: fmt.Sprintf("Test address %d", i),
				Description: fmt.Sprintf("Test description %d", i),
			},
		}); err != nil || resp.Id != uint64(i) {
			t.Fail()
			return
		}
	}
}

func TestGetPlacesRequest(t *testing.T) {
	const offset = 1
	resp, err := mockClient.GetPlaces(context.Background(), &GetPlacesRequest{
		Offset:           offset,
		Amount:           4,
		ShortDescription: false,
	})
	if err != nil {
		t.Fail()
		return
	}
	for i, place := range resp.Places {
		if place.Id != uint64(i + offset) ||
			place.Title != fmt.Sprintf("Test title %d", i + offset) ||
			place.Address != fmt.Sprintf("Test address %d", i + offset) ||
			place.Description != fmt.Sprintf("Test description %d", i + offset) {
			t.Fail()
			return
		}
	}
}
