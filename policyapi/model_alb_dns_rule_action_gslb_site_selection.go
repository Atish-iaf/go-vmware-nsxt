/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer DnsRuleActionGslbSiteSelection object
type AlbDnsRuleActionGslbSiteSelection struct {
	// GSLB fallback sites to use in case the desired site is down. Maximum of 64 items allowed. 
	FallbackSiteNames []string `json:"fallback_site_names,omitempty"`
	// When set to true, GSLB site is a preferred site. This setting comes into play when the site is down, as well as no configured fallback site is available (all fallback sites are also down), then any one available site is selected based on the default algorithm for GSLB pool member selection. Default value when not specified in API or module is interpreted by ALB Controller as true. 
	IsSitePreferred bool `json:"is_site_preferred,omitempty"`
	// GSLB site name.
	SiteName string `json:"site_name"`
}
