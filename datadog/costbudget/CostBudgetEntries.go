// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package costbudget


type CostBudgetEntries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/cost_budget#amount CostBudget#amount}.
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/cost_budget#month CostBudget#month}.
	Month *float64 `field:"required" json:"month" yaml:"month"`
	// tag_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/cost_budget#tag_filters CostBudget#tag_filters}
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

