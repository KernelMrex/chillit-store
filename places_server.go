package main

import (
	"chillit-store/places"
	"context"
)

type PlacesStoreServer struct{}

// TODO
func (s *PlacesStoreServer) GetPlaces(ctx context.Context, req *places.GetPlacesRequest) (*places.GetPlacesResponse, error) {
	return &places.GetPlacesResponse{}, nil
}

// TODO
func (s *PlacesStoreServer) AddPlace(ctx context.Context, req *places.AddPlaceRequest) (*places.AddPlaceResponse, error) {
	return &places.AddPlaceResponse{}, nil
}
