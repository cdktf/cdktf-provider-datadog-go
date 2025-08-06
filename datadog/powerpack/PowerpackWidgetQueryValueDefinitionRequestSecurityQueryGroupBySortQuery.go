// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetQueryValueDefinitionRequestSecurityQueryGroupBySortQuery struct {
	// The aggregation method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#aggregation Powerpack#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// Widget sorting methods. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#order Powerpack#order}
	Order *string `field:"required" json:"order" yaml:"order"`
	// The facet name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#facet Powerpack#facet}
	Facet *string `field:"optional" json:"facet" yaml:"facet"`
}

