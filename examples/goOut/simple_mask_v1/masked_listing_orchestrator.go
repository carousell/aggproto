package simple_mask_v1


import (
	"context"

	
	"github.com/carousell/aggproto/examples/goOut/listing"
)


type maskedListingSvc struct {
	UnimplementedMaskedListingServiceServer
	getListingClient *getListingClient
}


func New(listings listing.ListingsClient) MaskedListingServiceServer {
	return &maskedListingSvc{
		getListingClient: &getListingClient{listings},
	}
}


func (s *maskedListingSvc) InvokeMaskedListing(ctx context.Context, req *MaskedListingRequest) (*MaskedListingResponse, error){
	getListingRequest := transformMaskedListingRequest(req)
	getListingResponse, err := s.getListingClient.getListing(ctx, getListingRequest)
	if err != nil {
		return nil, err
	}
	resp, err := adaptMaskedListingResponse(getListingResponse)
	if err != nil {
		return nil, err
	}

	
	return resp, nil
}
