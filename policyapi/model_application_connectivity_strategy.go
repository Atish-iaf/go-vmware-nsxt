/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Allows more granular policies for application workloads 
type ApplicationConnectivityStrategy struct {
	// App connectivity strategies 
	ApplicationConnectivityStrategy string `json:"application_connectivity_strategy"`
	// Based on the value of the app connectivity strategy, a default rule is created for the security policy. The rule id is internally assigned by the system for this default rule. 
	DefaultApplicationRuleId int64 `json:"default_application_rule_id,omitempty"`
	// Flag to enable packet logging. Default is disabled.
	LoggingEnabled bool `json:"logging_enabled,omitempty"`
}
