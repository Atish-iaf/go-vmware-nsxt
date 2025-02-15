/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Guest virtual machine details include OS name and computer name of guest VM. 
type GuestInfo struct {
	// Computer name of guest virtual machine, which is set inside guest OS. Currently this is supported for guests on ESXi that have VMware Tools installed. 
	ComputerName string `json:"computer_name,omitempty"`
	// OS name of guest virtual machine. Currently this is supported for guests on ESXi that have VMware Tools installed. 
	OsName string `json:"os_name,omitempty"`
}
