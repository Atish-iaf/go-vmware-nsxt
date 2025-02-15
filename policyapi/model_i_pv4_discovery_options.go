/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Contains IPv4 related discovery options.
type IPv4DiscoveryOptions struct {
	ArpSnoopingConfig *ArpSnoopingConfig `json:"arp_snooping_config,omitempty"`
	// Indicates whether DHCP snooping is enabled
	DhcpSnoopingEnabled bool `json:"dhcp_snooping_enabled,omitempty"`
	// Indicates whether fetching IP using vm-tools is enabled. This option is only supported on ESX where vm-tools is installed. 
	VmtoolsEnabled bool `json:"vmtools_enabled,omitempty"`
}
