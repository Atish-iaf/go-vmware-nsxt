/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type ContainerClusterStatus struct {
	// Identifier of the container cluster.
	ClusterId string `json:"cluster_id,omitempty"`
	// Detail information on status.
	Detail string `json:"detail,omitempty"`
	// Display the cluster check interval in seconds.
	Interval int32 `json:"interval,omitempty"`
	// Display the container cluster status.
	Status string `json:"status,omitempty"`
}
