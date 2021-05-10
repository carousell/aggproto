package simple_mask_v1


import (
	"github.com/carousell/aggproto/examples/goOut/listing"
)


func transformMaskedListingRequest(req *MaskedListingRequest) *listing.GetListingRequest {
	getListingRequest := &listing.GetListingRequest{}
	getListingRequest.ListingId = req.GetListingRequest.ListingId
	return getListingRequest
}
