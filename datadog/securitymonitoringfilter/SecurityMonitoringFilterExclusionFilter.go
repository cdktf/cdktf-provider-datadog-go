// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringfilter


type SecurityMonitoringFilterExclusionFilter struct {
	// Exclusion filter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/security_monitoring_filter#name SecurityMonitoringFilter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Exclusion filter query. Logs that match this query are excluded from the security filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/security_monitoring_filter#query SecurityMonitoringFilter#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

