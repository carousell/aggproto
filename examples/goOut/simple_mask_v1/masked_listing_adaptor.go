package simple_mask_v1

import (
	"github.com/carousell/aggproto/examples/goOut/listing"
)

func adaptMaskedListingResponse(getListingResponse *listing.GetListingResponse) *MaskedListingResponse{
	listing := getListingResponse.Listing
	resp := &MaskedListingResponse{}
	resp.Listing = &MaskedListingResponse_ListingGen{}
	resp.Listing.GetListingResponse = &MaskedListingResponse_ListingGen_GetListingResponseGen{}
	resp.Listing.GetListingResponse.Listing = &MaskedListingResponse_ListingGen_GetListingResponseGen_ListingGen{}
	resp.Listing.GetListingResponse.Listing.Title = listing.Title
	resp.Listing.GetListingResponse.Listing.Description = listing.Description
	return resp
}
