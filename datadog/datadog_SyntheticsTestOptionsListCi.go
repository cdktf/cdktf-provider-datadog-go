// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SyntheticsTestOptionsListCi struct {
	// Execution rule for a Synthetics test. Valid values are `blocking`, `non_blocking`, `skipped`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#execution_rule SyntheticsTest#execution_rule}
	ExecutionRule *string `field:"optional" json:"executionRule" yaml:"executionRule"`
}

