// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSunburstDefinitionRequestQueryEventQuerySearch struct {
	// The events search string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/dashboard#query Dashboard#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

