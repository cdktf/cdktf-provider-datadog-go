// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSplitGraphDefinitionSplitConfigSort struct {
	// Widget sorting methods. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.31.0/docs/resources/dashboard#order Dashboard#order}
	Order *string `field:"required" json:"order" yaml:"order"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.31.0/docs/resources/dashboard#compute Dashboard#compute}
	Compute *DashboardWidgetSplitGraphDefinitionSplitConfigSortCompute `field:"optional" json:"compute" yaml:"compute"`
}

