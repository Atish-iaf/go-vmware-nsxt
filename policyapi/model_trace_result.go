/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type TraceResult struct {
	// Trace action result analysis notes
	Analysis []string `json:"analysis,omitempty"`
	Counters *TraceflowObservationCounters `json:"counters,omitempty"`
	// Direction of a trace
	Direction string `json:"direction,omitempty"`
	LogicalCounters *TraceflowObservationCounters `json:"logical_counters,omitempty"`
	// Trace observation list
	Observations []TraceflowObservation `json:"observations,omitempty"`
	// Packet ID in the session
	PacketId string `json:"packet_id,omitempty"`
	// Whether some observations were deleted from the result set
	ResultOverflowed bool `json:"result_overflowed,omitempty"`
}
