package rep_inferred_input_v1


import (
	"github.com/pkg/errors"

	
	"github.com/carousell/aggproto/examples/goOut/media"
	"github.com/carousell/aggproto/examples/goOut/listing"
)


func adaptMaskedListingsWMediaResponse(bulkGetListingsResponse *listing.BulkGetListingsResponse, bulkGetMediaResponse *media.BulkGetMediaResponse) (*MaskedListingsWMediaResponse, error){
	resp := &MaskedListingsWMediaResponse{}
	if len(bulkGetListingsResponse.Listings) != len(bulkGetMediaResponse.Medias) {
		return nil, errors.Errorf("assertion failed %s != %s", len(bulkGetListingsResponse.Listings), len(bulkGetMediaResponse.Medias))
	}
	resp.Listings = make([]*MaskedListingsWMediaResponse_ListingsGen, len(bulkGetListingsResponse.Listings))
	for i := 0; i < len(bulkGetListingsResponse.Listings); i++ {
		resp.Listings[i] = &MaskedListingsWMediaResponse_ListingsGen{}
		resp.Listings[i].Title = bulkGetListingsResponse.Listings[i].Title
		resp.Listings[i].Description = bulkGetListingsResponse.Listings[i].Description
		resp.Listings[i].PhotoUrl = bulkGetMediaResponse.Medias[i].PhotoUrl
	}
	return resp, nil
}
