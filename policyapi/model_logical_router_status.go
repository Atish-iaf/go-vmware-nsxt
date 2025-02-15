/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type LogicalRouterStatus struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Egress mode for the logical router at given mode 
	LocaleOperationMode string `json:"locale_operation_mode,omitempty"`
	// The id of the logical router
	LogicalRouterId string `json:"logical_router_id"`
	// Per Node Status
	PerNodeStatus []LogicalRouterStatusPerNode `json:"per_node_status,omitempty"`
}
