/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// The reponse to a batch operation
type BatchResponse struct {
	// Indicates if any of the APIs failed
	HasErrors bool `json:"has_errors,omitempty"`
	// Bulk list results
	Results []BatchResponseItem `json:"results"`
	// Optional flag indicating that all items were rolled back even if succeeded initially
	RolledBack bool `json:"rolled_back,omitempty"`
}
