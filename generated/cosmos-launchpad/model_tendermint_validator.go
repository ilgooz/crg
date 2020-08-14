/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TendermintValidator struct for TendermintValidator
type TendermintValidator struct {
	// bech32 encoded address
	Address string `json:"address,omitempty"`
	PubKey string `json:"pub_key,omitempty"`
	VotingPower string `json:"voting_power,omitempty"`
	ProposerPriority string `json:"proposer_priority,omitempty"`
}
