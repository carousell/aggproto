package inferred_input_v1

import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/media"
)

func adaptMaskedListingWMediaResponse(getListingResponse *listing.GetListingResponse, getMediaResponse *media.GetMediaResponse) *MaskedListingWMediaResponse{
	listing := getListingResponse.Listing
	media := getMediaResponse.Media
	resp := &MaskedListingWMediaResponse{}
	resp.Listing = &MaskedListingWMediaResponse_ListingGen{}
	resp.Listing.Title = listing.Title
	resp.Listing.Description = listing.Description
	resp.Listing.PhotoUrl = media.PhotoUrl
	return resp
}
