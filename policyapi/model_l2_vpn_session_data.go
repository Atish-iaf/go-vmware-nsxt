/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// L2VPN Session Data represents meta data necessary to create the L2VPN Session. It is represented by an array of peer code for each tunnel. 
type L2VpnSessionData struct {
	// Description of L2VPN Session
	Description string `json:"description,omitempty"`
	// Defaults to id if not set.
	DisplayName string `json:"display_name,omitempty"`
	// Enable to extend all the associated segments.
	Enabled bool `json:"enabled,omitempty"`
	// List of L2VPN transport tunnel data.
	TransportTunnels []L2VpnSessionTransportTunnelData `json:"transport_tunnels,omitempty"`
}
