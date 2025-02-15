/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// List of String values
type StringArrayConstraintValue struct {
	ResourceType string `json:"resource_type"`
	// Array of string values
	Values []string `json:"values"`
}
