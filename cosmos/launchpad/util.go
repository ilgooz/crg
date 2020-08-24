package launchpad

import (
	"context"
	"strconv"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// getTxByHash calls
func (l Launchpad) getTxByHash(ctx context.Context, hash string) (*types.Transaction, *types.Error) {
	txQuery, _, err := l.cosmos.Transactions.TxsHashGet(ctx, hash)
	if err != nil {
		return nil, ErrNodeConnection
	}

	tx := cosmosTxToRosettaTx(txQuery)

	return tx, nil
}
func toBlockIdentifier(result tendermintclient.BlockComplete) (*types.BlockIdentifier, error) {
	if result.BlockId.Hash == "" {
		return nil, nil
	}
	height, err := strconv.ParseUint(result.Block.Header.Height, 10, 64)
	if err != nil {
		return nil, err
	}
	return &types.BlockIdentifier{
		Index: int64(height),
		Hash:  result.BlockId.Hash,
	}, nil
}

func toTransactions(txs []cosmosclient.TxQuery) (transactions []*types.Transaction, err error) {
	for _, tx := range txs {
		transactions = append(transactions, cosmosTxToRosettaTx(tx))
	}
	return
}

// tendermintTxToRosettaTx converts a Tendermint api TxResponseResult to a Transaction
// in the type expected by Rosetta.
func tendermintTxToRosettaTx(res tendermintclient.TxResponseResult) *types.Transaction {
	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: res.Hash,
		},
		Operations: nil, // TODO difficult to get the operations from the mempool (maybe not worth it due to block times).
	}
}

// cosmosTxToRosettaTx converts a Cosmos api TxQuery to a Transaction
// in the type expected by Rosetta.
func cosmosTxToRosettaTx(tx cosmosclient.TxQuery) *types.Transaction {
	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: tx.Txhash,
		},
		Operations: toOperations(tx.Tx.Value.Msg),
	}
}

func toOperations(msg []cosmosclient.Msg) (operations []*types.Operation) {
	for i, msg := range msg {
		account := msg.Value.Creator
		if account == "" {
			account = msg.Value.FromAddress
		}
		var amount *types.Amount
		amounts := msg.Value.Amount
		if len(amounts) > 0 {
			am := amounts[0]
			amount = &types.Amount{
				Value: am.Amount,
				Currency: &types.Currency{
					Symbol: am.Denom,
				},
			}
		}
		operations = append(operations, &types.Operation{
			OperationIdentifier: &types.OperationIdentifier{
				Index: int64(i),
			},
			Type:   msg.Type,
			Status: "TODO",
			Account: &types.AccountIdentifier{
				Address: account,
			},
			Amount: amount,
		})
	}
	return
}