// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SyntheticsTestApiStepExtractedValue struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#name SyntheticsTest#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// parser block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#parser SyntheticsTest#parser}
	Parser *SyntheticsTestApiStepExtractedValueParser `field:"required" json:"parser" yaml:"parser"`
	// Property of the Synthetics Test Response to use for the variable. Valid values are `http_body`, `http_header`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// When type is `http_header`, name of the header to use to extract the value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#field SyntheticsTest#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
}

