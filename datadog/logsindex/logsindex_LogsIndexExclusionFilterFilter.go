package logsindex


type LogsIndexExclusionFilterFilter struct {
	// Only logs matching the filter criteria and the query of the parent index will be considered for this exclusion filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_index#query LogsIndex#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// The fraction of logs excluded by the exclusion filter, when active.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_index#sample_rate LogsIndex#sample_rate}
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}
