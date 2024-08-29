/*
 * xoxoapi.proto
 *
 * The realtime protocol for the template match engine.
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type RuntimeError struct {

	Error string `json:"error,omitempty"`

	Code int32 `json:"code,omitempty"`

	Message string `json:"message,omitempty"`

	Details []ProtobufAny `json:"details,omitempty"`
}
