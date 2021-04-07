package simple_mask_v1
func AdaptMaskedListingResponse(getListingResponse *listing.GetListingResponse) *MaskedListingResponse{
	listing := getListingResponse.listing
	resp := &MaskedListingResponse{}
	resp.Listing = &ListingGen{}
	resp.Listing.GetListingResponse = &GetListingResponseGen{}
	resp.Listing.GetListingResponse.Listing = &ListingGen{}
	resp.Listing.GetListingResponse.Listing.Title = listing.Title
	resp.Listing.GetListingResponse.Listing.Description = listing.Description
	return resp
}
