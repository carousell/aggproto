package rep_inferred_input_v1

import (
	"context"
	
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/media"
)

type maskedListingsWMediaSvc struct {
	UnimplementedMaskedListingsWMediaServiceServer
	bulkGetListingsClient *bulkGetListingsClient
	bulkGetMediaClient *bulkGetMediaClient
}

func New(listings listing.ListingsClient, mediaService media.MediaServiceClient) MaskedListingsWMediaServiceServer {
	return &maskedListingsWMediaSvc{
		bulkGetListingsClient: &bulkGetListingsClient{listings},
		bulkGetMediaClient: &bulkGetMediaClient{mediaService},
	}
}

func (s *maskedListingsWMediaSvc) InvokeMaskedListingsWMedia(ctx context.Context, req *MaskedListingsWMediaRequest) (*MaskedListingsWMediaResponse, error){
	bulkGetListingsRequest, bulkGetMediaRequest := transformMaskedListingsWMediaRequest(req)
	bulkGetListingsResponse, err := s.bulkGetListingsClient.bulkGetListings(ctx, bulkGetListingsRequest)
	if err != nil {
		return nil, err
	}
	transformToBulkGetMediaRequest(bulkGetMediaRequest, bulkGetListingsResponse)
	bulkGetMediaResponse, err := s.bulkGetMediaClient.bulkGetMedia(ctx, bulkGetMediaRequest)
	if err != nil {
		return nil, err
	}
	resp, err := adaptMaskedListingsWMediaResponse(bulkGetListingsResponse, bulkGetMediaResponse)
	if err != nil {
		return nil, err
	}
	
	return resp, nil
}
