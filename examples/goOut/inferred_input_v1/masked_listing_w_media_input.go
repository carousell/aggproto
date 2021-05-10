package inferred_input_v1


import (
	"github.com/carousell/aggproto/examples/goOut/0"
)


func transformMaskedListingWMediaRequest(req *MaskedListingWMediaRequest) *listing.GetListingRequest {
	getListingRequest := &listing.GetListingRequest{}
	getListingRequest.ListingId = req.GetListing.Id
	return getListingRequest
}
