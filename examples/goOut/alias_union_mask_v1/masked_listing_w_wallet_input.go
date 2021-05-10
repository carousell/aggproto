package alias_union_mask_v1


import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/listing_comments"
	"github.com/carousell/aggproto/examples/goOut/wallet"
)


func transformMaskedListingWWalletRequest(req *MaskedListingWWalletRequest) (*wallet.GetUserWalletRequest, *listing.GetListingRequest, *listing_comments.GetListingCommentsRequest) {
	getUserWalletRequest := &wallet.GetUserWalletRequest{}
	getListingRequest := &listing.GetListingRequest{}
	getListingCommentsRequest := &listing_comments.GetListingCommentsRequest{}
	getUserWalletRequest.UserId = req.GetWallet.Id
	getListingRequest.ListingId = req.GetListing.Id
	getListingCommentsRequest.ListingId = req.GetListing.Id
	return getUserWalletRequest, getListingRequest, getListingCommentsRequest
}
