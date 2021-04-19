package static_primitives_v1
func adaptMockListingResponse() *MockListingResponse{
	resp := &MockListingResponse{}
	resp.Listing = &MockListingResponse_ListingGen{}
	resp.Listing.Title = "iPhone"
	resp.Listing.Description = "BNIB iPhone X"
	return resp
}
