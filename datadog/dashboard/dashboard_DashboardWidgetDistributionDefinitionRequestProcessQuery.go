package dashboard


type DashboardWidgetDistributionDefinitionRequestProcessQuery struct {
	// Your chosen metric.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#metric Dashboard#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// A list of processes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#filter_by Dashboard#filter_by}
	FilterBy *[]*string `field:"optional" json:"filterBy" yaml:"filterBy"`
	// The max number of items in the filter list.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#limit Dashboard#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Your chosen search term.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#search_by Dashboard#search_by}
	SearchBy *string `field:"optional" json:"searchBy" yaml:"searchBy"`
}
