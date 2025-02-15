/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Match criteria based on a community list
type CommunityMatchCriteria struct {
	// Match criteria specified as a community list path or a regular expression. 
	Criteria string `json:"criteria"`
	// Match operator for community list entries. Not valid when a regular expression is specified for criteria. 
	MatchOperator string `json:"match_operator,omitempty"`
}
