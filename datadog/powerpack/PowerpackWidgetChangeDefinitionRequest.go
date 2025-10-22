// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetChangeDefinitionRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#apm_query Powerpack#apm_query}
	ApmQuery *PowerpackWidgetChangeDefinitionRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// Whether to show absolute or relative change. Valid values are `absolute`, `relative`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#change_type Powerpack#change_type}
	ChangeType *string `field:"optional" json:"changeType" yaml:"changeType"`
	// Choose from when to compare current data to. Valid values are `hour_before`, `day_before`, `week_before`, `month_before`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#compare_to Powerpack#compare_to}
	CompareTo *string `field:"optional" json:"compareTo" yaml:"compareTo"`
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#formula Powerpack#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// A Boolean indicating whether an increase in the value is good (displayed in green) or not (displayed in red).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#increase_good Powerpack#increase_good}
	IncreaseGood interface{} `field:"optional" json:"increaseGood" yaml:"increaseGood"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#log_query Powerpack#log_query}
	LogQuery *PowerpackWidgetChangeDefinitionRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// What to order by. Valid values are `change`, `name`, `present`, `past`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#order_by Powerpack#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Widget sorting method. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#order_dir Powerpack#order_dir}
	OrderDir *string `field:"optional" json:"orderDir" yaml:"orderDir"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#process_query Powerpack#process_query}
	ProcessQuery *PowerpackWidgetChangeDefinitionRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#q Powerpack#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#query Powerpack#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#rum_query Powerpack#rum_query}
	RumQuery *PowerpackWidgetChangeDefinitionRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#security_query Powerpack#security_query}
	SecurityQuery *PowerpackWidgetChangeDefinitionRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// If set to `true`, displays the current value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/powerpack#show_present Powerpack#show_present}
	ShowPresent interface{} `field:"optional" json:"showPresent" yaml:"showPresent"`
}

