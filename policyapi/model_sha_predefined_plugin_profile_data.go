/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Describes a config item for System Health profile. 
type ShaPredefinedPluginProfileData struct {
	// The interval of plugin to check the status.
	CheckInterval int64 `json:"check_interval,omitempty"`
	// The interval of plugin to report the status.
	ReportInterval int64 `json:"report_interval,omitempty"`
	// The smallest report interval if the status is changed. The value of smallest_report_interval_if_change should be less than the value of report_interval 
	SmallestReportIntervalIfChange int64 `json:"smallest_report_interval_if_change,omitempty"`
}
