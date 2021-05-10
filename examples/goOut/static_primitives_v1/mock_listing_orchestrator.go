package static_primitives_v1

import (
	"context"
	
)

type mockListingSvc struct {
	UnimplementedMockListingServiceServer
}

func New() MockListingServiceServer {
	return &mockListingSvc{
	}
}

func (s *mockListingSvc) InvokeMockListing(ctx context.Context, req *MockListingRequest) (*MockListingResponse, error){
	resp, err := adaptMockListingResponse()
	if err != nil {
		return nil, err
	}
	
	return resp, nil
}
