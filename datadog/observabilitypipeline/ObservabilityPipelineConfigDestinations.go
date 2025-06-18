// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinations struct {
	// amazon_opensearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#amazon_opensearch ObservabilityPipeline#amazon_opensearch}
	AmazonOpensearch interface{} `field:"optional" json:"amazonOpensearch" yaml:"amazonOpensearch"`
	// azure_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#azure_storage ObservabilityPipeline#azure_storage}
	AzureStorage interface{} `field:"optional" json:"azureStorage" yaml:"azureStorage"`
	// datadog_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#datadog_logs ObservabilityPipeline#datadog_logs}
	DatadogLogs interface{} `field:"optional" json:"datadogLogs" yaml:"datadogLogs"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#elasticsearch ObservabilityPipeline#elasticsearch}
	Elasticsearch interface{} `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// google_chronicle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#google_chronicle ObservabilityPipeline#google_chronicle}
	GoogleChronicle interface{} `field:"optional" json:"googleChronicle" yaml:"googleChronicle"`
	// google_cloud_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#google_cloud_storage ObservabilityPipeline#google_cloud_storage}
	GoogleCloudStorage interface{} `field:"optional" json:"googleCloudStorage" yaml:"googleCloudStorage"`
	// microsoft_sentinel block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#microsoft_sentinel ObservabilityPipeline#microsoft_sentinel}
	MicrosoftSentinel interface{} `field:"optional" json:"microsoftSentinel" yaml:"microsoftSentinel"`
	// new_relic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#new_relic ObservabilityPipeline#new_relic}
	NewRelic interface{} `field:"optional" json:"newRelic" yaml:"newRelic"`
	// opensearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#opensearch ObservabilityPipeline#opensearch}
	Opensearch interface{} `field:"optional" json:"opensearch" yaml:"opensearch"`
	// rsyslog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#rsyslog ObservabilityPipeline#rsyslog}
	Rsyslog interface{} `field:"optional" json:"rsyslog" yaml:"rsyslog"`
	// sentinel_one block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#sentinel_one ObservabilityPipeline#sentinel_one}
	SentinelOne interface{} `field:"optional" json:"sentinelOne" yaml:"sentinelOne"`
	// splunk_hec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#splunk_hec ObservabilityPipeline#splunk_hec}
	SplunkHec interface{} `field:"optional" json:"splunkHec" yaml:"splunkHec"`
	// sumo_logic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#sumo_logic ObservabilityPipeline#sumo_logic}
	SumoLogic interface{} `field:"optional" json:"sumoLogic" yaml:"sumoLogic"`
	// syslog_ng block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#syslog_ng ObservabilityPipeline#syslog_ng}
	SyslogNg interface{} `field:"optional" json:"syslogNg" yaml:"syslogNg"`
}

