/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Object to identify an uplink based on its type and name
type Uplink struct {
	// Name of this uplink
	UplinkName string `json:"uplink_name"`
	// Type of the uplink
	UplinkType string `json:"uplink_type"`
}
