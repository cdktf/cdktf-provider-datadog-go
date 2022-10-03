package syntheticstest


type SyntheticsTestApiStepAssertionTargetjsonpath struct {
	// The JSON path to assert.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#jsonpath SyntheticsTest#jsonpath}
	Jsonpath *string `field:"required" json:"jsonpath" yaml:"jsonpath"`
	// The specific operator to use on the path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#operator SyntheticsTest#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Expected matching value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#targetvalue SyntheticsTest#targetvalue}
	Targetvalue *string `field:"required" json:"targetvalue" yaml:"targetvalue"`
}

