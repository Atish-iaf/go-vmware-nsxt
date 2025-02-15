/*
 * NSX Manager API
 *
 * VMware NSX Manager REST API
 *
 * API version: 4.1.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package containerinventory

// It represents the condition of a service point.
type PortStatus struct {
	// Specifies port of service point.
	Port int64 `json:"port,omitempty"`
	// Specifies protocol of service point. e.g. TCP, UDP, SCTP.
	Protocol string `json:"protocol,omitempty"`
}
