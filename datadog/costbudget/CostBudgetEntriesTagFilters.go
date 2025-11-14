// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package costbudget


type CostBudgetEntriesTagFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/cost_budget#tag_key CostBudget#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/cost_budget#tag_value CostBudget#tag_value}.
	TagValue *string `field:"required" json:"tagValue" yaml:"tagValue"`
}

