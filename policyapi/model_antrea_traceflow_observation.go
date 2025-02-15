/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Observation result for Antrea traceflow. 
type AntreaTraceflowObservation struct {
	// The type of component. 
	ComponentType string `json:"component_type,omitempty"`
	// UID of the container node that observed a traceflow packet. 
	ContainerNodeId string `json:"container_node_id,omitempty"`
	// The type of observation. AntreaTraceflowObservationDelivered: The packet was delivered to destination Pod properly AntreaTraceflowObservationReceived: The packet was received from another ContainerNode AntreaTraceflowObservationForwarded: The packet was forwarded to next logical node or ContainerNode AntreaTraceflowObservationDropped: The packet was dropped 
	ObservationType string `json:"observation_type"`
	// Timestamp when the observation was collect by Antrea controller. 
	Timestamp int64 `json:"timestamp,omitempty"`
}
