package union_mask_v1

import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/wallet"
)

func transformMaskedListingWWalletRequest(req *MaskedListingWWalletRequest) (*listing.GetListingRequest, *wallet.GetUserWalletRequest) {
	getListingRequest := &listing.GetListingRequest{}
	getUserWalletRequest := &wallet.GetUserWalletRequest{}
	getListingRequest.ListingId = req.GetListingRequest.ListingId
	getUserWalletRequest.UserId = req.GetUserWalletRequest.UserId
	return getListingRequest, getUserWalletRequest
}
