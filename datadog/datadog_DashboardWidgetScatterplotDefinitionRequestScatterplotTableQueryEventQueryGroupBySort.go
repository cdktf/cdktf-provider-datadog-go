// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQueryGroupBySort struct {
	// The aggregation methods for the event platform queries.
	//
	// Valid values are `count`, `cardinality`, `median`, `pc75`, `pc90`, `pc95`, `pc98`, `pc99`, `sum`, `min`, `max`, `avg`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#aggregation Dashboard#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// The metric used for sorting group by results.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#metric Dashboard#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Direction of sort. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#order Dashboard#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

