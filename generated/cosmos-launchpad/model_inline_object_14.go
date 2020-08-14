/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject14 struct for InlineObject14
type InlineObject14 struct {
	BaseReq BaseReq `json:"base_req,omitempty"`
	// bech32 encoded address
	WithdrawAddress string `json:"withdraw_address,omitempty"`
}
