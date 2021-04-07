package simple_mask_v1

import "context"

type maskedListingSvc struct {
	UnimplementedMaskedListingServiceServer
}

func New() MaskedListingServiceServer {
	return &maskedListingSvc{}
}

func (s *maskedListingSvc) InvokeMaskedListing(ctx context.Context, req *MaskedListingRequest) (*MaskedListingResponse, error) {
}
