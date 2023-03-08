package monitor


type MonitorVariables struct {
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#event_query Monitor#event_query}
	EventQuery interface{} `field:"optional" json:"eventQuery" yaml:"eventQuery"`
}

