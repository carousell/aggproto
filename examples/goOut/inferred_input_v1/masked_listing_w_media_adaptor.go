package inferred_input_v1


import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/media"
)


func adaptMaskedListingWMediaResponse(getListingResponse *listing.GetListingResponse, getMediaResponse *media.GetMediaResponse) (*MaskedListingWMediaResponse, error){
	resp := &MaskedListingWMediaResponse{}
	resp.Listing = &MaskedListingWMediaResponse_ListingGen{}
	resp.Listing.Title = getListingResponse.Listing.Title
	resp.Listing.Description = getListingResponse.Listing.Description
	resp.Listing.PhotoUrl = getMediaResponse.Media.PhotoUrl
	return resp, nil
}
