// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessors struct {
	// add_env_vars block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#add_env_vars ObservabilityPipeline#add_env_vars}
	AddEnvVars interface{} `field:"optional" json:"addEnvVars" yaml:"addEnvVars"`
	// add_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#add_fields ObservabilityPipeline#add_fields}
	AddFields interface{} `field:"optional" json:"addFields" yaml:"addFields"`
	// custom_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#custom_processor ObservabilityPipeline#custom_processor}
	CustomProcessor interface{} `field:"optional" json:"customProcessor" yaml:"customProcessor"`
	// datadog_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#datadog_tags ObservabilityPipeline#datadog_tags}
	DatadogTags interface{} `field:"optional" json:"datadogTags" yaml:"datadogTags"`
	// dedupe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#dedupe ObservabilityPipeline#dedupe}
	Dedupe interface{} `field:"optional" json:"dedupe" yaml:"dedupe"`
	// enrichment_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#enrichment_table ObservabilityPipeline#enrichment_table}
	EnrichmentTable interface{} `field:"optional" json:"enrichmentTable" yaml:"enrichmentTable"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#filter ObservabilityPipeline#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// generate_datadog_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#generate_datadog_metrics ObservabilityPipeline#generate_datadog_metrics}
	GenerateDatadogMetrics interface{} `field:"optional" json:"generateDatadogMetrics" yaml:"generateDatadogMetrics"`
	// ocsf_mapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#ocsf_mapper ObservabilityPipeline#ocsf_mapper}
	OcsfMapper interface{} `field:"optional" json:"ocsfMapper" yaml:"ocsfMapper"`
	// parse_grok block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#parse_grok ObservabilityPipeline#parse_grok}
	ParseGrok interface{} `field:"optional" json:"parseGrok" yaml:"parseGrok"`
	// parse_json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#parse_json ObservabilityPipeline#parse_json}
	ParseJson interface{} `field:"optional" json:"parseJson" yaml:"parseJson"`
	// quota block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#quota ObservabilityPipeline#quota}
	Quota interface{} `field:"optional" json:"quota" yaml:"quota"`
	// reduce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#reduce ObservabilityPipeline#reduce}
	Reduce interface{} `field:"optional" json:"reduce" yaml:"reduce"`
	// remove_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#remove_fields ObservabilityPipeline#remove_fields}
	RemoveFields interface{} `field:"optional" json:"removeFields" yaml:"removeFields"`
	// rename_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#rename_fields ObservabilityPipeline#rename_fields}
	RenameFields interface{} `field:"optional" json:"renameFields" yaml:"renameFields"`
	// sample block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#sample ObservabilityPipeline#sample}
	Sample interface{} `field:"optional" json:"sample" yaml:"sample"`
	// sensitive_data_scanner block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#sensitive_data_scanner ObservabilityPipeline#sensitive_data_scanner}
	SensitiveDataScanner interface{} `field:"optional" json:"sensitiveDataScanner" yaml:"sensitiveDataScanner"`
	// throttle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/observability_pipeline#throttle ObservabilityPipeline#throttle}
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
}

