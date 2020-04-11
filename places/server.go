package places

import (
	"chillit-store/environment"
	"chillit-store/models"
	"context"
	"errors"
	"time"
)

type StoreServer struct {
	Env *environment.Env
}

func (s *StoreServer) GetPlaces(ctx context.Context, req *GetPlacesRequest) (*GetPlacesResponse, error) {
	timeoutCtx, _ := context.WithTimeout(ctx, 1*time.Second)
	dbPlaces, err := s.Env.DB.GetPlacesById(timeoutCtx, req.Offset, req.Amount)
	if err != nil {
		return &GetPlacesResponse{}, errors.New("error requesting data from database: " + err.Error())
	}

	pbPlaces := make([]*Place, len(dbPlaces))
	for i, place := range dbPlaces {
		pbPlaces[i] = &Place{
			Id:          place.Id,
			Title:       place.Title,
			Address:     place.Address,
			Description: place.Description,
		}
	}

	return &GetPlacesResponse{
		Places: pbPlaces,
	}, nil
}

func (s *StoreServer) AddPlace(ctx context.Context, req *AddPlaceRequest) (*AddPlaceResponse, error) {
	timeoutCtx, _ := context.WithTimeout(ctx, 1*time.Second)
	insertedId, err := s.Env.DB.AddPlace(timeoutCtx, &models.Place{
		Id:          req.Place.Id,
		Title:       req.Place.Title,
		Address:     req.Place.Address,
		Description: req.Place.Description,
	})
	if err != nil {
		return &AddPlaceResponse{}, errors.New("error requesting data from database: " + err.Error())
	}

	return &AddPlaceResponse{
		Id: insertedId,
	}, nil
}
