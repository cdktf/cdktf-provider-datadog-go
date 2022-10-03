package logsmetric


type LogsMetricCompute struct {
	// The type of aggregation to use. This field can't be updated after creation. Valid values are `count`, `distribution`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_metric#aggregation_type LogsMetric#aggregation_type}
	AggregationType *string `field:"required" json:"aggregationType" yaml:"aggregationType"`
	// The path to the value the log-based metric will aggregate on (only used if the aggregation type is a "distribution").
	//
	// This field can't be updated after creation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_metric#path LogsMetric#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

