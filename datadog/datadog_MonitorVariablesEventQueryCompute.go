// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type MonitorVariablesEventQueryCompute struct {
	// The aggregation methods for event platform queries.
	//
	// Valid values are `count`, `cardinality`, `median`, `pc75`, `pc90`, `pc95`, `pc98`, `pc99`, `sum`, `min`, `max`, `avg`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#aggregation Monitor#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// A time interval in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#interval Monitor#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The measurable attribute to compute.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#metric Monitor#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
}
