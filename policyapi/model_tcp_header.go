/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type TcpHeader struct {
	// Destination port of tcp header
	DstPort int64 `json:"dst_port,omitempty"`
	// Source port of tcp header
	SrcPort int64 `json:"src_port,omitempty"`
	// TCP flags (9bits)
	TcpFlags int64 `json:"tcp_flags,omitempty"`
}
