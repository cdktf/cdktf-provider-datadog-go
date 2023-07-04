package logscustompipeline


type LogsCustomPipelineProcessorGrokParserGrok struct {
	// Match rules for your grok parser.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/logs_custom_pipeline#match_rules LogsCustomPipeline#match_rules}
	MatchRules *string `field:"required" json:"matchRules" yaml:"matchRules"`
	// Support rules for your grok parser.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/logs_custom_pipeline#support_rules LogsCustomPipeline#support_rules}
	SupportRules *string `field:"required" json:"supportRules" yaml:"supportRules"`
}

