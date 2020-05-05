package placesstore

import (
	"chillit-store/internal/app/models"
	"chillit-store/internal/app/places"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type storeServer struct {
	datastore models.Datastore
	logger    *logrus.Logger
}

func newServer(datastore models.Datastore) *storeServer {
	return &storeServer{
		datastore: datastore,
		logger:    logrus.New(),
	}
}

func (s *storeServer) AddPlace(ctx context.Context, req *places.AddPlaceRequest) (*places.AddPlaceResponse, error) {
	return &places.AddPlaceResponse{}, fmt.Errorf("not implemented yet")
}

func (s *storeServer) GetRandomPlaceByCityName(ctx context.Context, req *places.GetRandomPlaceByCityNameRequest) (*places.GetRandomPlaceByCityNameResponse, error) {
	dbPlaceModel, err := s.datastore.GetRandomPlaceByCityName(ctx, req.CityName)
	if err != nil {
		s.logger.Errorf("could not request datastore for city '%s' error: %v", req.CityName, err)
		return &places.GetRandomPlaceByCityNameResponse{}, fmt.Errorf("error while requesting datastore")
	}

	return &places.GetRandomPlaceByCityNameResponse{
		Place: &places.Place{
			Id:          dbPlaceModel.ID,
			Title:       dbPlaceModel.Title,
			Description: dbPlaceModel.Description,
			Address:     dbPlaceModel.Address,
		},
	}, nil
}

func (s *storeServer) GetCities(ctx context.Context, req *places.GetCitiesRequest) (*places.GetCitiesResponse, error) {
	return &places.GetCitiesResponse{}, fmt.Errorf("not implemented yet")
}
