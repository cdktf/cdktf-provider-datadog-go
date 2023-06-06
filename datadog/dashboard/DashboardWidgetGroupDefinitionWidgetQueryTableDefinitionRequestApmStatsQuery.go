package dashboard


type DashboardWidgetGroupDefinitionWidgetQueryTableDefinitionRequestApmStatsQuery struct {
	// The environment name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/dashboard#env Dashboard#env}
	Env *string `field:"required" json:"env" yaml:"env"`
	// The operation name associated with the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The organization's host group name and value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/dashboard#primary_tag Dashboard#primary_tag}
	PrimaryTag *string `field:"required" json:"primaryTag" yaml:"primaryTag"`
	// The level of detail for the request. Valid values are `service`, `resource`, `span`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/dashboard#row_type Dashboard#row_type}
	RowType *string `field:"required" json:"rowType" yaml:"rowType"`
	// The service name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/dashboard#service Dashboard#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/dashboard#columns Dashboard#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/dashboard#resource Dashboard#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

