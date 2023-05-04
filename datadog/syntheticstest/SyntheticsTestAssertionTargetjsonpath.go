package syntheticstest


type SyntheticsTestAssertionTargetjsonpath struct {
	// The JSON path to assert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/synthetics_test#jsonpath SyntheticsTest#jsonpath}
	Jsonpath *string `field:"required" json:"jsonpath" yaml:"jsonpath"`
	// The specific operator to use on the path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/synthetics_test#operator SyntheticsTest#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Expected matching value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/synthetics_test#targetvalue SyntheticsTest#targetvalue}
	Targetvalue *string `field:"optional" json:"targetvalue" yaml:"targetvalue"`
}

