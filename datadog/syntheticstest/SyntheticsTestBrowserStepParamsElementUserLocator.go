package syntheticstest


type SyntheticsTestBrowserStepParamsElementUserLocator struct {
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/synthetics_test#value SyntheticsTest#value}
	Value *SyntheticsTestBrowserStepParamsElementUserLocatorValue `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/synthetics_test#fail_test_on_cannot_locate SyntheticsTest#fail_test_on_cannot_locate}.
	FailTestOnCannotLocate interface{} `field:"optional" json:"failTestOnCannotLocate" yaml:"failTestOnCannotLocate"`
}

