// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type LogsMetricGroupBy struct {
	// The path to the value the log-based metric will be aggregated over.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_metric#path LogsMetric#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Name of the tag that gets created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_metric#tag_name LogsMetric#tag_name}
	TagName *string `field:"required" json:"tagName" yaml:"tagName"`
}
