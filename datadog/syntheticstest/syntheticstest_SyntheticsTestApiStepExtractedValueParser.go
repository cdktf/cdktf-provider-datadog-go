package syntheticstest


type SyntheticsTestApiStepExtractedValueParser struct {
	// Type of parser for a Synthetics global variable from a synthetics test. Valid values are `raw`, `json_path`, `regex`, `x_path`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Regex or JSON path used for the parser. Not used with type `raw`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#value SyntheticsTest#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}
