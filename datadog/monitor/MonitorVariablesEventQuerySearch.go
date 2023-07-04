package monitor


type MonitorVariablesEventQuerySearch struct {
	// The events search string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/monitor#query Monitor#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

