// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetListStreamDefinitionRequestQuery struct {
	// Source from which to query items to display in the stream.
	//
	// Valid values are `logs_stream`, `audit_stream`, `ci_pipeline_stream`, `ci_test_stream`, `rum_issue_stream`, `apm_issue_stream`, `trace_stream`, `logs_issue_stream`, `logs_pattern_stream`, `logs_transaction_stream`, `event_stream`, `rum_stream`, `llm_observability_stream`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#data_source Powerpack#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Specifies the field for logs pattern clustering. Can only be used with `logs_pattern_stream`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#clustering_pattern_field_path Powerpack#clustering_pattern_field_path}
	ClusteringPatternFieldPath *string `field:"optional" json:"clusteringPatternFieldPath" yaml:"clusteringPatternFieldPath"`
	// Size of events displayed in widget. Required if `data_source` is `event_stream`. Valid values are `s`, `l`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#event_size Powerpack#event_size}
	EventSize *string `field:"optional" json:"eventSize" yaml:"eventSize"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#group_by Powerpack#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// List of indexes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#indexes Powerpack#indexes}
	Indexes *[]*string `field:"optional" json:"indexes" yaml:"indexes"`
	// Widget query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#query_string Powerpack#query_string}
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#sort Powerpack#sort}
	Sort *PowerpackWidgetListStreamDefinitionRequestQuerySort `field:"optional" json:"sort" yaml:"sort"`
	// Storage location (private beta).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#storage Powerpack#storage}
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
}

