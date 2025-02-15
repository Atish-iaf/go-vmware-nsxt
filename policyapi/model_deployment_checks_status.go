/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Pre/Post deployment status.
type DeploymentChecksStatus struct {
	// Show more information about pre/post check performed. 
	Description string `json:"description,omitempty"`
	// Name of pre/post check.
	Name string `json:"name,omitempty"`
	// Reason for failure of pre/post check. Otherwise empty. 
	Reason string `json:"reason,omitempty"`
	// Status pre/post check. SUCCESS - Successfully completed pre/post check. FAILED - Failed pre/post check. WARNING - Warning in pre/post check. SKIPPED - Pre/post check skipped. IN_PROGRESS - Pre/post check in progress. STOPPING - Stopping pre/post check. STOPPED - Pre/post check stopped. NOT_STARTED - Pre/post check not started 
	Status string `json:"status,omitempty"`
}
