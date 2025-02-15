/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type NdraAdvertisedRoute struct {
	// Lifetime of advertised route in seconds. 
	RouteLifetime int64 `json:"route_lifetime,omitempty"`
	// NDRA Route preference. Indicates preference of the router associated with a prefix over others, when multiple identical prefixes (for different routers) have been received. 
	RoutePreference string `json:"route_preference,omitempty"`
	// Advertised route subnet 
	Subnet string `json:"subnet"`
}
