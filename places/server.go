package places

import (
	"chillit-store/environment"
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

// TODO
func (s *StoreServer) AddPlace(ctx context.Context, req *AddPlaceRequest) (*AddPlaceResponse, error) {
	return &AddPlaceResponse{}, nil
}
