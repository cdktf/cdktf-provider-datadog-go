// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetGeomapDefinitionRequestFormulaLimit struct {
	// The number of results to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/powerpack#count Powerpack#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The direction of the sort. Valid values are `asc`, `desc`. Defaults to `"desc"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/powerpack#order Powerpack#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

