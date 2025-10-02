// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incidentnotificationtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IncidentNotificationTemplateConfig struct {
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
	// The category of the notification template. Valid values are `alert`, `incident`, `recovery`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/incident_notification_template#category IncidentNotificationTemplate#category}
	Category *string `field:"required" json:"category" yaml:"category"`
	// The content body of the notification template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/incident_notification_template#content IncidentNotificationTemplate#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The ID of the incident type this notification template is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/incident_notification_template#incident_type IncidentNotificationTemplate#incident_type}
	IncidentType *string `field:"required" json:"incidentType" yaml:"incidentType"`
	// The name of the notification template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/incident_notification_template#name IncidentNotificationTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The subject line of the notification template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/incident_notification_template#subject IncidentNotificationTemplate#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

