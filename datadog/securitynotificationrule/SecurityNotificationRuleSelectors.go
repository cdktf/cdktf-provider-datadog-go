// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitynotificationrule


type SecurityNotificationRuleSelectors struct {
	// Specifies security rule types for filtering signals and vulnerabilities that generate notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_notification_rule#rule_types SecurityNotificationRule#rule_types}
	RuleTypes *[]*string `field:"required" json:"ruleTypes" yaml:"ruleTypes"`
	// The type of security issues the rule applies to.
	//
	// Use `security_signals` for rules based on security signals and `security_findings` for those based on vulnerabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_notification_rule#trigger_source SecurityNotificationRule#trigger_source}
	TriggerSource *string `field:"required" json:"triggerSource" yaml:"triggerSource"`
	// Comprises one or several key:value pairs for filtering security issues based on tags and attributes. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_notification_rule#query SecurityNotificationRule#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// The security rules severities to consider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_notification_rule#severities SecurityNotificationRule#severities}
	Severities *[]*string `field:"optional" json:"severities" yaml:"severities"`
}

