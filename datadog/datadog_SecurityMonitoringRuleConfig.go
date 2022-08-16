// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityMonitoringRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// case block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#case SecurityMonitoringRule#case}
	Case interface{} `field:"required" json:"case" yaml:"case"`
	// Message for generated signals.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#message SecurityMonitoringRule#message}
	Message *string `field:"required" json:"message" yaml:"message"`
	// The name of the rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#name SecurityMonitoringRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#query SecurityMonitoringRule#query}
	Query interface{} `field:"required" json:"query" yaml:"query"`
	// Whether the rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#enabled SecurityMonitoringRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#filter SecurityMonitoringRule#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Whether the notifications include the triggering group-by values in their title.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#has_extended_title SecurityMonitoringRule#has_extended_title}
	HasExtendedTitle interface{} `field:"optional" json:"hasExtendedTitle" yaml:"hasExtendedTitle"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#id SecurityMonitoringRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#options SecurityMonitoringRule#options}
	Options *SecurityMonitoringRuleOptions `field:"optional" json:"options" yaml:"options"`
	// Tags for generated signals.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#tags SecurityMonitoringRule#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// The rule type. Valid values are `log_detection`, `infrastructure_configuration`, `workload_security`, `cloud_configuration`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#type SecurityMonitoringRule#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

