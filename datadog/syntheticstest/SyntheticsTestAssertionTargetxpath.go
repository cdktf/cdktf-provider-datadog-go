package syntheticstest


type SyntheticsTestAssertionTargetxpath struct {
	// The specific operator to use on the path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#operator SyntheticsTest#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Expected matching value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#targetvalue SyntheticsTest#targetvalue}
	Targetvalue *string `field:"required" json:"targetvalue" yaml:"targetvalue"`
	// The xpath to assert.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#xpath SyntheticsTest#xpath}
	Xpath *string `field:"required" json:"xpath" yaml:"xpath"`
}

