package logsmetric


type LogsMetricFilter struct {
	// The search query - following the log search syntax.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_metric#query LogsMetric#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

