package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorCategoryProcessorCategoryFilter struct {
	// Filter criteria of the category.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#query LogsCustomPipeline#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}
