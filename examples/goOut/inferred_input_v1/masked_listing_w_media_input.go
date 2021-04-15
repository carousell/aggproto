package inferred_input_v1

import (
	"github.com/carousell/aggproto/examples/goOut/listing"
)

func transformMaskedListingWMediaRequest(req *MaskedListingWMediaRequest) *listing.GetListingRequest {
	getListingRequest := &listing.GetListingRequest{}
	getListingRequest.ListingId = req.GetListingRequest.ListingId
	return getListingRequest
}
