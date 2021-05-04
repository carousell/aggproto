package union_mask_v1

import (
	"context"
	
	"github.com/carousell/aggproto/examples/goOut/wallet"
	"github.com/carousell/aggproto/examples/goOut/listing"
)

type maskedListingWWalletSvc struct {
	UnimplementedMaskedListingWWalletServiceServer
	getListingClient *getListingClient
	getUserWalletClient *getUserWalletClient
}

func New(wallet wallet.WalletClient, listings listing.ListingsClient) MaskedListingWWalletServiceServer {
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
