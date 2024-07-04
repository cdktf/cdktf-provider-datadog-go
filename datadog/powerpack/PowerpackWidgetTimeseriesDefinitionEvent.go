// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetTimeseriesDefinitionEvent struct {
	// The event query to use in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#q Powerpack#q}
	Q *string `field:"required" json:"q" yaml:"q"`
	// The execution method for multi-value filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#tags_execution Powerpack#tags_execution}
	TagsExecution *string `field:"optional" json:"tagsExecution" yaml:"tagsExecution"`
}

