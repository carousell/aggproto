package alias_union_mask_v1

import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/listing_comments"
	"github.com/carousell/aggproto/examples/goOut/wallet"
)

func adaptMaskedListingWWalletResponse(getListingResponse *listing.GetListingResponse, getListingCommentsResponse *listing_comments.GetListingCommentsResponse, getUserWalletResponse *wallet.GetUserWalletResponse) *MaskedListingWWalletResponse{
	listing := getListingResponse.Listing
	userWallet := getUserWalletResponse.UserWallet
	resp := &MaskedListingWWalletResponse{}
	resp.Listing = &MaskedListingWWalletResponse_ListingGen{}
	resp.Listing.Title = listing.Title
	resp.Listing.Description = listing.Description
	resp.Wallet = &MaskedListingWWalletResponse_WalletGen{}
	resp.Wallet.Balance = userWallet.CoinBalance
	return resp
}
