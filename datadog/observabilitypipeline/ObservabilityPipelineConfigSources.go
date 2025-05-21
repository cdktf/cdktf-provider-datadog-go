// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSources struct {
	// amazon_data_firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#amazon_data_firehose ObservabilityPipeline#amazon_data_firehose}
	AmazonDataFirehose interface{} `field:"optional" json:"amazonDataFirehose" yaml:"amazonDataFirehose"`
	// amazon_s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#amazon_s3 ObservabilityPipeline#amazon_s3}
	AmazonS3 interface{} `field:"optional" json:"amazonS3" yaml:"amazonS3"`
	// datadog_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#datadog_agent ObservabilityPipeline#datadog_agent}
	DatadogAgent interface{} `field:"optional" json:"datadogAgent" yaml:"datadogAgent"`
	// fluent_bit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#fluent_bit ObservabilityPipeline#fluent_bit}
	FluentBit interface{} `field:"optional" json:"fluentBit" yaml:"fluentBit"`
	// fluentd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#fluentd ObservabilityPipeline#fluentd}
	Fluentd interface{} `field:"optional" json:"fluentd" yaml:"fluentd"`
	// google_pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#google_pubsub ObservabilityPipeline#google_pubsub}
	GooglePubsub interface{} `field:"optional" json:"googlePubsub" yaml:"googlePubsub"`
	// http_client block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#http_client ObservabilityPipeline#http_client}
	HttpClient interface{} `field:"optional" json:"httpClient" yaml:"httpClient"`
	// http_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#http_server ObservabilityPipeline#http_server}
	HttpServer interface{} `field:"optional" json:"httpServer" yaml:"httpServer"`
	// kafka block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#kafka ObservabilityPipeline#kafka}
	Kafka interface{} `field:"optional" json:"kafka" yaml:"kafka"`
	// logstash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#logstash ObservabilityPipeline#logstash}
	Logstash interface{} `field:"optional" json:"logstash" yaml:"logstash"`
	// rsyslog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#rsyslog ObservabilityPipeline#rsyslog}
	Rsyslog interface{} `field:"optional" json:"rsyslog" yaml:"rsyslog"`
	// splunk_hec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#splunk_hec ObservabilityPipeline#splunk_hec}
	SplunkHec interface{} `field:"optional" json:"splunkHec" yaml:"splunkHec"`
	// splunk_tcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#splunk_tcp ObservabilityPipeline#splunk_tcp}
	SplunkTcp interface{} `field:"optional" json:"splunkTcp" yaml:"splunkTcp"`
	// sumo_logic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#sumo_logic ObservabilityPipeline#sumo_logic}
	SumoLogic interface{} `field:"optional" json:"sumoLogic" yaml:"sumoLogic"`
	// syslog_ng block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/observability_pipeline#syslog_ng ObservabilityPipeline#syslog_ng}
	SyslogNg interface{} `field:"optional" json:"syslogNg" yaml:"syslogNg"`
}

