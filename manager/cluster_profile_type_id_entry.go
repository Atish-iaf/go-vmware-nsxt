/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type ClusterProfileTypeIdEntry struct {

	// key value
	ProfileId string `json:"profile_id"`

	// Supported cluster profiles.
	ResourceType string `json:"resource_type,omitempty"`
}
