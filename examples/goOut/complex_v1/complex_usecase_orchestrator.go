package complex_v1

import (
	"context"
	
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/media"
	"github.com/carousell/aggproto/examples/goOut/wallet"
)

type complexUsecaseSvc struct {
	UnimplementedComplexUsecaseServiceServer
	bulkGetListingsClient *bulkGetListingsClient
	bulkGetMediaClient *bulkGetMediaClient
	getUserWalletClient *getUserWalletClient
}

func New(listings listing.ListingsClient, mediaService media.MediaServiceClient, wallet wallet.WalletClient) ComplexUsecaseServiceServer {
	return &complexUsecaseSvc{
		bulkGetListingsClient: &bulkGetListingsClient{listings},
		bulkGetMediaClient: &bulkGetMediaClient{mediaService},
		getUserWalletClient: &getUserWalletClient{wallet},
	}
}

func (s *complexUsecaseSvc) InvokeComplexUsecase(ctx context.Context, req *ComplexUsecaseRequest) (*ComplexUsecaseResponse, error){
	getUserWalletRequest, bulkGetMediaRequest, bulkGetListingsRequest := transformComplexUsecaseRequest(req)
	bulkGetListingsResponse, err := s.bulkGetListingsClient.bulkGetListings(ctx, bulkGetListingsRequest)
	if err != nil {
		return nil, err
	}
	transformToBulkGetMediaRequest(bulkGetMediaRequest, bulkGetListingsResponse)
	bulkGetMediaResponse, err := s.bulkGetMediaClient.bulkGetMedia(ctx, bulkGetMediaRequest)
	if err != nil {
		return nil, err
	}
	getUserWalletResponse, err := s.getUserWalletClient.getUserWallet(ctx, getUserWalletRequest)
	if err != nil {
		return nil, err
	}
	resp, err := adaptComplexUsecaseResponse(bulkGetListingsResponse, bulkGetMediaResponse, getUserWalletResponse)
	if err != nil {
		return nil, err
	}
	
	return resp, nil
}
