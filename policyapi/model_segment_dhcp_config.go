/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// DHCP IPv4 and IPv6 configurations are extended from this abstract class. 
type SegmentDhcpConfig struct {
	// IP address of DNS servers for subnet. DNS server IP address must belong to the same address family as segment gateway_address property. 
	DnsServers []string `json:"dns_servers,omitempty"`
	// DHCP lease time in seconds. When specified, this property overwrites lease time configured DHCP server config. 
	LeaseTime int64 `json:"lease_time,omitempty"`
	ResourceType string `json:"resource_type"`
	// IP address of the DHCP server in CIDR format. The server_address is mandatory in case this segment has provided a dhcp_config_path and it represents a DHCP server config. If this SegmentDhcpConfig is a SegmentDhcpV4Config, the address must be an IPv4 address. If this is a SegmentDhcpV6Config, the address must be an IPv6 address. This address must not overlap the ip-ranges of the subnet, or the gateway address of the subnet, or the DHCP static-binding addresses of this segment. 
	ServerAddress string `json:"server_address,omitempty"`
}
