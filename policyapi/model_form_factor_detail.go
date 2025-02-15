/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Form factor contains, resources required to deploy NSX Application Platform deployment and available features for a given form factor. 
type FormFactorDetail struct {
	NodeResources *NodeResources `json:"node_resources,omitempty"`
	// Features supported in this form factor. 
	SupportedFeatures []string `json:"supported_features,omitempty"`
}
