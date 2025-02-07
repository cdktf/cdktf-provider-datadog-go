// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sensitivedatascannerrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SensitiveDataScannerRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Id of the scanning group the rule belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#group_id SensitiveDataScannerRule#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#description SensitiveDataScannerRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Attributes excluded from the scan. If namespaces is provided, it has to be a sub-path of the namespaces array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#excluded_namespaces SensitiveDataScannerRule#excluded_namespaces}
	ExcludedNamespaces *[]*string `field:"optional" json:"excludedNamespaces" yaml:"excludedNamespaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#id SensitiveDataScannerRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// included_keyword_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#included_keyword_configuration SensitiveDataScannerRule#included_keyword_configuration}
	IncludedKeywordConfiguration *SensitiveDataScannerRuleIncludedKeywordConfiguration `field:"optional" json:"includedKeywordConfiguration" yaml:"includedKeywordConfiguration"`
	// Whether or not the rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#is_enabled SensitiveDataScannerRule#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Name of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#name SensitiveDataScannerRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Attributes included in the scan.
	//
	// If namespaces is empty or missing, all attributes except excluded_namespaces are scanned. If both are missing the whole event is scanned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#namespaces SensitiveDataScannerRule#namespaces}
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// Not included if there is a relationship to a standard pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#pattern SensitiveDataScannerRule#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Priority level of the rule (optional).
	//
	// Used to order sensitive data discovered in the sds summary page. It must be between 1 and 5 (1 being the most important).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#priority SensitiveDataScannerRule#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Id of the standard pattern the rule refers to. If provided, then pattern must not be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#standard_pattern_id SensitiveDataScannerRule#standard_pattern_id}
	StandardPatternId *string `field:"optional" json:"standardPatternId" yaml:"standardPatternId"`
	// List of tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#tags SensitiveDataScannerRule#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// text_replacement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/sensitive_data_scanner_rule#text_replacement SensitiveDataScannerRule#text_replacement}
	TextReplacement *SensitiveDataScannerRuleTextReplacement `field:"optional" json:"textReplacement" yaml:"textReplacement"`
}

