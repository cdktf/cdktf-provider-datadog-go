package integrationslackchannel


type IntegrationSlackChannelDisplay struct {
	// Show the main body of the alert event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/integration_slack_channel#message IntegrationSlackChannel#message}
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// Show the list of @-handles in the alert event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/integration_slack_channel#notified IntegrationSlackChannel#notified}
	Notified interface{} `field:"optional" json:"notified" yaml:"notified"`
	// Show the alert event's snapshot image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/integration_slack_channel#snapshot IntegrationSlackChannel#snapshot}
	Snapshot interface{} `field:"optional" json:"snapshot" yaml:"snapshot"`
	// Show the scopes on which the monitor alerted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/integration_slack_channel#tags IntegrationSlackChannel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

