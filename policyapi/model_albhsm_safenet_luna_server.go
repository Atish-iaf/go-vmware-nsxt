/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer HSMSafenetLunaServer object
type AlbhsmSafenetLunaServer struct {
	// Number of index.
	Index int64 `json:"index"`
	// Password of the partition assigned to this client.
	PartitionPasswd string `json:"partition_passwd,omitempty"`
	// Serial number of the partition assigned to this client.
	PartitionSerialNumber string `json:"partition_serial_number,omitempty"`
	// IP address of the Safenet/Gemalto HSM device.
	RemoteIp string `json:"remote_ip"`
	// CA certificate of the server.
	ServerCert string `json:"server_cert"`
}
