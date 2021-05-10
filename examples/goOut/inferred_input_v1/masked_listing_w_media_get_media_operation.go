package inferred_input_v1


import (
	"context"

	
	"github.com/pkg/errors"

	
	"github.com/carousell/aggproto/examples/goOut/media"
)


type getMediaClient struct {
	client media.MediaServiceClient
}


func (c *getMediaClient) getMedia(ctx context.Context, req *media.GetMediaRequest) ( *media.GetMediaResponse, error) {
	resp, err := c.client.GetMedia(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error invoking GetMedia")
	}
	return resp, nil
}
