// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type LogsCustomPipelineProcessorGeoIpParser struct {
	// List of source attributes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#sources LogsCustomPipeline#sources}
	Sources *[]*string `field:"required" json:"sources" yaml:"sources"`
	// Name of the parent attribute that contains all the extracted details from the sources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#target LogsCustomPipeline#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// If the processor is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#is_enabled LogsCustomPipeline#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Name of the processor.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}
