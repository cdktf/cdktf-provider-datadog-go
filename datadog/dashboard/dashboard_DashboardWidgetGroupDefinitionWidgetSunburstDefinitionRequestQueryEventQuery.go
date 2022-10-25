package dashboard


type DashboardWidgetGroupDefinitionWidgetSunburstDefinitionRequestQueryEventQuery struct {
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#compute Dashboard#compute}
	Compute interface{} `field:"required" json:"compute" yaml:"compute"`
	// The data source for event platform-based queries. Valid values are `logs`, `spans`, `network`, `rum`, `security_signals`, `profiles`, `audit`, `events`, `ci_tests`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#data_source Dashboard#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#group_by Dashboard#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// An array of index names to query in the stream.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#indexes Dashboard#indexes}
	Indexes *[]*string `field:"optional" json:"indexes" yaml:"indexes"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#search Dashboard#search}
	Search *DashboardWidgetGroupDefinitionWidgetSunburstDefinitionRequestQueryEventQuerySearch `field:"optional" json:"search" yaml:"search"`
	// Storage location (private beta).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#storage Dashboard#storage}
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
}

