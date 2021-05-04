package complex_v1

import (
	"github.com/pkg/errors"
	
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/media"
	"github.com/carousell/aggproto/examples/goOut/listing_comments"
	"github.com/carousell/aggproto/examples/goOut/wallet"
)

func adaptComplexUsecaseResponse(bulkGetListingsResponse *listing.BulkGetListingsResponse, bulkGetMediaResponse *media.BulkGetMediaResponse, bulkGetListingCommentsResponse *listing_comments.BulkGetListingCommentsResponse, getUserWalletResponse *wallet.GetUserWalletResponse) (*ComplexUsecaseResponse, error){
	resp := &ComplexUsecaseResponse{}
	if len(bulkGetListingsResponse.Listings) != len(bulkGetMediaResponse.Medias) {
		return nil, errors.Errorf("assertion failed %s != %s", len(bulkGetListingsResponse.Listings), len(bulkGetMediaResponse.Medias))
	}
	if len(bulkGetListingsResponse.Listings) != len(bulkGetListingCommentsResponse.ListingComments) {
		return nil, errors.Errorf("assertion failed %s != %s", len(bulkGetListingsResponse.Listings), len(bulkGetListingCommentsResponse.ListingComments))
	}
	resp.Listings = make([]*ComplexUsecaseResponse_ListingsGen, len(bulkGetListingsResponse.Listings))
	for i := 0; i < len(bulkGetListingsResponse.Listings); i++ {
		resp.Listings[i] = &ComplexUsecaseResponse_ListingsGen{}
		resp.Listings[i].Info = &ComplexUsecaseResponse_ListingsGen_InfoGen{}
		resp.Listings[i].Info.Title = bulkGetListingsResponse.Listings[i].Title
		resp.Listings[i].Info.Description = bulkGetListingsResponse.Listings[i].Description
		resp.Listings[i].Media = &ComplexUsecaseResponse_ListingsGen_MediaGen{}
		resp.Listings[i].Media.PhotoUrl = bulkGetMediaResponse.Medias[i].PhotoUrl
		resp.Listings[i].Comments = make([]*ComplexUsecaseResponse_ListingsGen_CommentsGen, len(bulkGetListingCommentsResponse.ListingComments[i].Comments))
		for j := 0; j < len(bulkGetListingCommentsResponse.ListingComments[i].Comments); j++ {
			resp.Listings[i].Comments[j] = &ComplexUsecaseResponse_ListingsGen_CommentsGen{}
			resp.Listings[i].Comments[j].Info = &ComplexUsecaseResponse_ListingsGen_CommentsGen_InfoGen{}
			resp.Listings[i].Comments[j].Info.Title = bulkGetListingCommentsResponse.ListingComments[i].Comments[j].Title
		}
	}
	resp.UserWallet = &ComplexUsecaseResponse_UserWalletGen{}
	resp.UserWallet.Balance = getUserWalletResponse.UserWallet.CoinBalance
	resp.ApiVersion = 42
	resp.CommitId = "123sdlfkj"
	return resp, nil
}
