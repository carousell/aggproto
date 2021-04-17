package inferred_input_v1

import (
	"context"
	
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/media"
)

type maskedListingWMediaSvc struct {
	UnimplementedMaskedListingWMediaServiceServer
	getListingClient *getListingClient
	getMediaClient *getMediaClient
}

func New(mediaService media.MediaServiceClient, listings listing.ListingsClient) MaskedListingWMediaServiceServer {
	return &maskedListingWMediaSvc{
		getListingClient: &getListingClient{listings},
		getMediaClient: &getMediaClient{mediaService},
	}
}

func (s *maskedListingWMediaSvc) InvokeMaskedListingWMedia(ctx context.Context, req *MaskedListingWMediaRequest) (*MaskedListingWMediaResponse, error){
	getListingRequest := transformMaskedListingWMediaRequest(req)
	getListingResponse, err := s.getListingClient.getListing(ctx, getListingRequest)
	if err != nil {
		return nil, err
	}
	getMediaRequest := transformToGetMediaRequest(getListingResponse)
	getMediaResponse, err := s.getMediaClient.getMedia(ctx, getMediaRequest)
	if err != nil {
		return nil, err
	}
	resp := adaptMaskedListingWMediaResponse(getListingResponse, getMediaResponse)
	return resp, nil
}
