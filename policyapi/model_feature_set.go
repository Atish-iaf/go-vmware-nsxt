/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Represents list of features required to view the widget.
type FeatureSet struct {
	// List of features required for to view widget.
	FeatureList []string `json:"feature_list,omitempty"`
	// Flag for specifying if permission to all features is required If set to false, then if there is permission for any of the feature from feature list, widget will be available.
	RequireAllPermissions bool `json:"require_all_permissions,omitempty"`
}
