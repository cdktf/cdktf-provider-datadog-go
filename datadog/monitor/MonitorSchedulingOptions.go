package monitor


type MonitorSchedulingOptions struct {
	// evaluation_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/monitor#evaluation_window Monitor#evaluation_window}
	EvaluationWindow interface{} `field:"required" json:"evaluationWindow" yaml:"evaluationWindow"`
}

