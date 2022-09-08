// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SyntheticsGlobalVariableParseTestOptions struct {
	// parser block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_global_variable#parser SyntheticsGlobalVariable#parser}
	Parser *SyntheticsGlobalVariableParseTestOptionsParser `field:"required" json:"parser" yaml:"parser"`
	// Defines the source to use to extract the value. Valid values are `http_body`, `http_header`, `local_variable`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_global_variable#type SyntheticsGlobalVariable#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Required when type = `http_header`. Defines the header to use to extract the value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_global_variable#field SyntheticsGlobalVariable#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
}

