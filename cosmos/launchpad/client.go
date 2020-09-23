package launchpad

import (
	"context"
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
)

type CosmosAPI struct {
	Auth         CosmosAuthAPI
	Bank         CosmosBankAPI
	Tendermint   CosmosTendermintAPI
	Transactions CosmosTransactionsAPI
}

type CosmosTransactionsAPI interface {
	TxsHashGet(ctx context.Context, hash string) (cosmosclient.TxQuery, *http.Response, error)
	TxsPost(ctx context.Context, txBroadcast cosmosclient.InlineObject) (cosmosclient.BroadcastTxCommitResult, *http.Response, error)
}
type CosmosBankAPI interface {
	BankBalancesAddressGet(ctx context.Context, address string) (cosmosclient.InlineResponse2005, *http.Response, error)
}

type CosmosAuthAPI interface {
	AuthAccountsAddressGet(ctx context.Context, address string) (cosmosclient.InlineResponse2006, *http.Response, error)
}

type CosmosTendermintAPI interface {
	NodeInfoGet(ctx context.Context) (cosmosclient.InlineResponse200, *http.Response, error)
}

type TendermintClient interface {
	NetInfo() (tendermint.NetInfoResponse, error)
	Block(height uint64) (tendermint.BlockResponse, error)
	BlockByHash(hash string) (tendermint.BlockResponse, error)
	UnconfirmedTxs() (tendermint.UnconfirmedTxsResponse, error)
	Tx(hash string) (tendermint.TxResponse, error)
	TxSearch(query string) (tendermint.TxSearchResponse, error)
}
