package inferred_input_v1
import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/media"
)
func transformToGetMediaRequest(getListingResponse *listing.GetListingResponse) *media.GetMediaRequest{
	getMediaRequest := &media.GetMediaRequest{}
	return getMediaRequest
}