package union_mask_v1


import (
	"github.com/carousell/aggproto/examples/goOut/wallet"
	"github.com/carousell/aggproto/examples/goOut/listing"
)


func adaptMaskedListingWWalletResponse(getListingResponse *listing.GetListingResponse, getUserWalletResponse *wallet.GetUserWalletResponse) (*MaskedListingWWalletResponse, error){
	resp := &MaskedListingWWalletResponse{}
	resp.Listing = &MaskedListingWWalletResponse_ListingGen{}
	resp.Listing.GetListingResponse = &MaskedListingWWalletResponse_ListingGen_GetListingResponseGen{}
	resp.Listing.GetListingResponse.Listing = &MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen{}
	resp.Listing.GetListingResponse.Listing.Title = getListingResponse.Listing.Title
	resp.Listing.GetListingResponse.Listing.Description = getListingResponse.Listing.Description
	resp.Wallet = &MaskedListingWWalletResponse_WalletGen{}
	resp.Wallet.GetUserWalletResponse = &MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen{}
	resp.Wallet.GetUserWalletResponse.UserWallet = &MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen{}
	resp.Wallet.GetUserWalletResponse.UserWallet.CoinBalance = getUserWalletResponse.UserWallet.CoinBalance
	return resp, nil
}
