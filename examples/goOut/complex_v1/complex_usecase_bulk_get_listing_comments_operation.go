package complex_v1

import (
	"context"
	
	"github.com/pkg/errors"
	
	"github.com/carousell/aggproto/examples/goOut/listing_comments"
)

type bulkGetListingCommentsClient struct {
	client listing_comments.ListingCommentsClient
}

func (c *bulkGetListingCommentsClient) bulkGetListingComments(ctx context.Context, req *listing_comments.BulkGetListingCommentsRequest) ( *listing_comments.BulkGetListingCommentsResponse, error) {
	resp, err := c.client.BulkGetListingComments(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error invoking BulkGetListingComments")
	}
	return resp, nil
}
