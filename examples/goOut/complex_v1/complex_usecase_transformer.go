package complex_v1
import (
	"github.com/carousell/aggproto/examples/goOut/listing"
	"github.com/carousell/aggproto/examples/goOut/media"
)
func transformToBulkGetMediaRequest(bulkGetMediaRequest *media.BulkGetMediaRequest, bulkGetListingsResponse *listing.BulkGetListingsResponse) {
	bulkGetMediaRequest.MediaIds = make([]string, len(bulkGetListingsResponse.Listings))
	for i:=0; i < len(bulkGetListingsResponse.Listings); i+=1 {
		bulkGetMediaRequest.MediaIds[i] = bulkGetListingsResponse.Listings[i].MediaId
	}
	return 
}
