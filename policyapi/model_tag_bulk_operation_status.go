/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Status of tag bulk operation. 
type TagBulkOperationStatus struct {
	// Tag apply operation status per resource type. 
	ApplyTo []ResourceTypeTagStatus `json:"apply_to,omitempty"`
	// Intent path corresponding to tag operation
	Path string `json:"path"`
	// Tag remove operation status per resource type. 
	RemoveFrom []ResourceTypeTagStatus `json:"remove_from,omitempty"`
	// Overall status
	Status string `json:"status"`
	Tag *Tag `json:"tag"`
}
