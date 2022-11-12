package monitor


type MonitorSchedulingOptions struct {
	// evaluation_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#evaluation_window Monitor#evaluation_window}
	EvaluationWindow interface{} `field:"required" json:"evaluationWindow" yaml:"evaluationWindow"`
}

