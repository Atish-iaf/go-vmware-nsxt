/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type PolicyTraceflowObservationDroppedLogical struct {
	// The name of the component that issued the observation.
	ComponentName string `json:"component_name,omitempty"`
	// The sub type of the component that issued the observation.
	ComponentSubType string `json:"component_sub_type,omitempty"`
	// The type of the component that issued the observation.
	ComponentType string `json:"component_type,omitempty"`
	ResourceType string `json:"resource_type"`
	// the hop count for observations on the transport node that a traceflow packet is injected in will be 0. The hop count is incremented each time a subsequent transport node receives the traceflow packet. The sequence number of 999 indicates that the hop count could not be determined for the containing observation.
	SequenceNo int64 `json:"sequence_no,omitempty"`
	// Timestamp when the observation was created by the transport node (milliseconds epoch)
	Timestamp int64 `json:"timestamp,omitempty"`
	// Timestamp when the observation was created by the transport node (microseconds epoch)
	TimestampMicro int64 `json:"timestamp_micro,omitempty"`
	// id of the transport node that observed a traceflow packet
	TransportNodeId string `json:"transport_node_id,omitempty"`
	// name of the transport node that observed a traceflow packet
	TransportNodeName string `json:"transport_node_name,omitempty"`
	// type of the transport node that observed a traceflow packet
	TransportNodeType string `json:"transport_node_type,omitempty"`
	// This field is specified when the traceflow packet matched a L3 firewall rule. 
	AclRuleId int64 `json:"acl_rule_id,omitempty"`
	// This field specifies the ARP fails reason ARP_TIMEOUT - ARP failure due to query control plane timeout ARP_CPFAIL - ARP failure due post ARP query message to control plane failure ARP_FROMCP - ARP failure due to deleting ARP entry from control plane ARP_PORTDESTROY - ARP failure due to port destruction ARP_TABLEDESTROY - ARP failure due to ARP table destruction ARP_NETDESTROY - ARP failure due to overlay network destruction
	ArpFailReason string `json:"arp_fail_reason,omitempty"`
	// This field is specified when the traceflow packet matched a jump-to rule. 
	JumptoRuleId int64 `json:"jumpto_rule_id,omitempty"`
	// This field is specified when the traceflow packet matched a l2 rule. 
	L2RuleId int64 `json:"l2_rule_id,omitempty"`
	// The id of the logical port at which the traceflow packet was dropped
	LportId string `json:"lport_id,omitempty"`
	// The name of the logical port at which the traceflow packet was dropped
	LportName string `json:"lport_name,omitempty"`
	// This field is specified when the traceflow packet matched a NAT rule. 
	NatRuleId int64 `json:"nat_rule_id,omitempty"`
	// The reason traceflow packet was dropped
	Reason string `json:"reason,omitempty"`
	// The id of the component that dropped the traceflow packet.
	ComponentId string `json:"component_id,omitempty"`
	// The index of service path that is a chain of services represents the point where the traceflow packet was dropped. 
	ServicePathIndex int64 `json:"service_path_index,omitempty"`
	// The path of the ACL rule that was applied to forward the traceflow packet
	AclRulePath string `json:"acl_rule_path,omitempty"`
	// The path of the component that dropped the traceflow packet
	ComponentPath string `json:"component_path,omitempty"`
	// The path of the jump-to rule that was applied to the traceflow packet
	JumptoRulePath string `json:"jumpto_rule_path,omitempty"`
	// The path of the l2 rule that was applied to the traceflow packet
	L2RulePath string `json:"l2_rule_path,omitempty"`
	// The path of the NAT rule that was applied to forward the traceflow packet
	NatRulePath string `json:"nat_rule_path,omitempty"`
	// The path of the segment port at which traceflow packet was dropped
	SegmentPortPath string `json:"segment_port_path,omitempty"`
}
