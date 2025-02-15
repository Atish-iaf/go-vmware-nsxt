/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer DnsRateLimiter object
type AlbDnsRateLimiter struct {
	Action *AlbDnsRuleRlAction `json:"action"`
	RateLimiterObject *AlbRateLimiter `json:"rate_limiter_object"`
}
