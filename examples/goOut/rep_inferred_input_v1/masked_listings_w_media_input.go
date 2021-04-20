package rep_inferred_input_v1

import (
	"github.com/carousell/aggproto/examples/goOut/listing"
)

func transformMaskedListingsWMediaRequest(req *MaskedListingsWMediaRequest) *listing.BulkGetListingsRequest {
	bulkGetListingsRequest := &listing.BulkGetListingsRequest{}
	bulkGetListingsRequest.ListingIds = req.GetListings.Ids
	return bulkGetListingsRequest
}
