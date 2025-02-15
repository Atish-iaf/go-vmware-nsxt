/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Query statistics counters to an upstream server including successfully forwarded queries and failed queries. 
type NsxTUpstreamServerStatistics struct {
	// Queries failed to forward.
	QueriesFailed int64 `json:"queries_failed,omitempty"`
	// Queries forwarded successfully
	QueriesSucceeded int64 `json:"queries_succeeded,omitempty"`
	// Upstream server ip
	UpstreamServer string `json:"upstream_server,omitempty"`
}
