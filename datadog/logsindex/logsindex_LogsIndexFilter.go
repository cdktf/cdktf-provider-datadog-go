package logsindex


type LogsIndexFilter struct {
	// Logs filter criteria. Only logs matching this filter criteria are considered for this index.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_index#query LogsIndex#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

