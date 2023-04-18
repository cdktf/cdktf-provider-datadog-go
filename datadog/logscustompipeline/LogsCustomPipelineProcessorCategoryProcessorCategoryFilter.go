package logscustompipeline


type LogsCustomPipelineProcessorCategoryProcessorCategoryFilter struct {
	// Filter criteria of the category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.23.0/docs/resources/logs_custom_pipeline#query LogsCustomPipeline#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

