/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package upgrade

import (
	"github.com/vmware/go-vmware-nsxt/common"
)

type AwsSubnet struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Timestamp of last modification
	LastSyncTime int64 `json:"_last_sync_time,omitempty"`

	// Description of this resource
	Description string `json:"description,omitempty"`

	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`

	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`

	// Opaque identifiers meaningful to the API user
	Tags []common.Tag `json:"tags,omitempty"`

	AvailabilityZone string `json:"availability_zone,omitempty"`

	// IPV4 CIDR Block for the Vpc
	Cidr string `json:"cidr"`

	// ID of subnet
	Id string `json:"id"`

	// ID of the vpc
	VpcId string `json:"vpc_id"`
}
