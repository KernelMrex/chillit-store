package places

import (
	"context"

	grpc "google.golang.org/grpc"
)

// MockPlacesStoreServer mock store places server
type MockPlacesStoreServer struct {
	GetPlacesResponse *GetPlacesResponse
	GetPlacesError    error

	AddPlaceResponse *AddPlaceResponse
	AddPlaceError    error
}

// GetPlaces mock method
func (m *MockPlacesStoreServer) GetPlaces(ctx context.Context, req *GetPlacesRequest) (*GetPlacesResponse, error) {
	return m.GetPlacesResponse, m.GetPlacesError
}

// AddPlace mock method
func (m *MockPlacesStoreServer) AddPlace(ctx context.Context, req *AddPlaceRequest) (*AddPlaceResponse, error) {
	return m.AddPlaceResponse, m.AddPlaceError
}

// MockPlacesStoreClient mock store places server
type MockPlacesStoreClient struct {
	GetPlacesResponse *GetPlacesResponse
	GetPlacesError    error

	AddPlaceResponse *AddPlaceResponse
	AddPlaceError    error
}

// GetPlaces mock method
func (m *MockPlacesStoreClient) GetPlaces(ctx context.Context, req *GetPlacesRequest, opts ...grpc.CallOption) (*GetPlacesResponse, error) {
	return m.GetPlacesResponse, m.GetPlacesError
}

// AddPlace mock method
func (m *MockPlacesStoreClient) AddPlace(ctx context.Context, req *AddPlaceRequest, opts ...grpc.CallOption) (*AddPlaceResponse, error) {
	return m.AddPlaceResponse, m.AddPlaceError
}
