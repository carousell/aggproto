package alias_union_mask_v1


import (
	"context"

	
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/listing_comments"
	"github.com/carousell/aggproto/examples/goOut/wallet"
)


type maskedListingWWalletSvc struct {
	UnimplementedMaskedListingWWalletServiceServer
	getListingClient *getListingClient
	getListingCommentsClient *getListingCommentsClient
	getUserWalletClient *getUserWalletClient
}


func New(listingComments listing_comments.ListingCommentsClient, wallet wallet.WalletClient, listings listing.ListingsClient) MaskedListingWWalletServiceServer {
	return &maskedListingWWalletSvc{
		getListingClient: &getListingClient{listings},
		getListingCommentsClient: &getListingCommentsClient{listingComments},
		getUserWalletClient: &getUserWalletClient{wallet},
	}
}


func (s *maskedListingWWalletSvc) InvokeMaskedListingWWallet(ctx context.Context, req *MaskedListingWWalletRequest) (*MaskedListingWWalletResponse, error){
	getListingRequest, getListingCommentsRequest, getUserWalletRequest := transformMaskedListingWWalletRequest(req)
	getListingResponse, err := s.getListingClient.getListing(ctx, getListingRequest)
	if err != nil {
		return nil, err
	}
	getListingCommentsResponse, err := s.getListingCommentsClient.getListingComments(ctx, getListingCommentsRequest)
	if err != nil {
		return nil, err
	}
	getUserWalletResponse, err := s.getUserWalletClient.getUserWallet(ctx, getUserWalletRequest)
	if err != nil {
		return nil, err
	}
	resp, err := adaptMaskedListingWWalletResponse(getListingResponse, getListingCommentsResponse, getUserWalletResponse)
	if err != nil {
		return nil, err
	}

	
	return resp, nil
}
