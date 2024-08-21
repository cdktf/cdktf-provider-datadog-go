// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringsuppression

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityMonitoringSuppressionConfig struct {
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
	// Whether the suppression rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/security_monitoring_suppression#enabled SecurityMonitoringSuppression#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the suppression rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/security_monitoring_suppression#name SecurityMonitoringSuppression#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The rule query of the suppression rule, with the same syntax as the search bar for detection rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/security_monitoring_suppression#rule_query SecurityMonitoringSuppression#rule_query}
	RuleQuery *string `field:"required" json:"ruleQuery" yaml:"ruleQuery"`
	// An exclusion query on the input data of the security rules, which could be logs, Agent events, or other types of data based on the security rule.
	//
	// Events matching this query are ignored by any detection rules referenced in the suppression rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/security_monitoring_suppression#data_exclusion_query SecurityMonitoringSuppression#data_exclusion_query}
	DataExclusionQuery *string `field:"optional" json:"dataExclusionQuery" yaml:"dataExclusionQuery"`
	// A description for the suppression rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/security_monitoring_suppression#description SecurityMonitoringSuppression#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A RFC3339 timestamp giving an expiration date for the suppression rule. After this date, it won't suppress signals anymore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/security_monitoring_suppression#expiration_date SecurityMonitoringSuppression#expiration_date}
	ExpirationDate *string `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// The suppression query of the suppression rule.
	//
	// If a signal matches this query, it is suppressed and is not triggered. It uses the same syntax as the queries to search signals in the Signals Explorer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/security_monitoring_suppression#suppression_query SecurityMonitoringSuppression#suppression_query}
	SuppressionQuery *string `field:"optional" json:"suppressionQuery" yaml:"suppressionQuery"`
}

