package logscustompipeline


type LogsCustomPipelineProcessorCategoryProcessor struct {
	// category block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#category LogsCustomPipeline#category}
	Category interface{} `field:"required" json:"category" yaml:"category"`
	// Name of the target attribute whose value is defined by the matching category.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#target LogsCustomPipeline#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// If the processor is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#is_enabled LogsCustomPipeline#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Name of the category.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

