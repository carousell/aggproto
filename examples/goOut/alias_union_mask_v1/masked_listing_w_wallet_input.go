package alias_union_mask_v1

import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/wallet"
)

func transformMaskedListingWWalletRequest(req *MaskedListingWWalletRequest) (*listing.GetListingRequest, *wallet.GetUserWalletRequest) {
	getListingRequest := &listing.GetListingRequest{}
	getListingRequest.ListingId = req.GetListingRequest.ListingId
	getUserWalletRequest := &wallet.GetUserWalletRequest{}
	getUserWalletRequest.UserId = req.GetUserWalletRequest.UserId
	return getListingRequest, getUserWalletRequest
}
