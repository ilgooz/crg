// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	openapi "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
)

// CosmosAuthAPI is an autogenerated mock type for the CosmosAuthAPI type
type CosmosAuthAPI struct {
	mock.Mock
}

// AuthAccountsAddressGet provides a mock function with given fields: ctx, address
func (_m *CosmosAuthAPI) AuthAccountsAddressGet(ctx context.Context, address string) (openapi.InlineResponse2006, *http.Response, error) {
	ret := _m.Called(ctx, address)

	var r0 openapi.InlineResponse2006
	if rf, ok := ret.Get(0).(func(context.Context, string) openapi.InlineResponse2006); ok {
		r0 = rf(ctx, address)
	} else {
		r0 = ret.Get(0).(openapi.InlineResponse2006)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string) *http.Response); ok {
		r1 = rf(ctx, address)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, address)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
