// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type MetricTagConfigurationAggregations struct {
	// A space aggregation for use in query. Valid values are `avg`, `max`, `min`, `sum`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/metric_tag_configuration#space MetricTagConfiguration#space}
	Space *string `field:"required" json:"space" yaml:"space"`
	// A time aggregation for use in query. Valid values are `avg`, `count`, `max`, `min`, `sum`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/metric_tag_configuration#time MetricTagConfiguration#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

