/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Detailed information about an API Error
type PolicyApiError struct {
	// Further details about the error
	Details string `json:"details,omitempty"`
	// A numeric error code
	ErrorCode int64 `json:"error_code,omitempty"`
	// Additional data about the error
	ErrorData interface{} `json:"error_data,omitempty"`
	// A description of the error
	ErrorMessage string `json:"error_message,omitempty"`
	// The module name where the error occurred
	ModuleName string `json:"module_name,omitempty"`
	// Other errors related to this error
	RelatedErrors []PolicyRelatedApiError `json:"related_errors,omitempty"`
}
