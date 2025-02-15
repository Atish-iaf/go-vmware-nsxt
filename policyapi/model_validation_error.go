/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Error while applying transport node profile on discovered node
type ValidationError struct {
	// Discovered Node Id
	DiscoveredNodeId string `json:"discovered_node_id,omitempty"`
	// Validation error message
	ErrorMessage string `json:"error_message,omitempty"`
}
