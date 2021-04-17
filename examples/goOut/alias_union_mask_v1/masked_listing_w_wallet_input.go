package alias_union_mask_v1

import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/listing_comments"
	"github.com/carousell/aggproto/examples/goOut/wallet"
)

func transformMaskedListingWWalletRequest(req *MaskedListingWWalletRequest) (*listing.GetListingRequest, *listing_comments.GetListingCommentsRequest, *wallet.GetUserWalletRequest) {
	getListingRequest := &listing.GetListingRequest{}
	getListingCommentsRequest := &listing_comments.GetListingCommentsRequest{}
	getListingRequest.ListingId = req.GetListing.Id
	getListingCommentsRequest.ListingId = req.GetListing.Id
	getUserWalletRequest := &wallet.GetUserWalletRequest{}
	getUserWalletRequest.UserId = req.GetWallet.Id
	return getListingRequest, getListingCommentsRequest, getUserWalletRequest
}
