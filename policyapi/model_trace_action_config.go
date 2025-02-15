/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type TraceActionConfig struct {
	ActionArgument *TraceActionArgument `json:"action_argument,omitempty"`
	SamplingArgument *SamplingArgument `json:"sampling_argument,omitempty"`
}
