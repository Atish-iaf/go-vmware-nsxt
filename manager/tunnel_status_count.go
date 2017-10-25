/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type TunnelStatusCount struct {

	// Degraded count
	DegradedCount int32 `json:"degraded_count,omitempty"`

	// Down count
	DownCount int32 `json:"down_count,omitempty"`

	// Roll-up status
	Status string `json:"status,omitempty"`

	// Up count
	UpCount int32 `json:"up_count,omitempty"`

	// BFD Diagnostic
	BfdDiagnostic *BfdDiagnosticCount `json:"bfd_diagnostic,omitempty"`

	// BFD Status
	BfdStatus *BfdStatusCount `json:"bfd_status,omitempty"`
}
