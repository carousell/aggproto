package union_mask_v1

import (
	"github.com/carousell/aggproto/examples/goOut/wallet"
	"github.com/carousell/aggproto/examples/goOut/listing"
)

func transformMaskedListingWWalletRequest(req *MaskedListingWWalletRequest) (*wallet.GetUserWalletRequest, *listing.GetListingRequest) {
	getUserWalletRequest := &wallet.GetUserWalletRequest{}
	getUserWalletRequest.UserId = req.GetUserWalletRequest.UserId
	getListingRequest := &listing.GetListingRequest{}
	getListingRequest.ListingId = req.GetListingRequest.ListingId
	return getUserWalletRequest, getListingRequest
}
