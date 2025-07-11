// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessor struct {
	// arithmetic_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#arithmetic_processor LogsCustomPipeline#arithmetic_processor}
	ArithmeticProcessor *LogsCustomPipelineProcessorArithmeticProcessor `field:"optional" json:"arithmeticProcessor" yaml:"arithmeticProcessor"`
	// attribute_remapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#attribute_remapper LogsCustomPipeline#attribute_remapper}
	AttributeRemapper *LogsCustomPipelineProcessorAttributeRemapper `field:"optional" json:"attributeRemapper" yaml:"attributeRemapper"`
	// category_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#category_processor LogsCustomPipeline#category_processor}
	CategoryProcessor *LogsCustomPipelineProcessorCategoryProcessor `field:"optional" json:"categoryProcessor" yaml:"categoryProcessor"`
	// date_remapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#date_remapper LogsCustomPipeline#date_remapper}
	DateRemapper *LogsCustomPipelineProcessorDateRemapper `field:"optional" json:"dateRemapper" yaml:"dateRemapper"`
	// geo_ip_parser block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#geo_ip_parser LogsCustomPipeline#geo_ip_parser}
	GeoIpParser *LogsCustomPipelineProcessorGeoIpParser `field:"optional" json:"geoIpParser" yaml:"geoIpParser"`
	// grok_parser block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#grok_parser LogsCustomPipeline#grok_parser}
	GrokParser *LogsCustomPipelineProcessorGrokParser `field:"optional" json:"grokParser" yaml:"grokParser"`
	// lookup_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#lookup_processor LogsCustomPipeline#lookup_processor}
	LookupProcessor *LogsCustomPipelineProcessorLookupProcessor `field:"optional" json:"lookupProcessor" yaml:"lookupProcessor"`
	// message_remapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#message_remapper LogsCustomPipeline#message_remapper}
	MessageRemapper *LogsCustomPipelineProcessorMessageRemapper `field:"optional" json:"messageRemapper" yaml:"messageRemapper"`
	// pipeline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#pipeline LogsCustomPipeline#pipeline}
	Pipeline *LogsCustomPipelineProcessorPipeline `field:"optional" json:"pipeline" yaml:"pipeline"`
	// reference_table_lookup_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#reference_table_lookup_processor LogsCustomPipeline#reference_table_lookup_processor}
	ReferenceTableLookupProcessor *LogsCustomPipelineProcessorReferenceTableLookupProcessor `field:"optional" json:"referenceTableLookupProcessor" yaml:"referenceTableLookupProcessor"`
	// service_remapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#service_remapper LogsCustomPipeline#service_remapper}
	ServiceRemapper *LogsCustomPipelineProcessorServiceRemapper `field:"optional" json:"serviceRemapper" yaml:"serviceRemapper"`
	// span_id_remapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#span_id_remapper LogsCustomPipeline#span_id_remapper}
	SpanIdRemapper *LogsCustomPipelineProcessorSpanIdRemapper `field:"optional" json:"spanIdRemapper" yaml:"spanIdRemapper"`
	// status_remapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#status_remapper LogsCustomPipeline#status_remapper}
	StatusRemapper *LogsCustomPipelineProcessorStatusRemapper `field:"optional" json:"statusRemapper" yaml:"statusRemapper"`
	// string_builder_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#string_builder_processor LogsCustomPipeline#string_builder_processor}
	StringBuilderProcessor *LogsCustomPipelineProcessorStringBuilderProcessor `field:"optional" json:"stringBuilderProcessor" yaml:"stringBuilderProcessor"`
	// trace_id_remapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#trace_id_remapper LogsCustomPipeline#trace_id_remapper}
	TraceIdRemapper *LogsCustomPipelineProcessorTraceIdRemapper `field:"optional" json:"traceIdRemapper" yaml:"traceIdRemapper"`
	// url_parser block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#url_parser LogsCustomPipeline#url_parser}
	UrlParser *LogsCustomPipelineProcessorUrlParser `field:"optional" json:"urlParser" yaml:"urlParser"`
	// user_agent_parser block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/logs_custom_pipeline#user_agent_parser LogsCustomPipeline#user_agent_parser}
	UserAgentParser *LogsCustomPipelineProcessorUserAgentParser `field:"optional" json:"userAgentParser" yaml:"userAgentParser"`
}

