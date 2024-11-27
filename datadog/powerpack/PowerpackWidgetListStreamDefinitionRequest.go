// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetListStreamDefinitionRequest struct {
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#columns Powerpack#columns}
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#query Powerpack#query}
	Query *PowerpackWidgetListStreamDefinitionRequestQuery `field:"required" json:"query" yaml:"query"`
	// Widget response format. Valid values are `event_list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#response_format Powerpack#response_format}
	ResponseFormat *string `field:"required" json:"responseFormat" yaml:"responseFormat"`
}

