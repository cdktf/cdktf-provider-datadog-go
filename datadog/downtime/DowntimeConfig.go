// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package downtime

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DowntimeConfig struct {
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
	// specify the group scope to which this downtime applies. For everything use '*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/downtime#scope Downtime#scope}
	Scope *[]*string `field:"required" json:"scope" yaml:"scope"`
	// Optionally specify an end date when this downtime should expire. Accepts a Unix timestamp in UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/downtime#end Downtime#end}
	End *float64 `field:"optional" json:"end" yaml:"end"`
	// String representing date and time to end the downtime in RFC3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/downtime#end_date Downtime#end_date}
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/downtime#id Downtime#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An optional message to provide when creating the downtime, can include notification handles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/downtime#message Downtime#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// When specified, this downtime will only apply to this monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/downtime#monitor_id Downtime#monitor_id}
	MonitorId *float64 `field:"optional" json:"monitorId" yaml:"monitorId"`
	// A list of monitor tags (up to 32) to base the scheduled downtime on.
	//
	// Only monitors that have all selected tags are silenced
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/downtime#monitor_tags Downtime#monitor_tags}
	MonitorTags *[]*string `field:"optional" json:"monitorTags" yaml:"monitorTags"`
	// When true the first recovery notification during the downtime will be muted Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/downtime#mute_first_recovery_notification Downtime#mute_first_recovery_notification}
	MuteFirstRecoveryNotification interface{} `field:"optional" json:"muteFirstRecoveryNotification" yaml:"muteFirstRecoveryNotification"`
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/downtime#recurrence Downtime#recurrence}
	Recurrence *DowntimeRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
	// Specify when this downtime should start. Accepts a Unix timestamp in UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/downtime#start Downtime#start}
	Start *float64 `field:"optional" json:"start" yaml:"start"`
	// String representing date and time to start the downtime in RFC3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/downtime#start_date Downtime#start_date}
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// The timezone for the downtime. Follows IANA timezone database identifiers. Defaults to `"UTC"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/downtime#timezone Downtime#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

