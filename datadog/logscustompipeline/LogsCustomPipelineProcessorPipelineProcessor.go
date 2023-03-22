package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessor struct {
	// arithmetic_processor block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#arithmetic_processor LogsCustomPipeline#arithmetic_processor}
	ArithmeticProcessor *LogsCustomPipelineProcessorPipelineProcessorArithmeticProcessor `field:"optional" json:"arithmeticProcessor" yaml:"arithmeticProcessor"`
	// attribute_remapper block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#attribute_remapper LogsCustomPipeline#attribute_remapper}
	AttributeRemapper *LogsCustomPipelineProcessorPipelineProcessorAttributeRemapper `field:"optional" json:"attributeRemapper" yaml:"attributeRemapper"`
	// category_processor block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#category_processor LogsCustomPipeline#category_processor}
	CategoryProcessor *LogsCustomPipelineProcessorPipelineProcessorCategoryProcessor `field:"optional" json:"categoryProcessor" yaml:"categoryProcessor"`
	// date_remapper block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#date_remapper LogsCustomPipeline#date_remapper}
	DateRemapper *LogsCustomPipelineProcessorPipelineProcessorDateRemapper `field:"optional" json:"dateRemapper" yaml:"dateRemapper"`
	// geo_ip_parser block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#geo_ip_parser LogsCustomPipeline#geo_ip_parser}
	GeoIpParser *LogsCustomPipelineProcessorPipelineProcessorGeoIpParser `field:"optional" json:"geoIpParser" yaml:"geoIpParser"`
	// grok_parser block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#grok_parser LogsCustomPipeline#grok_parser}
	GrokParser *LogsCustomPipelineProcessorPipelineProcessorGrokParser `field:"optional" json:"grokParser" yaml:"grokParser"`
	// lookup_processor block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#lookup_processor LogsCustomPipeline#lookup_processor}
	LookupProcessor *LogsCustomPipelineProcessorPipelineProcessorLookupProcessor `field:"optional" json:"lookupProcessor" yaml:"lookupProcessor"`
	// message_remapper block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#message_remapper LogsCustomPipeline#message_remapper}
	MessageRemapper *LogsCustomPipelineProcessorPipelineProcessorMessageRemapper `field:"optional" json:"messageRemapper" yaml:"messageRemapper"`
	// reference_table_lookup_processor block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#reference_table_lookup_processor LogsCustomPipeline#reference_table_lookup_processor}
	ReferenceTableLookupProcessor *LogsCustomPipelineProcessorPipelineProcessorReferenceTableLookupProcessor `field:"optional" json:"referenceTableLookupProcessor" yaml:"referenceTableLookupProcessor"`
	// service_remapper block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#service_remapper LogsCustomPipeline#service_remapper}
	ServiceRemapper *LogsCustomPipelineProcessorPipelineProcessorServiceRemapper `field:"optional" json:"serviceRemapper" yaml:"serviceRemapper"`
	// status_remapper block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#status_remapper LogsCustomPipeline#status_remapper}
	StatusRemapper *LogsCustomPipelineProcessorPipelineProcessorStatusRemapper `field:"optional" json:"statusRemapper" yaml:"statusRemapper"`
	// string_builder_processor block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#string_builder_processor LogsCustomPipeline#string_builder_processor}
	StringBuilderProcessor *LogsCustomPipelineProcessorPipelineProcessorStringBuilderProcessor `field:"optional" json:"stringBuilderProcessor" yaml:"stringBuilderProcessor"`
	// trace_id_remapper block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#trace_id_remapper LogsCustomPipeline#trace_id_remapper}
	TraceIdRemapper *LogsCustomPipelineProcessorPipelineProcessorTraceIdRemapper `field:"optional" json:"traceIdRemapper" yaml:"traceIdRemapper"`
	// url_parser block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#url_parser LogsCustomPipeline#url_parser}
	UrlParser *LogsCustomPipelineProcessorPipelineProcessorUrlParser `field:"optional" json:"urlParser" yaml:"urlParser"`
	// user_agent_parser block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#user_agent_parser LogsCustomPipeline#user_agent_parser}
	UserAgentParser *LogsCustomPipelineProcessorPipelineProcessorUserAgentParser `field:"optional" json:"userAgentParser" yaml:"userAgentParser"`
}

