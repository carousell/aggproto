package alias_union_mask_v1


import (
	"github.com/carousell/aggproto/examples/goOut/0"
	"github.com/carousell/aggproto/examples/goOut/1"
	"github.com/carousell/aggproto/examples/goOut/2"
)


func transformMaskedListingWWalletRequest(req *MaskedListingWWalletRequest) (*listing.GetListingRequest, *listing_comments.GetListingCommentsRequest, *wallet.GetUserWalletRequest) {
	getListingRequest := &listing.GetListingRequest{}
	getListingCommentsRequest := &listing_comments.GetListingCommentsRequest{}
	getUserWalletRequest := &wallet.GetUserWalletRequest{}
	getListingRequest.ListingId = req.GetListing.Id
	getListingCommentsRequest.ListingId = req.GetListing.Id
	getUserWalletRequest.UserId = req.GetWallet.Id
	return getListingRequest, getListingCommentsRequest, getUserWalletRequest
}
