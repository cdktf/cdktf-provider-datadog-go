package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorTraceIdRemapper struct {
	// List of source attributes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#sources LogsCustomPipeline#sources}
	Sources *[]*string `field:"required" json:"sources" yaml:"sources"`
	// If the processor is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#is_enabled LogsCustomPipeline#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Name of the processor.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}
