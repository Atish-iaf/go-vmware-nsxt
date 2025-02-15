/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer AuthProfileHTTPClientParams object
type AlbAuthProfileHttpClientParams struct {
	// The max allowed length of time a clients authentication is cached. Allowed values are 1-30. Unit is SEC. Default value when not specified in API or module is interpreted by ALB Controller as 5. 
	CacheExpirationTime int64 `json:"cache_expiration_time,omitempty"`
	// Insert an HTTP header. This field is used to define the header name. The value of the header is set to the client's HTTP Auth user ID. 
	RequestHeader string `json:"request_header,omitempty"`
	// A user should be a member of these groups. Each group is defined by the DN. For example, CN=testgroup,OU=groups,dc=example,dc=avinetworks,DC=com. 
	RequireUserGroups []string `json:"require_user_groups,omitempty"`
}
