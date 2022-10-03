package dashboard


type DashboardWidgetGroupDefinitionWidgetSunburstDefinitionRequestQueryProcessQuery struct {
	// The data source for process queries. Valid values are `process`, `container`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#data_source Dashboard#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The process metric name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#metric Dashboard#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The aggregation methods available for metrics queries. Valid values are `avg`, `min`, `max`, `sum`, `last`, `area`, `l2norm`, `percentile`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#aggregator Dashboard#aggregator}
	Aggregator *string `field:"optional" json:"aggregator" yaml:"aggregator"`
	// Whether to normalize the CPU percentages.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#is_normalized_cpu Dashboard#is_normalized_cpu}
	IsNormalizedCpu interface{} `field:"optional" json:"isNormalizedCpu" yaml:"isNormalizedCpu"`
	// The number of hits to return.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#limit Dashboard#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// The direction of the sort. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#sort Dashboard#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// An array of tags to filter by.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#tag_filters Dashboard#tag_filters}
	TagFilters *[]*string `field:"optional" json:"tagFilters" yaml:"tagFilters"`
	// The text to use as a filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#text_filter Dashboard#text_filter}
	TextFilter *string `field:"optional" json:"textFilter" yaml:"textFilter"`
}

