// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SyntheticsTestApiStepRetry struct {
	// Number of retries needed to consider a location as failed before sending a notification alert.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#count SyntheticsTest#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Interval between a failed test and the next retry in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#interval SyntheticsTest#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

