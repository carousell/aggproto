package complex_v1


import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/listing_comments"
	"github.com/carousell/aggproto/examples/goOut/media"
	"github.com/carousell/aggproto/examples/goOut/wallet"
)


func transformComplexUsecaseRequest(req *ComplexUsecaseRequest) (*wallet.GetUserWalletRequest, *media.BulkGetMediaRequest, *listing.BulkGetListingsRequest, *listing_comments.BulkGetListingCommentsRequest) {
	getUserWalletRequest := &wallet.GetUserWalletRequest{}
	bulkGetMediaRequest := &media.BulkGetMediaRequest{}
	bulkGetListingsRequest := &listing.BulkGetListingsRequest{}
	bulkGetListingCommentsRequest := &listing_comments.BulkGetListingCommentsRequest{}
	getUserWalletRequest.UserId = req.UserInfo.Id
	bulkGetMediaRequest.UserId = req.UserInfo.Id
	bulkGetMediaRequest.MediaType = req.Meta.MediaType
	for i := 0; i < len(req.GetListings.ListingIds); i += 1 {
		bulkGetListingsRequest.ListingIds[i] = req.GetListings.ListingIds[i].Ids
		bulkGetListingCommentsRequest.ListingId[i] = req.GetListings.ListingIds[i].Ids
	}
	return getUserWalletRequest, bulkGetMediaRequest, bulkGetListingsRequest, bulkGetListingCommentsRequest
}
