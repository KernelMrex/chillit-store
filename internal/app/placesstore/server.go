package placesstore

import (
	"chillit-store/internal/app/models"
	"chillit-store/internal/app/places"
	"context"
	"errors"
	"time"

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

func (s *storeServer) GetPlaces(ctx context.Context, req *places.GetPlacesRequest) (*places.GetPlacesResponse, error) {
	timeoutCtx, _ := context.WithTimeout(ctx, 1*time.Second)
	dbPlaces, err := s.datastore.GetPlacesById(timeoutCtx, req.Offset, req.Amount)
	if err != nil {
		return &places.GetPlacesResponse{}, errors.New("error requesting data from database: " + err.Error())
	}

	pbPlaces := make([]*places.Place, len(dbPlaces))
	for i, place := range dbPlaces {
		pbPlaces[i] = &places.Place{
			Id:          place.Id,
			Title:       place.Title,
			Address:     place.Address,
			Description: place.Description,
		}
	}

	return &places.GetPlacesResponse{
		Places: pbPlaces,
	}, nil
}

func (s *storeServer) AddPlace(ctx context.Context, req *places.AddPlaceRequest) (*places.AddPlaceResponse, error) {
	timeoutCtx, _ := context.WithTimeout(ctx, 1*time.Second)
	insertedID, err := s.datastore.AddPlace(timeoutCtx, &models.Place{
		Id:          req.Place.Id,
		Title:       req.Place.Title,
		Address:     req.Place.Address,
		Description: req.Place.Description,
	})
	if err != nil {
		return &places.AddPlaceResponse{}, errors.New("error requesting data from database: " + err.Error())
	}
	return &places.AddPlaceResponse{
		Id: insertedID,
	}, nil
}
