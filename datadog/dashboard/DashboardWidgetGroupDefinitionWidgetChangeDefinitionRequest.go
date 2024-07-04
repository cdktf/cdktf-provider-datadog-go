// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetChangeDefinitionRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#apm_query Dashboard#apm_query}
	ApmQuery *DashboardWidgetGroupDefinitionWidgetChangeDefinitionRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// Whether to show absolute or relative change. Valid values are `absolute`, `relative`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#change_type Dashboard#change_type}
	ChangeType *string `field:"optional" json:"changeType" yaml:"changeType"`
	// Choose from when to compare current data to. Valid values are `hour_before`, `day_before`, `week_before`, `month_before`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#compare_to Dashboard#compare_to}
	CompareTo *string `field:"optional" json:"compareTo" yaml:"compareTo"`
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#formula Dashboard#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// A Boolean indicating whether an increase in the value is good (displayed in green) or not (displayed in red).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#increase_good Dashboard#increase_good}
	IncreaseGood interface{} `field:"optional" json:"increaseGood" yaml:"increaseGood"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#log_query Dashboard#log_query}
	LogQuery *DashboardWidgetGroupDefinitionWidgetChangeDefinitionRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// What to order by. Valid values are `change`, `name`, `present`, `past`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#order_by Dashboard#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Widget sorting method. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#order_dir Dashboard#order_dir}
	OrderDir *string `field:"optional" json:"orderDir" yaml:"orderDir"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#process_query Dashboard#process_query}
	ProcessQuery *DashboardWidgetGroupDefinitionWidgetChangeDefinitionRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#q Dashboard#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#query Dashboard#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#rum_query Dashboard#rum_query}
	RumQuery *DashboardWidgetGroupDefinitionWidgetChangeDefinitionRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#security_query Dashboard#security_query}
	SecurityQuery *DashboardWidgetGroupDefinitionWidgetChangeDefinitionRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// If set to `true`, displays the current value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/dashboard#show_present Dashboard#show_present}
	ShowPresent interface{} `field:"optional" json:"showPresent" yaml:"showPresent"`
}

