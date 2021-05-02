package rep_inferred_input_v1

import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/media"
)

func transformMaskedListingsWMediaRequest(req *MaskedListingsWMediaRequest) (*listing.BulkGetListingsRequest, *media.BulkGetMediaRequest) {
	bulkGetListingsRequest := &listing.BulkGetListingsRequest{}
	bulkGetMediaRequest := &media.BulkGetMediaRequest{}
	bulkGetListingsRequest.ListingIds = req.GetListings.Ids
	bulkGetMediaRequest.UserId = req.BulkGetMediaRequest.UserId
	bulkGetMediaRequest.MediaType = req.BulkGetMediaRequest.MediaType
	return bulkGetListingsRequest, bulkGetMediaRequest
}
