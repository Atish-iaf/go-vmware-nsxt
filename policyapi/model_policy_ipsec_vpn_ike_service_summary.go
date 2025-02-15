/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Summarized view of all IPSec VPN sessions for a specified service.
type PolicyIpsecVpnIkeServiceSummary struct {
	AggregateTrafficCounters *IpSecVpnTrafficCounters `json:"aggregate_traffic_counters,omitempty"`
	// Timestamp when the data was last updated.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	SessionSummary *IPsecVpnikeSessionSummary `json:"session_summary,omitempty"`
	// Traffic summary per session.
	TrafficSummaryPerSession []IpSecVpnSessionTrafficSummary `json:"traffic_summary_per_session,omitempty"`
	// Display name of IPSec VPN service
	DisplayName string `json:"display_name,omitempty"`
	// Policy Path referencing the Primary site's enforcement point where the info is fetched. This is applicable only on a GlobalManager. 
	EnforcementPointPath string `json:"enforcement_point_path,omitempty"`
	// Policy path of IPSec VPN service
	IpsecVpnServicePath string `json:"ipsec_vpn_service_path,omitempty"`
}
