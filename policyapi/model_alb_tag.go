/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer Tag object
type AlbTag struct {
	// Enum options - AVI_DEFINED, USER_DEFINED, VCENTER_DEFINED. Default value when not specified in API or module is interpreted by ALB Controller as USER_DEFINED. 
	Type_ string `json:"type,omitempty"`
	// value of Tag.
	Value string `json:"value"`
}
