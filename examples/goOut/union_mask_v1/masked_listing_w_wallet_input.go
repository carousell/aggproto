package union_mask_v1


import (
	"github.com/carousell/aggproto/examples/goOut/0"
	"github.com/carousell/aggproto/examples/goOut/1"
)


func transformMaskedListingWWalletRequest(req *MaskedListingWWalletRequest) (*wallet.GetUserWalletRequest, *listing.GetListingRequest) {
	getUserWalletRequest := &wallet.GetUserWalletRequest{}
	getListingRequest := &listing.GetListingRequest{}
	getUserWalletRequest.UserId = req.GetUserWalletRequest.UserId
	getListingRequest.ListingId = req.GetListingRequest.ListingId
	return getUserWalletRequest, getListingRequest
}
