/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type Traceflow struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// The id of the traceflow round
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// Traceflow result analysis notes
	Analysis []string `json:"analysis,omitempty"`
	Counters *TraceflowObservationCounters `json:"counters,omitempty"`
	LogicalCounters *TraceflowObservationCounters `json:"logical_counters,omitempty"`
	// id of the source logical port used for injecting the traceflow packet
	LportId string `json:"lport_id,omitempty"`
	// Represents the traceflow operation state
	OperationState string `json:"operation_state,omitempty"`
	// The status of the traceflow RPC request. SUCCESS - The traceflow request is sent successfully. TIMEOUT - The traceflow request gets timeout. SOURCE_PORT_NOT_FOUND - The source port of the request cannot be found. DATA_PATH_NOT_READY - The datapath component cannot be ready to receive request. CONNECTION_ERROR - There is connection error on datapath component. UNKNOWN - The status of traceflow request cannot be determined.
	RequestStatus string `json:"request_status,omitempty"`
	// A flag, when set true, indicates some observations were deleted from the result set.
	ResultOverflowed bool `json:"result_overflowed,omitempty"`
	// Maximum time (in ms) the management plane will be waiting for this traceflow round. Upper limit for federation case is 90000, for non-federation case is 15000, the maximum is set to 90000 as the higher of the two cases.
	Timeout int64 `json:"timeout,omitempty"`
}
