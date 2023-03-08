package dashboard


type DashboardWidgetHostmapDefinitionRequestFill struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#apm_query Dashboard#apm_query}
	ApmQuery *DashboardWidgetHostmapDefinitionRequestFillApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#log_query Dashboard#log_query}
	LogQuery *DashboardWidgetHostmapDefinitionRequestFillLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#process_query Dashboard#process_query}
	ProcessQuery *DashboardWidgetHostmapDefinitionRequestFillProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#q Dashboard#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#rum_query Dashboard#rum_query}
	RumQuery *DashboardWidgetHostmapDefinitionRequestFillRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#security_query Dashboard#security_query}
	SecurityQuery *DashboardWidgetHostmapDefinitionRequestFillSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
}

