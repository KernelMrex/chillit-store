package placesstore

import (
	"chillit-store/internal/app/models"
	"chillit-store/internal/app/places"
	"context"
	"fmt"
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

func (s *storeServer) GetRandomPlaceByCityName(ctx context.Context, req *places.GetRandomPlaceByCityNameRequest) (*places.GetRandomPlaceByCityNameResponse, error) {
	timeoutContext, _ := context.WithTimeout(ctx, time.Second*1)
	dbPlaceModel, err := s.datastore.GetRandomPlaceByCityName(timeoutContext, req.GetCityName())
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
	timeoutContext, _ := context.WithTimeout(ctx, time.Second*1)
	dbCityModels, err := s.datastore.GetCities(timeoutContext, req.GetAmount(), req.GetOffset())
	if err != nil {
		s.logger.Errorf("could not request datastore for limit '%d' and offset '%d' error: %v", req.GetAmount(), req.GetOffset(), err)
		return &places.GetCitiesResponse{}, fmt.Errorf("error while requesting datastore")
	}

	cities := make([]*places.City, len(dbCityModels))
	for i, city := range dbCityModels {
		cities[i] = &places.City{
			Id:    city.ID,
			Title: city.Title,
		}
	}

	return &places.GetCitiesResponse{
		Cities: cities,
	}, nil
}

func (s *storeServer) AddPlace(ctx context.Context, req *places.AddPlaceRequest) (*places.AddPlaceResponse, error) {
	timeoutContext, _ := context.WithTimeout(ctx, time.Second*1)
	id, err := s.datastore.SavePlace(timeoutContext, &models.Place{
		Title:       req.GetPlace().GetTitle(),
		Address:     req.GetPlace().GetAddress(),
		Description: req.GetPlace().GetDescription(),
		ImageURL:    req.GetPlace().GetImgURL(),
	}, req.CityName)
	if err != nil {
		s.logger.Errorf("could not save place '%s' error while requesting datastore: %v", req.Place.Title, err)
		return &places.AddPlaceResponse{}, fmt.Errorf("error while requesting datastore")
	}

	return &places.AddPlaceResponse{Id: id}, nil
}

func (s *storeServer) GetPlacesByCityID(ctx context.Context, req *places.GetPlacesByCityIDRequest) (*places.GetPlacesByCityIDResponse, error) {
	timeoutContext, _ := context.WithTimeout(ctx, time.Second*1)
	placesModels, err := s.datastore.GetPlacesByCityID(timeoutContext, req.GetCityID(), req.Amount, req.Offset)
	if err != nil {
		s.logger.Errorf("could not get placesModels from datastore error: %v", err)
		return &places.GetPlacesByCityIDResponse{}, fmt.Errorf("error while requesting datastore")
	}

	grpcPlaces := make([]*places.Place, len(placesModels))
	for i, placeModel := range placesModels {
		grpcPlaces[i] = &places.Place{
			Id:          placeModel.ID,
			Title:       placeModel.Title,
			Address:     placeModel.Address,
			Description: placeModel.Description,
			ImgURL:      placeModel.ImageURL,
		}
	}

	return &places.GetPlacesByCityIDResponse{
		Places: grpcPlaces,
	}, nil
}
