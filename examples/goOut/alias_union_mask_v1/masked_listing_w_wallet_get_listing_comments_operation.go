package alias_union_mask_v1

import (
	"context"
	
	"github.com/pkg/errors"
	
	"github.com/carousell/aggproto/examples/goOut/listing_comments"
)

type getListingCommentsClient struct {
	client listing_comments.ListingCommentsClient
}

func (c *getListingCommentsClient) getListingComments(ctx context.Context, req *listing_comments.GetListingCommentsRequest) ( *listing_comments.GetListingCommentsResponse, error) {
	resp, err := c.client.GetListingComments(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error invoking GetListingComments")
	}
	return resp, nil
}
