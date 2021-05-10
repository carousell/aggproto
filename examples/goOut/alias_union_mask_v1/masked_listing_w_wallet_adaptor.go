package alias_union_mask_v1


import (
	"github.com/carousell/aggproto/examples/goOut/listing_comments"
	"github.com/carousell/aggproto/examples/goOut/wallet"
	"github.com/carousell/aggproto/examples/goOut/listing"
)


func adaptMaskedListingWWalletResponse(getListingResponse *listing.GetListingResponse, getListingCommentsResponse *listing_comments.GetListingCommentsResponse, getUserWalletResponse *wallet.GetUserWalletResponse) (*MaskedListingWWalletResponse, error){
	resp := &MaskedListingWWalletResponse{}
	resp.Listing = &MaskedListingWWalletResponse_ListingGen{}
	resp.Listing.Title = getListingResponse.Listing.Title
	resp.Listing.Description = getListingResponse.Listing.Description
	resp.Listing.Comments = make([]*MaskedListingWWalletResponse_ListingGen_CommentsGen, len(getListingCommentsResponse.ListingComment.Comments))
	for i := 0; i < len(getListingCommentsResponse.ListingComment.Comments); i++ {
		resp.Listing.Comments[i] = &MaskedListingWWalletResponse_ListingGen_CommentsGen{}
		resp.Listing.Comments[i].Title = getListingCommentsResponse.ListingComment.Comments[i].Title
	}
	resp.Wallet = &MaskedListingWWalletResponse_WalletGen{}
	resp.Wallet.Balance = getUserWalletResponse.UserWallet.CoinBalance
	return resp, nil
}
