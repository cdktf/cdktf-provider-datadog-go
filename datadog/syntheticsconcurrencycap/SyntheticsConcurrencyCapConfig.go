// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsconcurrencycap

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsConcurrencyCapConfig struct {
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
	// Value of the on-demand concurrency cap, customizing the number of Synthetic tests run in parallel.
	//
	// Value must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/synthetics_concurrency_cap#on_demand_concurrency_cap SyntheticsConcurrencyCap#on_demand_concurrency_cap}
	OnDemandConcurrencyCap *float64 `field:"required" json:"onDemandConcurrencyCap" yaml:"onDemandConcurrencyCap"`
}

