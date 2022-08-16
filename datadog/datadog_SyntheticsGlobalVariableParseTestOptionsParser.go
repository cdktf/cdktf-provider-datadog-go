// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SyntheticsGlobalVariableParseTestOptionsParser struct {
	// Type of parser to extract the value. Valid values are `raw`, `json_path`, `regex`, `x_path`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_global_variable#type SyntheticsGlobalVariable#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Value for the parser to use, required for type `json_path` or `regex`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_global_variable#value SyntheticsGlobalVariable#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

