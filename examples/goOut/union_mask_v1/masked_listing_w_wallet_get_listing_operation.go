package union_mask_v1


import (
	"context"

	
	"github.com/pkg/errors"

	
	"github.com/carousell/aggproto/examples/goOut/listing"
)


type getListingClient struct {
	client listing.ListingsClient
}


func (c *getListingClient) getListing(ctx context.Context, req *listing.GetListingRequest) ( *listing.GetListingResponse, error) {
	resp, err := c.client.GetListing(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error invoking GetListing")
	}
	return resp, nil
}
