// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetGroupDefinitionWidgetHostmapDefinition struct {
	// custom_link block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#custom_link Dashboard#custom_link}
	CustomLink interface{} `field:"optional" json:"customLink" yaml:"customLink"`
	// The list of tags to group nodes by.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#group Dashboard#group}
	Group *[]*string `field:"optional" json:"group" yaml:"group"`
	// The type of node used. Valid values are `host`, `container`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#node_type Dashboard#node_type}
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// A Boolean indicating whether to show ungrouped nodes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#no_group_hosts Dashboard#no_group_hosts}
	NoGroupHosts interface{} `field:"optional" json:"noGroupHosts" yaml:"noGroupHosts"`
	// A Boolean indicating whether to show nodes with no metrics.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#no_metric_hosts Dashboard#no_metric_hosts}
	NoMetricHosts interface{} `field:"optional" json:"noMetricHosts" yaml:"noMetricHosts"`
	// request block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#request Dashboard#request}
	Request *DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequest `field:"optional" json:"request" yaml:"request"`
	// The list of tags to filter nodes by.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#scope Dashboard#scope}
	Scope *[]*string `field:"optional" json:"scope" yaml:"scope"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#style Dashboard#style}
	Style *DashboardWidgetGroupDefinitionWidgetHostmapDefinitionStyle `field:"optional" json:"style" yaml:"style"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#title Dashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#title_align Dashboard#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#title_size Dashboard#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}
