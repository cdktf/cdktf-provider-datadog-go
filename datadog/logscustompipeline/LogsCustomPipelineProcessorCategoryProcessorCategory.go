package logscustompipeline


type LogsCustomPipelineProcessorCategoryProcessorCategory struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.23.0/docs/resources/logs_custom_pipeline#filter LogsCustomPipeline#filter}
	Filter *LogsCustomPipelineProcessorCategoryProcessorCategoryFilter `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.23.0/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

