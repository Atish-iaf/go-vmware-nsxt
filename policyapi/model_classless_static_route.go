/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// DHCP classless static route option.
type ClasslessStaticRoute struct {
	// Destination network in CIDR format.
	Network string `json:"network"`
	// IP address of next hop of the route.
	NextHop string `json:"next_hop"`
}
