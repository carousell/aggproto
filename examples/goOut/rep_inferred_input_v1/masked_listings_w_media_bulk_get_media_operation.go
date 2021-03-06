// Code generated by aggproto. DO NOT EDIT.
package rep_inferred_input_v1


import (
	"context"

	
	"github.com/pkg/errors"

	
	"github.com/carousell/aggproto/examples/goOut/media"
)


type bulkGetMediaClient struct {
	client media.MediaServiceClient
}


func (c *bulkGetMediaClient) bulkGetMedia(ctx context.Context, req *media.BulkGetMediaRequest) ( *media.BulkGetMediaResponse, error) {
	resp, err := c.client.BulkGetMedia(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error invoking BulkGetMedia")
	}
	return resp, nil
}
