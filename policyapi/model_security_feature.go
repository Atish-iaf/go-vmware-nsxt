/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// T1 Security feature entity with feature details
type SecurityFeature struct {
	// true - enable the feature, false - disable the feture
	Enable bool `json:"enable"`
	// Feature to be enabled/disabled. IDPS - Intrusion Detection System TLS - Transport Layer Security Inspection MALWAREPREVENTION - Malware Prevention Use any one of this to enable/disabe it. 
	Feature string `json:"feature"`
}
