package complex_v1


import (
	"context"

	
	"github.com/pkg/errors"

	
	"github.com/carousell/aggproto/examples/goOut/wallet"
)


type getUserWalletClient struct {
	client wallet.WalletClient
}


func (c *getUserWalletClient) getUserWallet(ctx context.Context, req *wallet.GetUserWalletRequest) ( *wallet.GetUserWalletResponse, error) {
	resp, err := c.client.GetUserWallet(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error invoking GetUserWallet")
	}
	return resp, nil
}
