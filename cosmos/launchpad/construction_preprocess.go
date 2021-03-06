package launchpad

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) ConstructionPreprocess(ctx context.Context, r *types.ConstructionPreprocessRequest) (*types.ConstructionPreprocessResponse, *types.Error) {
	operations := r.Operations
	if len(operations) < 1 {
		return nil, ErrInterpreting
	}

	txData, err := getTransferTxDataFromOperations(operations)
	if err != nil {
		return nil, ErrInvalidAddress
	}
	if txData.From == nil {
		return nil, ErrInvalidAddress
	}

	var res = &types.ConstructionPreprocessResponse{
		Options: map[string]interface{}{
			OptionAddress: txData.From.String(),
			OptionGas:     r.SuggestedFeeMultiplier,
			// TODO: Check if memo is needed
		},
	}
	return res, nil
}
