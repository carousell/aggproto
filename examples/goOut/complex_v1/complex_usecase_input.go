package complex_v1

import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/listing_comments"
	"github.com/carousell/aggproto/examples/goOut/wallet"
	"github.com/carousell/aggproto/examples/goOut/media"
)

func transformComplexUsecaseRequest(req *ComplexUsecaseRequest) (*media.BulkGetMediaRequest, *listing.BulkGetListingsRequest, *listing_comments.BulkGetListingCommentsRequest, *wallet.GetUserWalletRequest) {
	bulkGetMediaRequest := &media.BulkGetMediaRequest{}
	bulkGetListingsRequest := &listing.BulkGetListingsRequest{}
	bulkGetListingCommentsRequest := &listing_comments.BulkGetListingCommentsRequest{}
	getUserWalletRequest := &wallet.GetUserWalletRequest{}
	bulkGetMediaRequest.MediaType = req.Meta.MediaType
	for i := 0; i < len(req.GetListings.ListingIds); i += 1 {
		bulkGetListingsRequest.ListingIds[i] = req.GetListings.ListingIds[i].Ids
		bulkGetListingCommentsRequest.ListingId[i] = req.GetListings.ListingIds[i].Ids
	}
	getUserWalletRequest.UserId = req.UserInfo.Id
	bulkGetMediaRequest.UserId = req.UserInfo.Id
	return bulkGetMediaRequest, bulkGetListingsRequest, bulkGetListingCommentsRequest, getUserWalletRequest
}
