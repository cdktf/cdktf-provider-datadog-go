// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type LogsCustomPipelineProcessorPipelineProcessorCategoryProcessorCategory struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#filter LogsCustomPipeline#filter}
	Filter *LogsCustomPipelineProcessorPipelineProcessorCategoryProcessorCategoryFilter `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#name LogsCustomPipeline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}
