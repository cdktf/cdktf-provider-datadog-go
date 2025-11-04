// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetHostmapDefinition struct {
	// custom_link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#custom_link Powerpack#custom_link}
	CustomLink interface{} `field:"optional" json:"customLink" yaml:"customLink"`
	// The list of tags to group nodes by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#group Powerpack#group}
	Group *[]*string `field:"optional" json:"group" yaml:"group"`
	// The type of node used. Valid values are `host`, `container`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#node_type Powerpack#node_type}
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// A Boolean indicating whether to show ungrouped nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#no_group_hosts Powerpack#no_group_hosts}
	NoGroupHosts interface{} `field:"optional" json:"noGroupHosts" yaml:"noGroupHosts"`
	// A Boolean indicating whether to show nodes with no metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#no_metric_hosts Powerpack#no_metric_hosts}
	NoMetricHosts interface{} `field:"optional" json:"noMetricHosts" yaml:"noMetricHosts"`
	// request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#request Powerpack#request}
	Request *PowerpackWidgetHostmapDefinitionRequest `field:"optional" json:"request" yaml:"request"`
	// The list of tags to filter nodes by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#scope Powerpack#scope}
	Scope *[]*string `field:"optional" json:"scope" yaml:"scope"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#style Powerpack#style}
	Style *PowerpackWidgetHostmapDefinitionStyle `field:"optional" json:"style" yaml:"style"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#title Powerpack#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#title_align Powerpack#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#title_size Powerpack#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}

