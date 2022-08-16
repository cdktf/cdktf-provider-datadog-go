// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type LogsCustomPipelineProcessorPipelineProcessorGrokParserGrok struct {
	// Match rules for your grok parser.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#match_rules LogsCustomPipeline#match_rules}
	MatchRules *string `field:"required" json:"matchRules" yaml:"matchRules"`
	// Support rules for your grok parser.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#support_rules LogsCustomPipeline#support_rules}
	SupportRules *string `field:"required" json:"supportRules" yaml:"supportRules"`
}

