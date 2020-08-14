/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BlockQueryBlockMeta struct for BlockQueryBlockMeta
type BlockQueryBlockMeta struct {
	Header BlockHeader `json:"header,omitempty"`
	BlockId BlockId `json:"block_id,omitempty"`
}
