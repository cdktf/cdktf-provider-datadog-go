package syntheticstest


type SyntheticsTestOptionsListCi struct {
	// Execution rule for a Synthetics test. Valid values are `blocking`, `non_blocking`, `skipped`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/synthetics_test#execution_rule SyntheticsTest#execution_rule}
	ExecutionRule *string `field:"optional" json:"executionRule" yaml:"executionRule"`
}

