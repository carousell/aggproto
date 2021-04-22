package complex_v1

import (
	"github.com/pkg/errors"
	
	"github.com/carousell/aggproto/examples/goOut/wallet"
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/media"
)

func adaptComplexUsecaseResponse(bulkGetListingsResponse *listing.BulkGetListingsResponse, bulkGetMediaResponse *media.BulkGetMediaResponse, getUserWalletResponse *wallet.GetUserWalletResponse) (*ComplexUsecaseResponse, error){
	resp := &ComplexUsecaseResponse{}
	if len(bulkGetListingsResponse.Listings) != len(bulkGetMediaResponse.Medias) {
		return nil, errors.Errorf("assertion failed %s != %s", len(bulkGetListingsResponse.Listings), len(bulkGetMediaResponse.Medias))
	}
	resp.Listings = make([]*ComplexUsecaseResponse_ListingsGen, len(bulkGetListingsResponse.Listings))
	for i := 0; i < len(bulkGetListingsResponse.Listings); i++ {
		resp.Listings[i] = &ComplexUsecaseResponse_ListingsGen{}
		resp.Listings[i].Info = &ComplexUsecaseResponse_ListingsGen_InfoGen{}
		resp.Listings[i].Info.Title = bulkGetListingsResponse.Listings[i].Title
		resp.Listings[i].Info.Description = bulkGetListingsResponse.Listings[i].Description
		resp.Listings[i].Media = &ComplexUsecaseResponse_ListingsGen_MediaGen{}
		resp.Listings[i].Media.PhotoUrl = bulkGetMediaResponse.Medias[i].PhotoUrl
	}
	resp.UserWallet = &ComplexUsecaseResponse_UserWalletGen{}
	resp.UserWallet.Balance = getUserWalletResponse.UserWallet.CoinBalance
	resp.ApiVersion = 42
	resp.CommitId = "123sdlfkj"
	return resp, nil
}
