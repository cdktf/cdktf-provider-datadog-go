package dashboard


type DashboardWidgetChangeDefinitionRequestLogQuery struct {
	// The name of the index to query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#index Dashboard#index}
	Index *string `field:"required" json:"index" yaml:"index"`
	// compute_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#compute_query Dashboard#compute_query}
	ComputeQuery *DashboardWidgetChangeDefinitionRequestLogQueryComputeQuery `field:"optional" json:"computeQuery" yaml:"computeQuery"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#group_by Dashboard#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// multi_compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#multi_compute Dashboard#multi_compute}
	MultiCompute interface{} `field:"optional" json:"multiCompute" yaml:"multiCompute"`
	// The search query to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#search_query Dashboard#search_query}
	SearchQuery *string `field:"optional" json:"searchQuery" yaml:"searchQuery"`
}

