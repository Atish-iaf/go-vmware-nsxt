/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer DnsGeoLocationMatch object
type AlbDnsGeoLocationMatch struct {
	// Geographical location of the client IP to be used in the match. This location is of the format Country/State/City e.g. US/CA/Santa Clara. 
	GeolocationName string `json:"geolocation_name,omitempty"`
	// Geolocation tag for the client IP. This could be any string value for the client IP, e.g. client IPs from US East Coast geolocation would be tagged as 'East Coast'. 
	GeolocationTag string `json:"geolocation_tag,omitempty"`
	// Criterion to use for matching the client IP's geographical location. Enum options - IS_IN, IS_NOT_IN. 
	MatchCriteria string `json:"match_criteria"`
	// Use the IP address from the EDNS client subnet option, if available, to derive geo location of the DNS query. Default value when not specified in API or module is interpreted by ALB Controller as true. 
	UseEdnsClientSubnetIp bool `json:"use_edns_client_subnet_ip,omitempty"`
}
