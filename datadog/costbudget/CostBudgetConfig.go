// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package costbudget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CostBudgetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The month when the budget ends (YYYYMM).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/cost_budget#end_month CostBudget#end_month}
	EndMonth *float64 `field:"required" json:"endMonth" yaml:"endMonth"`
	// The cost query used to track against the budget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/cost_budget#metrics_query CostBudget#metrics_query}
	MetricsQuery *string `field:"required" json:"metricsQuery" yaml:"metricsQuery"`
	// The name of the budget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/cost_budget#name CostBudget#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The month when the budget starts (YYYYMM).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/cost_budget#start_month CostBudget#start_month}
	StartMonth *float64 `field:"required" json:"startMonth" yaml:"startMonth"`
	// entries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/cost_budget#entries CostBudget#entries}
	Entries interface{} `field:"optional" json:"entries" yaml:"entries"`
	// The ID of the budget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/cost_budget#id CostBudget#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

