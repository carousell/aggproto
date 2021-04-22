package complex_v1

import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/wallet"
	"github.com/carousell/aggproto/examples/goOut/media"
)

func transformComplexUsecaseRequest(req *ComplexUsecaseRequest) (*wallet.GetUserWalletRequest, *media.BulkGetMediaRequest, *listing.BulkGetListingsRequest) {
	getUserWalletRequest := &wallet.GetUserWalletRequest{}
	bulkGetMediaRequest := &media.BulkGetMediaRequest{}
	bulkGetListingsRequest := &listing.BulkGetListingsRequest{}
	getUserWalletRequest.UserId = req.UserInfo.Id
	bulkGetMediaRequest.UserId = req.UserInfo.Id
	bulkGetMediaRequest.MediaType = req.Meta.MediaType
	for i := 0; i < len(req.GetListings.ListingIds); i += 1 {
		bulkGetListingsRequest.ListingIds[i] = req.GetListings.ListingIds[i].Ids
	}
	return getUserWalletRequest, bulkGetMediaRequest, bulkGetListingsRequest
}
