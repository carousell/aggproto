package complex_v1

import (
	"context"
	
	"github.com/pkg/errors"
	
	"github.com/carousell/aggproto/examples/goOut/listing"
)

type bulkGetListingsClient struct {
	client listing.ListingsClient
}

func (c *bulkGetListingsClient) bulkGetListings(ctx context.Context, req *listing.BulkGetListingsRequest) ( *listing.BulkGetListingsResponse, error) {
	resp, err := c.client.BulkGetListings(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error invoking BulkGetListings")
	}
	return resp, nil
}
