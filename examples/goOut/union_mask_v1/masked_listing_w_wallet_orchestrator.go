package union_mask_v1


import (
	"context"

	
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/wallet"
)


type maskedListingWWalletSvc struct {
	UnimplementedMaskedListingWWalletServiceServer
	getListingClient *getListingClient
	getUserWalletClient *getUserWalletClient
}


func New(listings listing.ListingsClient, wallet wallet.WalletClient) MaskedListingWWalletServiceServer {
	return &maskedListingWWalletSvc{
		getListingClient: &getListingClient{listings},
		getUserWalletClient: &getUserWalletClient{wallet},
	}
}


func (s *maskedListingWWalletSvc) InvokeMaskedListingWWallet(ctx context.Context, req *MaskedListingWWalletRequest) (*MaskedListingWWalletResponse, error){
	getListingRequest, getUserWalletRequest := transformMaskedListingWWalletRequest(req)
	getListingResponse, err := s.getListingClient.getListing(ctx, getListingRequest)
	if err != nil {
		return nil, err
	}
	getUserWalletResponse, err := s.getUserWalletClient.getUserWallet(ctx, getUserWalletRequest)
	if err != nil {
		return nil, err
	}
	resp, err := adaptMaskedListingWWalletResponse(getListingResponse, getUserWalletResponse)
	if err != nil {
		return nil, err
	}

	
	return resp, nil
}
