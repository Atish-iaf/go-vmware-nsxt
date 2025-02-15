/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Dscp bit config
type DscpValue struct {
	// The method for indicating the existence of INT header.
	IndicatorType string `json:"indicator_type"`
	// A DSCP value is allocated to indicate the existence of INT header. It takes effects only when the INT indicator mode is DSCP_VALUE. The user should guarantee that the given DSCP value is specifically allocated for INT. 
	DscpValue int32 `json:"dscp_value"`
}
