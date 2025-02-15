/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer SamlServiceProviderNode object
type AlbSamlServiceProviderNode struct {
	// Globally unique entityID for this node. Entity ID on the IDP should match this. 
	EntityId string `json:"entity_id,omitempty"`
	// Refers to the Cluster name identifier (Virtual IP or FQDN).
	Name string `json:"name"`
	// Service Engines will use this SSL certificate to sign assertions going to the IdP. It is a reference to an object of type SSLKeyAndCertificate. 
	SigningSslKeyAndCertificatePath string `json:"signing_ssl_key_and_certificate_path,omitempty"`
	// Single Signon URL to be programmed on the IDP.
	SingleSignonUrl string `json:"single_signon_url,omitempty"`
}
