// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationslackchannel


type IntegrationSlackChannelDisplay struct {
	// Show the main body of the alert event. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/integration_slack_channel#message IntegrationSlackChannel#message}
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// Show interactive buttons to mute the alerting monitor. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/integration_slack_channel#mute_buttons IntegrationSlackChannel#mute_buttons}
	MuteButtons interface{} `field:"optional" json:"muteButtons" yaml:"muteButtons"`
	// Show the list of.
	Notified interface{} `field:"optional" json:"notified" yaml:"notified"`
	// Show the alert event's snapshot image. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/integration_slack_channel#snapshot IntegrationSlackChannel#snapshot}
	Snapshot interface{} `field:"optional" json:"snapshot" yaml:"snapshot"`
	// Show the scopes on which the monitor alerted. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/integration_slack_channel#tags IntegrationSlackChannel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

