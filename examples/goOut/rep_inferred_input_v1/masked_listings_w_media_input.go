package rep_inferred_input_v1

import (
	"github.com/carousell/aggproto/examples/goOut/media"
	"github.com/carousell/aggproto/examples/goOut/listing"
)

func transformMaskedListingsWMediaRequest(req *MaskedListingsWMediaRequest) (*media.BulkGetMediaRequest, *listing.BulkGetListingsRequest) {
	bulkGetMediaRequest := &media.BulkGetMediaRequest{}
	bulkGetListingsRequest := &listing.BulkGetListingsRequest{}
	bulkGetMediaRequest.UserId = req.BulkGetMediaRequest.UserId
	bulkGetMediaRequest.MediaType = req.BulkGetMediaRequest.MediaType
	bulkGetListingsRequest.ListingIds = req.GetListings.Ids
	return bulkGetMediaRequest, bulkGetListingsRequest
}
