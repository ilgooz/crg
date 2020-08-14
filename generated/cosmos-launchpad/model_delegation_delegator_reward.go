/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// DelegationDelegatorReward struct for DelegationDelegatorReward
type DelegationDelegatorReward struct {
	// bech32 encoded address
	ValidatorAddress string `json:"validator_address,omitempty"`
	Reward []Coin `json:"reward,omitempty"`
}
